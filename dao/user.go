package dao

import (
	"gorm.io/gorm"
	"tiktok/domain"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (u *UserDao) SelectCount(username string) int64 {
	var count int64
	u.db.Model(&domain.User{}).Where("username = ?", username).Count(&count)
	return count
}

func (u *UserDao) CreatUse(user *domain.User) error {
	return u.db.Create(user).Error
}

func (u *UserDao) GetUserByUserName(username string) (domain.User, error) {
	var user domain.User
	err := u.db.Where("username = ?", username).Find(&user).Error
	return user, err
}

func (u *UserDao) GetUserById(id int64) (domain.User, error) {
	var user domain.User
	err := u.db.Where("id = ?", id).Find(&user).Error
	return user, err
}

func (u *UserDao) GetUserByVideoId(id int64) (domain.User, error) {
	var user domain.User
	err := u.db.Where("id = ?").Find(&user).Error
	return user, err
}

func (u *UserDao) GetUserByIds(iDs []string) ([]domain.User, error) {
	var users []domain.User
	u.db.Where("id in ?", iDs).Find(&users)
	return users, nil
}
