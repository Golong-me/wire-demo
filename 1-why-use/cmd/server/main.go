package main

import (
	"wire-why/config"
	"wire-why/internal/controller"
)

func main() {
	//初始化所有配置
	config.InitAllConfig()

	userController := controller.UserController{}
	userController.GetUserInfo()
}
