package server

import (
	v1 "github.com/CyanPigeon/toktik/api/demo/v1"
	"github.com/CyanPigeon/toktik/app/demo/internal/conf"
	"github.com/CyanPigeon/toktik/app/demo/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer 创建一个gRPC服务。
// TODO 需要将第二个参数改为服务层的ApiService
func NewGRPCServer(c *conf.Server, serv *service.DemoService, logger log.Logger) *grpc.Server {
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

	// TODO 在此调用API的gRPC服务注册函数。
	// TODO 一般来说，gRPC服务注册函数为`Register${ApiName}Server`，其中`${ApiName}`为Api名。
	// TODO 如果有多个gRPC服务，**需要在此将它们全部注册**。
	// TODO 请注意：如果有多个gRPC服务，说明你有多个Service，因此你需要增加方法的参数，告诉wire你需要更多的Service。
	// TODO 下一行为示例代码。
	v1.RegisterDemoServer(srv, serv)

	return srv
}
