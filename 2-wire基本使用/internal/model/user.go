package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type User struct {
	Id        int `json:"id" gorm:"primaryKey"`
	Name      string
	Phone     int64
	Birth     *time.Time
	CreatedAt *time.Time
	//软删除
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
