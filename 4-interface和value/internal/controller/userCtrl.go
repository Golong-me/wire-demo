package controller

import (
	"log"
	"wire-interface/internal/service"
)

//设计原则
//依赖倒置

type UserController struct {
	//维护接口
	UserService service.IUserService
}

func (receiver *UserController) GetUserInfo() {
	info := receiver.UserService.GetUserInfo()

	//打印日志
	log.Println(info)
}
