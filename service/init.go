package service

import "tiktok/dao"

type Services struct {
	UserService *UserService
	FeedService *FeedService
}

func InitService(databases *dao.Databases) *Services {
	return &Services{
		UserService: NewUserService(databases.UseDao),
		FeedService: NewFeedService(databases.FeedDao),
	}
}
