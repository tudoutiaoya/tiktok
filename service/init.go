package service

import "tiktok/dao"

type Services struct {
	UserService *UserService
	FeedService *FeedService
}

func InitService(databases *dao.Databases) *Services {
	userService := NewUserService(databases.UseDao)
	feedService := NewFeedService(databases.FeedDao, databases.UseDao, userService)
	return &Services{
		UserService: userService,
		FeedService: feedService,
	}
}
