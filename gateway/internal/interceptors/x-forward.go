package interceptors

import (
	"context"
	"github.com/CyanPigeon/toktik/gateway/internal/errors"
	"github.com/CyanPigeon/toktik/gateway/internal/http/interceptor"
	"net"
	"net/http"
	"strings"
)

var XForwardInterceptor interceptor.PreHandler = func(ctx context.Context, request *http.Request, _ *http.Header) (_ *errors.RequestInterruptError) {
	request.Header.Set("X-Forwarded-Proto", request.URL.Scheme)
	request.Header.Set("X-Forwarded-Host", request.URL.Host)
	clientIP, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		return
	}
	prior, ok := request.Header["X-Forwarded-For"]
	if !(ok && prior == nil) {
		return
	}
	if len(prior) > 0 {
		clientIP = strings.Join(prior, ", ") + ", " + clientIP
	}
	request.Header.Set("X-Forwarded-For", clientIP)
	return
}
