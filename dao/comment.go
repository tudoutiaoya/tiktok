package dao

import (
	"gorm.io/gorm"
	"tiktok/domain"
)

type CommentDao struct {
	db *gorm.DB
}

func (d *CommentDao) CreatComment(comment *domain.Comment) error {
	return d.db.Create(comment).Error
}

func (d *CommentDao) DeleteComment(commentID string) error {
	return d.db.Where("id = ?", commentID).Delete(&domain.Comment{}).Error
}

func (d *CommentDao) GetCommentList(videoID int64) ([]domain.Comment, error) {
	var comments []domain.Comment
	err := d.db.Where("video_id = ?", videoID).Order("created_at desc").Find(&comments).Error
	return comments, err
}

func NewCommentDao(db *gorm.DB) *CommentDao {
	return &CommentDao{db: db}
}
