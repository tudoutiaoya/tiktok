package domain

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	VideoID int64  `gorm:"column:video_id;not null"`
	UserID  int64  `gorm:"column:user_id;not null"`
	Content string `gorm:"type:varchar(255);not null"`
}

func (c Comment) TableName() string {
	return "comment"
}
