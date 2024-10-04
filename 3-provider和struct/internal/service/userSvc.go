package service

import (
	"wire-provider/internal/dao"
	"wire-provider/internal/model"
)

type UserService struct {
	UserDao *dao.UserDao
}

func NewUserService(UserDao *dao.UserDao) *UserService {
	return &UserService{UserDao: UserDao}
}

func (receiver UserService) GetUserInfo() *model.User {
	return receiver.UserDao.GetUserInfo()
}
