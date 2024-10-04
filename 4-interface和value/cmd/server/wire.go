//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-interface/internal/conf"
	"wire-interface/internal/controller"
)

func InitCtrl() *controller.UserController {
	wire.Build(
		conf.NewGorm,
		controller.CtrlProvider,
	)
	return nil
}
