package controller

import (
	"gin-wire/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
)

//设计原则
//依赖倒置

type UserController struct {
	//维护接口
	UserService service.IUserService
	Logger      *zap.Logger
}

func NewUserController(engine *gin.Engine, UserSvc service.IUserService, Logger *zap.Logger) *UserController {
	ctl := &UserController{UserService: UserSvc, Logger: Logger}
	group := engine.Group("/user")
	group.GET("/info/:id", ctl.GetUserInfo())
	return ctl
}

func (receiver *UserController) GetUserInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		info := receiver.UserService.GetUserInfo()
		//打印日志
		log.Println(info)
	}
}
