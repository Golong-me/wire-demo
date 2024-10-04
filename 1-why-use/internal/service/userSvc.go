package service

import (
	"wire-why/internal/dao"
	"wire-why/internal/model"
)

type UserService struct {
}

var userDao = dao.UserDao{}

//product Service
//redis

func (receiver UserService) GetUserInfo() *model.User {
	return userDao.GetUserInfo()
}
