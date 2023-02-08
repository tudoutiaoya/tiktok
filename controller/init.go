package controller

import "tiktok/service"

type Controllers struct {
	UserController *UserController
	FeedController *FeedController
}

func InitController(services *service.Services) *Controllers {
	return &Controllers{
		UserController: NewUserController(services.UserService),
		FeedController: NewFeedController(services.FeedService),
	}
}
