package controller

import (
	"github.com/google/wire"
	"wire-interface/internal/dao"
	"wire-interface/internal/service"
)

var CtrlProvider = wire.NewSet(UserProvider)

var UserProvider = wire.NewSet(
	//绑定接口
	//参数1 接口类型
	//参数2 实现类
	wire.Bind(new(service.IUserService), new(service.UserService)),
	wire.Struct(new(UserController), "UserService"),
	wire.Struct(new(service.UserService), "UserDao"),
	wire.Struct(new(dao.UserDao), "Db"),
)
