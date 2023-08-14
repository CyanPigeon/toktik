package endpoint

import (
	"bytes"
	"context"
	"fmt"
	"github.com/CyanPigeon/toktik/gateway/internal/config"
	"github.com/CyanPigeon/toktik/gateway/internal/http"
	"github.com/CyanPigeon/toktik/gateway/internal/http/interceptor"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/selector"
	"io"
	std "net/http"
)

type Factory func(path string, errorWriter http.ErrorWriter) http.Endpoint
type endpoint struct {
	path        string
	errorWriter http.ErrorWriter
	transport   http.Transport
	selector    selector.Selector
	//filters     []selector.NodeFilter
}

func NewEndpointFactory(transport http.Transport) Factory {
	return func(path string, errorWriter http.ErrorWriter) http.Endpoint {
		return &endpoint{
			path:        path,
			errorWriter: errorWriter,
			transport:   transport,
			selector:    selector.GlobalSelector().Build(),
		}
	}
}

func (t *endpoint) Path() string {
	return t.path
}

func (t *endpoint) ServeHTTPEx(writer std.ResponseWriter, request *std.Request, interceptors interceptor.Interceptors) {
	originURL := request.URL.String()
	a := action{
		transport:    t.transport,
		selector:     t.selector,
		Interceptors: interceptors,
	}
	statusCode, err := a.forward(writer, request)

	name, version, url := "???", "?.?.?", "???"
	if a.selected != nil {
		name, version, url = a.selected.ServiceName(), a.selected.Version(), request.URL.String()
	}
	message := fmt.Sprintf(
		"%s | %s | %d - %d - %s -> [%s-%s] %s",
		request.Method, t.path, statusCode,
		a.sent, originURL, name, version, url,
	)

	if err != nil {
		a.done(context.Background(), selector.DoneInfo{Err: err})
	} else if a.response == nil {
		err = fmt.Errorf("unexcepted branch")
		a.done(context.Background(), selector.DoneInfo{Err: err})
	} else {
		a.done(context.Background(), selector.DoneInfo{ReplyMD: a.response.Trailer})
	}

	if err != nil {
		t.errorWriter.Write(writer, statusCode, err)
		log.Error(message)
	} else {
		log.Info(message)
	}
}

func (t *endpoint) Apply(nodes []selector.Node) {
	t.selector.Apply(nodes)
}

func (t *endpoint) Close() error {
	return t.transport.Close()
}

type action struct {
	interceptor.Interceptors
	transport std.RoundTripper
	selector  selector.Selector
	filters   []selector.NodeFilter

	response *std.Response
	sent     int64
	done     selector.DoneFunc
	selected selector.Node
}

func (t *action) forward(writer std.ResponseWriter, request *std.Request) (_ int, err error) {
	ctx := request.Context()
	t.selected, t.done, err = t.selector.Select(ctx, selector.WithNodeFilter(t.filters...))
	if err != nil {
		return std.StatusNotFound, errors.ServiceUnavailable(
			"[Gateway] node not found",
			err.Error(),
		)
	}
	ctx, cancel := context.WithTimeout(ctx, config.Endpoint.Timeout)
	defer cancel()

	request.URL.Scheme = t.selected.Scheme()
	request.URL.Host = t.selected.Address()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return std.StatusBadGateway, err
	}

	request.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(body)), nil
	}

	headers := writer.Header()
	if e := t.PreHandle(ctx, request, &headers); e != nil {
		return e.Status, e.Cause
	}
	// 需要将Host设置为目标微服务地址，否则会导致回环。
	request.Host = request.URL.Host
	tooManyRetry := true
	for i := 1; i <= config.Endpoint.MaxRetry; i++ {
		if err = ctx.Err(); err != nil {
			return std.StatusGatewayTimeout, errors.GatewayTimeout(
				"[Gateway] Timeout while preprocessing the request",
				err.Error(),
			)
		}
		t.response, err = t.transport.RoundTrip(request.Clone(ctx))
		if err == nil {
			tooManyRetry = false
			break
		}
		log.Warnf(
			"Failed to forward request(%d/%d): %s: %+v",
			i, config.Endpoint.MaxRetry,
			request.URL.String(), err,
		)
	}
	if tooManyRetry {
		return std.StatusServiceUnavailable, errors.ServiceUnavailable(
			"[Gateway] service unavailable",
			"maximum retry limit exceeded",
		)
	}

	if t.PostHandle(ctx, t.response); ctx.Err() != nil {
		return std.StatusGatewayTimeout, errors.GatewayTimeout(
			"[Gateway] Timeout while preprocessing the request",
			err.Error(),
		)
	}

	for k, v := range t.response.Header {
		headers[k] = v
	}
	writer.WriteHeader(t.response.StatusCode)

	if t.response.Body == nil {
		return 200, nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(t.response.Body)

	t.sent, err = io.Copy(writer, t.response.Body)
	if err != nil {
		return std.StatusBadGateway, fmt.Errorf("[Gateway] Failed to copy response body from service: %+v", err)
	}

	for k, v := range t.response.Trailer {
		headers[std.TrailerPrefix+k] = v
	}

	return t.response.StatusCode, nil
}
