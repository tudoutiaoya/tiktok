package domain

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	PlayUrl       string `gorm:"type:varchar(255);not null"`
	CoverUrl      string `gorm:"type:varchar(255)"`
	FavoriteCount int64  `gorm:"default:0"`
	CommentCount  int64  `gorm:"default:0"`
	Title         string `gorm:"type:varchar(255);not null"`
	AuthorID      int64  `gorm:"not null"`
}

func (v *Video) TableName() string {
	return "video"
}
