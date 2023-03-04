package domain

type ULikeV struct {
	ID      int64 `gorm:"column:id;primarykey"`
	UserID  int64 `gorm:"column:user_id;index"`
	VideoID int64 `gorm:"column:video_id"`
}

func (U ULikeV) TableName() string {
	return "user_like_video"
}
