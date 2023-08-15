package server

import (
	v1 "relation/api/toktik/relation"
	"relation/internal/conf"
	"relation/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, action *service.FollowActionService, followList *service.FollowListService,
	followerList *service.FollowerListService, friendList *service.FriendListService, logger log.Logger) *http.Server {
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
	v1.RegisterFollowActionHTTPServer(srv, action)
	v1.RegisterFollowListHTTPServer(srv, followList)
	v1.RegisterFollowerListHTTPServer(srv, followerList)
	v1.RegisterFriendListHTTPServer(srv, friendList)

	return srv
}
