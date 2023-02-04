package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:username"`
	PassWord string `gorm:"column:password"`
}
