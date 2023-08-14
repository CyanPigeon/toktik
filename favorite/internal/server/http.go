package server

import (
	fav "favorite/api/toktik/favorite"
	"favorite/internal/conf"
	favoriteService "favorite/internal/service/favorite"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, likeService *favoriteService.LikeActionService, likeListService *favoriteService.LikeListService, logger log.Logger) *http.Server {
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
	fav.RegisterLikeActionHTTPServer(srv, likeService)
	fav.RegisterLikeListHTTPServer(srv, likeListService)
	return srv
}
