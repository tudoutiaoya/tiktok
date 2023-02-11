package dao

import (
	"gorm.io/gorm"
	"tiktok/domain"
	"time"
)

type FeedDao struct {
	db *gorm.DB
}

func NewFeedDao(db *gorm.DB) *FeedDao {
	return &FeedDao{db: db}
}

func (dao *FeedDao) GetFeed(limit int, latestTime int64) ([]domain.Video, error) {
	videos := make([]domain.Video, 0)
	// select * from video where updated_at < . order by updated_at limit .
	//err := dao.db.Where("updated_at < ?", time.Unix(latestTime, 0).String()).Order("updated_at desc").Limit(limit).Find(&videos).Error
	//err := dao.db.Order("updated_at desc").Find(&videos).Error
	err := dao.db.Where("updated_at < ?", time.UnixMilli(latestTime).String()).Order("updated_at desc").Limit(limit).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (dao *FeedDao) SaveVideo(video domain.Video) error {
	return dao.db.Create(&video).Error
}

func (dao *FeedDao) GetUserPublishList(id int64) ([]domain.Video, error) {
	var videos []domain.Video
	err := dao.db.Where("author_id = ?", id).Find(&videos).Error
	return videos, err
}
