package controller

import (
	"gin-wire/internal/dao"
	"gin-wire/internal/service"
	"github.com/google/wire"
)

var CtrlProvider = wire.NewSet(UserProvider)

var UserProvider = wire.NewSet(
	NewUserController,
	wire.Bind(new(service.IUserService), new(service.UserService)),
	wire.Struct(new(service.UserService), "UserDao"),
	wire.Struct(new(dao.UserDao), "Db"),
)
