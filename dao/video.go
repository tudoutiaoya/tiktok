package dao

import (
	"gorm.io/gorm"
	"tiktok/domain"
	"time"
)

type VideoDao struct {
	db *gorm.DB
}

func NewVideoDao(db *gorm.DB) *VideoDao {
	return &VideoDao{db: db}
}

func (dao *VideoDao) GetFeed(limit int, latestTime int64) ([]domain.Video, error) {
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

func (dao *VideoDao) SaveVideo(video domain.Video) error {
	return dao.db.Create(&video).Error
}

func (dao *VideoDao) GetUserPublishList(id int64) ([]domain.Video, error) {
	var videos []domain.Video
	err := dao.db.Where("author_id = ?", id).Find(&videos).Error
	return videos, err
}

func (dao *VideoDao) GetVideoById(id int64) (domain.Video, error) {
	var video domain.Video
	err := dao.db.Where("id = ?", id).Find(&video).Error
	return video, err
}

func (d *VideoDao) Like(uLikeV domain.ULikeV) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&uLikeV).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		if err := tx.Model(&domain.Video{}).
			Where("id = ?", uLikeV.VideoID).
			Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).
			Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		return nil
	})
}

func (d *VideoDao) DisLike(uLikeV domain.ULikeV) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Where("user_id = ?", uLikeV.UserID).Delete(&uLikeV).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		if err := tx.Model(&domain.Video{}).
			Where("id = ?", uLikeV.VideoID).
			Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).
			Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		return nil
	})
}

// LikeListVideoIDs 获取作者喜欢视频id
func (d *VideoDao) LikeListVideoIDs(userID int64) ([]int64, error) {
	videoIDs := make([]int64, 0)
	err := d.db.Model(&domain.ULikeV{}).Where("user_id = ?", userID).Select("video_id").Find(&videoIDs).Error
	return videoIDs, err
}

func (d *VideoDao) GetVideoListByIDs(videoIDs []int64) ([]domain.Video, error) {
	var videos []domain.Video
	err := d.db.Where("id in ?", videoIDs).Find(&videos).Error
	return videos, err
}

// IsLike 是否已经点过赞
func (dao *VideoDao) IsLike(userID int64, videoID int64) (bool, error) {
	var count int64
	err := dao.db.Model(&domain.ULikeV{}).Where("user_id = ? and video_id = ?", userID, videoID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
