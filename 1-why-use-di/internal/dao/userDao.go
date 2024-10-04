package dao

import (
	"gorm.io/gorm"
	"wire-why/internal/model"
)

// 这里的db要注入进去，而非使用全局的
type UserDao struct {
	Db *gorm.DB
}

func (receiver UserDao) GetUserInfo() *model.User {
	var user model.User
	receiver.Db.Model(&model.User{}).Find(&user, 1)
	return &user
}
