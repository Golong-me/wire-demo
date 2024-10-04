package service

import (
	"wire-why/internal/dao"
	"wire-why/internal/model"
)

type UserService struct {
	UserDao *dao.UserDao
}

func (receiver UserService) GetUserInfo() *model.User {
	return receiver.UserDao.GetUserInfo()
}
