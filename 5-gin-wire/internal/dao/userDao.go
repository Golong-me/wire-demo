package dao

import (
	"gin-wire/internal/model"
	"gorm.io/gorm"
)

// 这里的db要注入进去，而非使用全局的
type UserDao struct {
	Db *gorm.DB
}

// 构造函数
func NewUserDao(Db *gorm.DB) *UserDao {
	return &UserDao{Db: Db}
}

func (receiver UserDao) GetUserInfo() *model.User {
	var user model.User
	receiver.Db.Model(&model.User{}).Find(&user, 1)
	return &user
}
