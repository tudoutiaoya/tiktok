package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"tiktok/config"
	"tiktok/domain"
)

// Databases 结构体
type Databases struct {
	*UserDao
	*VideoDao
	*CommentDao
	*RelationDao
	*MessageDao
}

func InitDao(config *config.Configuration) *Databases {
	db, err := gorm.Open(mysql.Open(config.DatabaseSettings.DatabaseURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("不能连接到数据库 : %s", err.Error()))
	}
	// 与数据库对应
	err = db.AutoMigrate(&domain.Video{}, &domain.User{}, &domain.ULikeV{}, &domain.Comment{}, &domain.Relation{}, &domain.Message{})
	if err != nil {
		panic(fmt.Sprintf("domain转移数据库失败 : %s", err.Error()))
	}
	return &Databases{
		UserDao:     NewUserDao(db),
		VideoDao:    NewVideoDao(db),
		CommentDao:  NewCommentDao(db),
		RelationDao: NewRelationDao(db),
		MessageDao:  NewMessageDao(db),
	}
}
