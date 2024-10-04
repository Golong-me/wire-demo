package controller

import (
	"github.com/google/wire"
	"wire-provider/internal/dao"
	"wire-provider/internal/service"
)

var CtrlProvider = wire.NewSet(UserProvider)

var UserProvider = wire.NewSet(
	wire.Struct(new(UserController), "*"),
	wire.Struct(new(service.UserService), "*"),
	wire.Struct(new(dao.UserDao), "*"))
