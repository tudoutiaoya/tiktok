package service

import "tiktok/dao"

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{
		userDao: userDao,
	}
}

func (s *UserService) Create() error {
	err := s.userDao.Create()
	if err != nil {
		return err
	}
	return nil
}
