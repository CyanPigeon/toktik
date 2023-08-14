package main

import (
	"context"
	"fmt"
	"github.com/CyanPigeon/toktik/gateway/internal/config"
	"github.com/CyanPigeon/toktik/gateway/internal/http"
	"github.com/CyanPigeon/toktik/gateway/internal/http/endpoint"
	"github.com/CyanPigeon/toktik/gateway/internal/interceptors"
	Proxy "github.com/CyanPigeon/toktik/gateway/internal/proxy"
	Router "github.com/CyanPigeon/toktik/gateway/internal/router"
	Discovery "github.com/CyanPigeon/toktik/middleware/discovery"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/p2c"
	Consul "github.com/hashicorp/consul/api"
)

func initialization(ctx context.Context) (server http.Server, err error) {
	selector.SetGlobalSelector(p2c.NewBuilder())
	discovery, err := Discovery.New(Consul.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("etcd client initialization failed: %+v", err)
	}
	transport := http.NewTransport()
	router := Router.New(
		ctx,
		endpoint.NewEndpointFactory(transport),
		Router.WithPreInterceptors(
			interceptors.XForwardInterceptor,
			interceptors.Authorization,
		),
	)
	server, err = Proxy.NewServer(
		ctx,
		discovery,
		router,
	)
	if err != nil {
		return nil, fmt.Errorf("reverse proxy server initialization failed: %+v", err)
	}
	return
}

func main() {
	ctx := context.Background()

	server, err := initialization(ctx)
	if err != nil {
		panic(err)
	}

	if err = server.Listen(config.ServerListenAddress()); err != nil {
		panic(err)
	}
}
