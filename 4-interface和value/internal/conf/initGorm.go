package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//核心的就已经到这里了
//wire的使用
//gin的案例
//gin优雅的使用wire

func NewGorm() *gorm.DB {
	url := "root:root@tcp(192.168.80.128:3306)/myTest?charset=utf8mb4&parseTime=True&loc=Local"
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	Db, _ := gorm.Open(mysql.Open(url), config)

	return Db
}
