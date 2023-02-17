package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName      string `gorm:"column:username;type:varchar(255);not null"`
	PassWord      string `gorm:"column:password;type:varchar(255);not null"`
	FollowCount   int64  `gorm:"default:0"`
	FollowerCount int64  `gorm:"default:0"`
	Avatar        string `gorm:"column:avatar"`
}

func (u User) TableName() string {
	return "user"
}
