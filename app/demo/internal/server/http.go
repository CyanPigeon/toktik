package server

import (
	v1 "github.com/CyanPigeon/toktik/api/demo/v1"
	"github.com/CyanPigeon/toktik/app/demo/internal/conf"
	"github.com/CyanPigeon/toktik/app/demo/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer 创建一个HTTP服务。
// TODO 需要将第二个参数改为服务层的ApiService
func NewHTTPServer(c *conf.Server, serv *service.DemoService, logger log.Logger) *http.Server {
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

	// TODO 在此调用API的HTTP服务注册函数。
	// TODO 一般来说，HTTP服务注册函数为`Register${ApiName}HTTPServer`，其中`${ApiName}`为Api名。
	// TODO 如果有多个HTTP服务，**需要在此将它们全部注册**。
	// TODO 请注意：如果有多个HTTP服务，说明你有多个Service，因此你需要增加方法的参数，告诉wire你需要更多的Service。
	// TODO 下一行为示例代码。
	v1.RegisterDemoHTTPServer(srv, serv)
	return srv
}
