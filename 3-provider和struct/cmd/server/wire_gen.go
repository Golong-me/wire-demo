// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"wire-provider/internal/conf"
	"wire-provider/internal/controller"
	"wire-provider/internal/dao"
	"wire-provider/internal/service"
)

// Injectors from wire.go:

func InitCtrl() *controller.UserController {
	db := conf.NewGorm()
	userDao := &dao.UserDao{
		Db: db,
	}
	userService := &service.UserService{
		UserDao: userDao,
	}
	userController := &controller.UserController{
		UserService: userService,
	}
	return userController
}
