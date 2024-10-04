//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-base-use/internal/conf"
	"wire-base-use/internal/controller"
	"wire-base-use/internal/dao"
	"wire-base-use/internal/service"
)

func InitCtrl() *controller.UserController {
	wire.Build(
		conf.NewGorm,
		controller.NewUserController,
		service.NewUserService,
		dao.NewUserDao,
	)
	return nil
}
