package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/CyanPigeon/toktik/gateway/internal"
	"github.com/CyanPigeon/toktik/gateway/internal/config"
	"github.com/CyanPigeon/toktik/gateway/internal/discovery"
	"github.com/CyanPigeon/toktik/gateway/internal/http"
	"github.com/CyanPigeon/toktik/gateway/internal/http/endpoint"
	"github.com/CyanPigeon/toktik/gateway/internal/http/interceptor"
	"github.com/CyanPigeon/toktik/gateway/internal/interceptors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/gorilla/mux"
	std "net/http"
	"runtime"
	"strings"
	"sync"
)

func WithErrorWriter(writer http.ErrorWriter) internal.OptionSetter[MuxRouter] {
	return func(t *MuxRouter) {
		t.errorWriter = writer
	}
}
func WithPreInterceptors(handlers ...interceptor.PreHandler) internal.OptionSetter[MuxRouter] {
	return func(o *MuxRouter) {
		if len(handlers) <= 0 {
			return
		}
		pre := &o.Interceptors.(*interceptor.DefaultInterceptors).Pre
		*pre = append(handlers, *pre...)
	}
}
func WithPostInterceptors(handlers ...interceptor.PostHandler) internal.OptionSetter[MuxRouter] {
	return func(o *MuxRouter) {
		if len(handlers) <= 0 {
			return
		}
		post := &o.Interceptors.(*interceptor.DefaultInterceptors).Post
		*post = append(handlers, *post...)
	}
}

type MuxRouter struct {
	interceptor.Interceptors
	errorWriter http.ErrorWriter
	factory     endpoint.Factory
	ctx         context.Context
	cancel      context.CancelFunc
	router      *mux.Router
	mutex       sync.Mutex
	semaphore   *sync.WaitGroup
	endpoints   map[string]http.Endpoint
	nodes       map[string]*[]selector.Node
}

func New(ctx context.Context, factory endpoint.Factory, opts ...internal.OptionSetter[MuxRouter]) Router {
	r := &MuxRouter{
		Interceptors: &interceptor.DefaultInterceptors{},
		errorWriter:  http.NewErrorWriter(),
		factory:      factory,
		mutex:        sync.Mutex{},
		semaphore:    &sync.WaitGroup{},
		endpoints:    make(map[string]http.Endpoint),
		nodes:        make(map[string]*[]selector.Node),
	}
	WithPreInterceptors(interceptors.XForwardInterceptor)(r)
	for _, opt := range opts {
		opt(r)
	}

	r.ctx, r.cancel = context.WithCancel(ctx)
	return r
}
func (t *MuxRouter) register(endpoint http.Endpoint, nodes []selector.Node) error {
	if _, ok := t.endpoints[endpoint.Path()]; !ok {
		t.endpoints[endpoint.Path()] = endpoint
	} else {
		return fmt.Errorf("the specified route already exists")
	}
	endpoint.Apply(nodes)

	route := t.router.NewRoute().HandlerFunc(func(writer std.ResponseWriter, request *std.Request) {
		endpoint.ServeHTTPEx(writer, request, t.Interceptors)
	})
	if strings.HasSuffix(endpoint.Path(), "*") {
		route = route.PathPrefix(strings.TrimRight(endpoint.Path(), "*"))
	} else {
		route = route.Path(endpoint.Path())
	}
	if err := route.GetError(); err != nil {
		return err
	}
	log.Infof("endpoint '%s' registered.", endpoint.Path())
	return nil
}
func (t *MuxRouter) Register(path string) error {
	nodes, ok := t.nodes[path]
	if !ok {
		return errors.New("invalid path")
	}
	return t.register(t.factory(path, t.errorWriter), *nodes)
}

func (t *MuxRouter) Apply(nodes []discovery.Node) (err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.nodes = make(map[string]*[]selector.Node, len(nodes))
	for _, node := range nodes {
		list, ok := t.nodes[node.PathPrefix()]
		if !ok {
			list = new([]selector.Node)
			t.nodes[node.PathPrefix()] = list
		}
		*list = append(*list, node)
	}

	endpoints := t.endpoints
	t.endpoints = make(map[string]http.Endpoint, len(nodes))
	t.router = mux.NewRouter()

	for path, ns := range t.nodes {
		if ept, ok := endpoints[path]; ok {
			err = t.register(ept, *ns)
		} else {
			err = t.register(t.factory(path, t.errorWriter), *ns)
		}
		if err != nil {
			if config.Router.StopApplyWhenError {
				return err
			} else {
				log.Errorf("Route '%s' failed: %+v", path, err)
			}
		}
	}

	return nil
}

func (t *MuxRouter) ServeHTTP(writer std.ResponseWriter, request *std.Request) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		writer.WriteHeader(std.StatusBadGateway)
		buf := make([]byte, 64<<10)
		n := runtime.Stack(buf, false)
		log.Errorf("panic recovered: %s", buf[:n])
		t.errorWriter.Write(writer, std.StatusInternalServerError, errors.New("500 Internal Server Error"))
	}()
	if t.router == nil {
		t.errorWriter.Write(writer, std.StatusNotFound, errors.New("404 Page Not Found"))
		return
	}
	t.semaphore.Add(1)
	defer t.semaphore.Done()
	t.router.ServeHTTP(writer, request)
}

func (t *MuxRouter) Context() context.Context {
	return t.ctx
}
func (t *MuxRouter) Endpoints() map[string]http.Endpoint {
	return t.endpoints
}

func (t *MuxRouter) Close() error {
	timeout := func() bool {
		c := make(chan struct{})
		go func() { defer close(c); t.semaphore.Wait() }()
		select {
		case <-c:
			return false
		case <-t.ctx.Done():
			return true
		}
	}()
	if timeout {
		log.Warnf("Timeout to wait all request complete.")
	}
	for _, closer := range t.endpoints {
		if err := closer.Close(); err != nil {
			log.Errorf("Failed to close an endpoint: %+v", err)
		}
	}
	return nil
}
