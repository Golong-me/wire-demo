//go:build wireinject
// +build wireinject

package main

import (
	"gin-wire/internal/conf"
	"gin-wire/internal/controller"
	"github.com/google/wire"
)

func wireApp() *conf.HttpServer {
	wire.Build(
		wire.Struct(new(conf.HttpServer), "*"),
		conf.ZapProvider,
		conf.NewGin,
		conf.NewGorm,
		controller.CtrlProvider,
	)
	return nil
}
