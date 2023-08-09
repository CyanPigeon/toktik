package http

import (
	"github.com/CyanPigeon/toktik/gateway/internal/http/interceptor"
	"github.com/go-kratos/kratos/v2/selector"
	"io"
	"net/http"
)

type Endpoint interface {
	Path() string
	ServeHTTPEx(writer http.ResponseWriter, request *http.Request, interceptors interceptor.Interceptors)
	selector.Rebalancer
	io.Closer
}

type Server interface {
	Listen(host string) error
	io.Closer
}

type ErrorWriter interface {
	Write(writer http.ResponseWriter, status int, err error)
}
