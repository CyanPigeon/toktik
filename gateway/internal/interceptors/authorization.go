package interceptors

import (
	"context"
	"errors"
	toktik "github.com/CyanPigeon/toktik/gateway/internal/errors"
	"github.com/CyanPigeon/toktik/gateway/internal/http/interceptor"
	"github.com/CyanPigeon/toktik/middleware"
	"net/http"
	"net/url"
	"strings"
)

var whitelist = []string{
	"/douyin/feed",
	"/douyin/user/",
}

// needAuthorization 用于检查url是否需要鉴权。
// true=需要鉴权；
// false=不需要鉴权。
func needAuthorization(url *url.URL) bool {
	for _, pattern := range whitelist {
		if strings.HasPrefix(url.Path, pattern) {
			return false
		}
	}
	return true
}

var Authorization interceptor.PreHandler = func(_ context.Context, request *http.Request, header *http.Header) *toktik.RequestInterruptError {
	if !needAuthorization(request.URL) {
		return nil
	}
	params := request.URL.Query()

	if !params.Has("token") {
		return &toktik.RequestInterruptError{
			Cause:  errors.New("authorize failed: token not found"),
			Status: http.StatusUnauthorized,
		}
	}

	success, _, err := middleware.ValidateToken(params.Get("token"))
	if !success && err == nil {
		err = errors.New("authorize failed: invalid token")
	}
	if err != nil {
		return &toktik.RequestInterruptError{
			Cause:  err,
			Status: http.StatusUnauthorized,
		}
	}

	return nil
}
