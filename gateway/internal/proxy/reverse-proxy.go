package proxy

import (
	"context"
	"github.com/CyanPigeon/toktik/gateway/internal"
	"github.com/CyanPigeon/toktik/gateway/internal/config"
	"github.com/CyanPigeon/toktik/gateway/internal/discovery"
	"github.com/CyanPigeon/toktik/gateway/internal/router"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

type ReverseProxyServer struct {
	ctx       context.Context
	router    router.Router
	watcher   registry.Watcher
	discovery discovery.Discovery
	server    *http.Server
}

type handler struct {
}

func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(404)
}

func NewServer(ctx context.Context, discovery discovery.Discovery, router router.Router) (*ReverseProxyServer, error) {
	server := &ReverseProxyServer{
		ctx:       ctx,
		router:    router,
		discovery: discovery,
		server: &http.Server{
			Handler: h2c.NewHandler(router, &http2.Server{
				IdleTimeout:          config.Server.IdleTimeout,
				MaxConcurrentStreams: config.Server.MaxCurrentStream,
			}),
			ReadHeaderTimeout: config.Server.ReadHeaderTimeout,
			ReadTimeout:       config.Server.ReadTimeout,
			WriteTimeout:      config.Server.WriteTimeout,
			IdleTimeout:       config.Server.IdleTimeout,
		},
		watcher: nil,
	}
	return server, server.tick()
}

func (t *ReverseProxyServer) tick() (err error) {
	if t.watcher != nil {
		err = t.watcher.Stop()
		if err != nil {
			log.Errorf("node discovery loop start failed caused by: %+v", err)
		}
	}
	t.watcher, err = t.discovery.Watch(t.ctx, discovery.DefaultServiceFilter)
	if err != nil {
		return err
	}
	go func() {
		for {
			if t.watcher == nil {
				return
			}
			services, e := t.watcher.Next()
			log.Info("microservice list modified.")
			if e != nil {
				return
			}
			nodes := make([]discovery.Node, 0, len(services))
			for _, service := range services {
				var path string
				path, e = internal.ParseEndpoint(service.Endpoints, "http")
				if e != nil {
					log.Errorf("ignore service '%s' caused by: ", e)
					continue
				}
				nodes = append(nodes, discovery.NewNode("http", path, service))
			}
			log.Infof("nodes: %+v", nodes)
			if e = t.router.Apply(nodes); e != nil {
				log.Errorf("update router failed: %+v", e)
			}
		}
	}()
	return nil
}

func (t *ReverseProxyServer) Listen(address string) error {
	t.server.Addr = address
	err := t.server.ListenAndServe()
	return err
}

func (t *ReverseProxyServer) Shutdown() error {
	return t.server.Shutdown(t.ctx)
}

func (t *ReverseProxyServer) Close() error {
	return t.server.Close()
}
