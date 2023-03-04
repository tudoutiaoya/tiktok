package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName      string `gorm:"column:username;type:varchar(255);not null"`
	PassWord      string `gorm:"column:password;type:varchar(255);not null"`
	FollowCount   int64  `gorm:"default:0"`
	FollowerCount int64  `gorm:"default:0"`
	Avatar        string `gorm:"column:avatar"`
	WorkCount     int64  `gorm:"column:work_count"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
}

func (u User) TableName() string {
	return "user"
}
