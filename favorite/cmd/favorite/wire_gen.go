// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"favorite/internal/biz/favorite"
	"favorite/internal/conf"
	"favorite/internal/data"
	"favorite/internal/server"
	"favorite/internal/service/favorite"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db, err := data.NewGormDB(confData)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(confData, db, logger)
	if err != nil {
		return nil, nil, err
	}
	favoriteServiceBiz := favorite.NewFavoriteServiceBiz(dataData)
	likeActionService := service.NewLikeActionService(favoriteServiceBiz)
	likeListService := service.NewLikeListService(favoriteServiceBiz)
	grpcServer := server.NewGRPCServer(confServer, likeActionService, likeListService, logger)
	httpServer := server.NewHTTPServer(confServer, likeActionService, likeListService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
