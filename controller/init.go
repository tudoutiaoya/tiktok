package controller

import "tiktok/service"

type Controllers struct {
	UserController *UserController
}

func InitController(service *service.Services) *Controllers {
	return &Controllers{
		UserController: NewUserController(service.UserService),
	}
}
