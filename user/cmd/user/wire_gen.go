// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"
	"user/internal/conf"
	"user/internal/data"
	"user/internal/server"
	"user/internal/service/user"
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
	userServiceBizImpl := biz.NewUserServiceImpl(dataData)
	loginService := user.NewUserLoginService(userServiceBizImpl)
	registerService := user.NewUserRegisterService(userServiceBizImpl)
	infoService := user.NewUserInfoService(userServiceBizImpl)
	grpcServer := server.NewGRPCServer(confServer, loginService, registerService, infoService, logger)
	httpServer := server.NewHTTPServer(confServer, loginService, registerService, infoService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
