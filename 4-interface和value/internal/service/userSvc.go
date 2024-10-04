package service

import (
	"wire-interface/internal/dao"
	"wire-interface/internal/model"
)

type IUserService interface {
	GetUserInfo() *model.User
}

type UserService struct {
	UserDao *dao.UserDao
}

/*
	func NewUserService(UserDao *dao.UserDao) *UserService {
		return &UserService{UserDao: UserDao}
	}
*/
func (receiver UserService) GetUserInfo() *model.User {
	return receiver.UserDao.GetUserInfo()
}
