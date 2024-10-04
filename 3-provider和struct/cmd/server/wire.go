//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-provider/internal/conf"
	"wire-provider/internal/controller"
)

// wire ./...
//可以生成当前目录以及当前目录下所有的子孙目录和文件的Build
//wire无法分析多个build中相同的对象

func InitCtrl() *controller.UserController {
	wire.Build(
		conf.NewGorm,
		controller.CtrlProvider,
	)
	return nil
}
