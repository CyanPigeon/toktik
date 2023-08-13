package interceptor

import (
	"context"
	toktik "github.com/CyanPigeon/toktik/gateway/internal/errors"
	kratos "github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

type PreHandler func(ctx context.Context, request *http.Request, header *http.Header) *toktik.RequestInterruptError
type PostHandler func(ctx context.Context, response *http.Response)

type Interceptors interface {
	PreHandle(ctx context.Context, request *http.Request, header *http.Header) *toktik.RequestInterruptError
	PostHandle(ctx context.Context, response *http.Response)
}

type DefaultInterceptors struct {
	Pre  []PreHandler
	Post []PostHandler
}

func (t *DefaultInterceptors) PreHandle(ctx context.Context, request *http.Request, header *http.Header) *toktik.RequestInterruptError {
	for i := len(t.Pre) - 1; i >= 0; i-- {
		if err := ctx.Err(); err != nil {
			return &toktik.RequestInterruptError{
				Cause: kratos.GatewayTimeout(
					"[Gateway] Timeout while preprocessing the request",
					err.Error(),
				),
				Status: http.StatusGatewayTimeout,
			}
		} else if err := t.Pre[i](ctx, request, header); err != nil {
			return err
		}
	}
	return nil
}

func (t *DefaultInterceptors) PostHandle(ctx context.Context, response *http.Response) {
	for i := len(t.Post) - 1; i >= 0; i-- {
		if ctx.Err() != nil {
			return
		}
		t.Post[i](ctx, response)
	}
}
