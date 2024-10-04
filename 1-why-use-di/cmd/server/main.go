package main

import (
	"wire-why/internal/conf"
	"wire-why/internal/controller"
)

func main() {
	db := conf.NewGorm()

	userController := controller.NewUserController(db)
	userController.GetUserInfo()
}
