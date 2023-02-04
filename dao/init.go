package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/config"
)

// Databases 结构体
type Databases struct {
	UseDao *UserDao
}

func InitDao(config *config.Configuration) *Databases {
	db, err := gorm.Open(mysql.Open(config.DatabaseSettings.DatabaseURI), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("不能连接到数据库 : %s", err.Error()))
	}
	return &Databases{
		UseDao: NewUserDao(db),
	}
}
