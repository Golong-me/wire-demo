package dao

import (
	"wire-why/internal/conf"
	"wire-why/internal/model"
)

type UserDao struct {
}

func (receiver UserDao) GetUserInfo() *model.User {
	var user model.User
	conf.Db.Model(&model.User{}).Find(&user, 1)
	return &user
}
