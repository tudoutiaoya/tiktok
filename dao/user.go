package dao

import (
	"gorm.io/gorm"
	"tiktok/domain"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	db.AutoMigrate(&domain.User{})
	return &UserDao{
		db: db,
	}
}

func (u *UserDao) Create() error {
	tx := u.db.Create(&domain.User{UserName: "haha", PassWord: "123456"})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
