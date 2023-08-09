package main

import (
	"context"
	"fmt"
	"github.com/CyanPigeon/toktik/gateway/internal/config"
	Discovery "github.com/CyanPigeon/toktik/gateway/internal/discovery"
	"github.com/CyanPigeon/toktik/gateway/internal/http"
	"github.com/CyanPigeon/toktik/gateway/internal/http/endpoint"
	Proxy "github.com/CyanPigeon/toktik/gateway/internal/proxy"
	Router "github.com/CyanPigeon/toktik/gateway/internal/router"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/p2c"
	Etcd "go.etcd.io/etcd/client/v3"
)

func initialization(ctx context.Context) (server http.Server, err error) {
	var client *Etcd.Client

	selector.SetGlobalSelector(p2c.NewBuilder())
	client, err = Etcd.New(Etcd.Config{Endpoints: config.Discovery.Endpoints})
	if err != nil {
		return nil, fmt.Errorf("etcd client initialization failed: %+v", err)
	}
	discovery := Discovery.New(client)
	transport := http.NewTransport()
	router := Router.New(ctx, discovery, endpoint.NewEndpointFactory(transport))
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
