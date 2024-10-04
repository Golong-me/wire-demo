package controller

import (
	"log"
	"wire-base-use/internal/service"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(UserService *service.UserService) *UserController {
	return &UserController{UserService: UserService}
}

//func NewUserController(db *gorm.DB) *UserController {
//	return &UserController{
//		UserService: &service.UserService{
//			UserDao: &dao.UserDao{
//				//那我们其他的productService呢?
//				//此时就要做好Gorm的单例处理了
//				//Db: conf.NewGorm(),
//
//				//也可以放到main里，然后传递过来
//				//推荐放到main里
//				Db: db,
//			},
//		},
//	}
//}

func (receiver *UserController) GetUserInfo() {
	info := receiver.UserService.GetUserInfo()

	//打印日志
	log.Println(info)
}
