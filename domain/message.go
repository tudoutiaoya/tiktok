package domain

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        uint   `gorm:"primarykey"`
	Content   string `gorm:"column:content;not null"`
	UserID    int64  `gorm:"column:user_id;not null;index"`
	ToUserID  int64  `gorm:"column:to_user_id;not null;index"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m Message) TableName() string {
	return "message"
}
