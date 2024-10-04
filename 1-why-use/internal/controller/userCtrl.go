package controller

import (
	"wire-why/internal/conf"
	"wire-why/internal/service"
)

type UserController struct {
}

var UserService = service.UserService{}

func (receiver *UserController) GetUserInfo() {
	info := UserService.GetUserInfo()

	//打印日志
	conf.Logger.Sugar().Info(info)
}
