package dao

import (
	"gorm.io/gorm"
	"tiktok/domain"
)

type RelationDao struct {
	db *gorm.DB
}

func (d *RelationDao) Follow(userID int64, toUserID int64) error {
	var relation = domain.Relation{
		UserID:   userID,
		ToUserID: toUserID,
	}
	return d.db.Create(&relation).Error
}

func (d *RelationDao) UnFollow(userID int64, toUserID int64) error {
	return d.db.Where("user_id = ? and to_user_id = ?", userID, toUserID).Delete(&domain.Relation{}).Error
}

func (d *RelationDao) IsHas(userID int64, toUserID int64) (bool, error) {
	var count int64
	err := d.db.Model(&domain.Relation{}).Where("user_id = ? and to_user_id = ?", userID, toUserID).Count(&count).Error
	return count > 0, err
}

// FollowList 关注列表
func (d *RelationDao) FollowList(userID int64) ([]domain.User, error) {
	var userIDs []int64
	subQuery := d.db.Model(&domain.Relation{}).Select("to_user_id").Where("user_id = ?", userID).Find(&userIDs)
	var users []domain.User
	err := d.db.Where("id in (?)", subQuery).Find(&users).Error
	return users, err
}

// FansList 粉丝列表
func (d *RelationDao) FansList(userID int64) ([]domain.User, error) {
	var userIDs []int64
	subQuery := d.db.Model(&domain.Relation{}).Select("user_id").Where("to_user_id = ?", userID).Find(&userIDs)
	var users []domain.User
	err := d.db.Where("id in (?)", subQuery).Find(&users).Error
	return users, err
}

// FriendList 朋友列表
func (d *RelationDao) FriendList(userID int64) ([]domain.User, error) {
	var userIDs []int64
	// 使用两个子查询
	// TODO 再优化，关系放在redis当中
	subQuery1 := d.db.Model(&domain.Relation{}).Select("user_id").Where("to_user_id = ?", userID).Find(&userIDs)
	// SELECT * FROM `user` WHERE id in
	//   (SELECT `to_user_id` FROM `relation` WHERE to_user_id in
	//           (SELECT `user_id` FROM `relation` WHERE to_user_id = 2) and user_id = 2);
	subQuery2 := d.db.Model(&domain.Relation{}).Select("to_user_id").Where("to_user_id in (?) and user_id = ?", subQuery1, userID).Find(&userIDs)
	var users []domain.User
	err := d.db.Where("id in (?)", subQuery2).Find(&users).Error
	return users, err
}

func NewRelationDao(db *gorm.DB) *RelationDao {
	return &RelationDao{db: db}
}
