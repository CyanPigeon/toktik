//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/CyanPigeon/toktik/app/demo/internal/conf"
	"github.com/CyanPigeon/toktik/app/demo/internal/data"
	"github.com/CyanPigeon/toktik/app/demo/internal/server"
	"github.com/CyanPigeon/toktik/app/demo/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		server.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}
