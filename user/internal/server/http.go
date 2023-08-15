package server

import (
	"user/api/toktik/user"
	"user/internal/conf"
	user2 "user/internal/service/user"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userLoginSrv *user2.UserLoginService, userRegisterSrv *user2.UserRegisterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	user.RegisterUserLoginHTTPServer(srv, userLoginSrv)
	user.RegisterUserRegisterHTTPServer(srv, userRegisterSrv)
	return srv
}
