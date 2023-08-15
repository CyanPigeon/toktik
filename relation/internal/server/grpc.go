package server

import (
	v1 "relation/api/toktik/relation"
	"relation/internal/conf"
	"relation/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, action *service.FollowActionService, followList *service.FollowListService,
	followerList *service.FollowerListService, friendList *service.FriendListService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterFollowActionServer(srv, action)
	v1.RegisterFollowListServer(srv, followList)
	v1.RegisterFollowerListServer(srv, followerList)
	v1.RegisterFriendListServer(srv, friendList)

	return srv
}
