package service

import "tiktok/dao"

type Services struct {
	UserService *UserService
}

func InitService(databases *dao.Databases) *Services {
	return &Services{
		UserService: NewUserService(databases.UseDao),
	}
}
