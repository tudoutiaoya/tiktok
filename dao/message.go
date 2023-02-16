package dao

import (
	"gorm.io/gorm"
	"tiktok/domain"
)

type MessageDao struct {
	db *gorm.DB
}

func (d *MessageDao) MessageAction(message domain.Message) error {
	return d.db.Create(&message).Error
}

func (d *MessageDao) MessageChat(userID int64, toUserID int64) ([]domain.Message, error) {
	var chatMessage []domain.Message
	err := d.db.Where("user_id = ? and to_user_id = ? and if_read = 0", userID, toUserID).Find(&chatMessage).Error
	d.db.Model(&chatMessage).Update("if_read", 1)
	return chatMessage, err
}

func NewMessageDao(db *gorm.DB) *MessageDao {
	return &MessageDao{db: db}
}
