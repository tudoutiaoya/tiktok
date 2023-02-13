package controller

import "tiktok/service"

type Controllers struct {
	*UserController
	*FeedController
	*PublishController
	*FavoriteController
	*CommentController
}

func InitController(services *service.Services) *Controllers {
	return &Controllers{
		UserController:     NewUserController(services.UserService),
		FeedController:     NewFeedController(services.VideoService),
		PublishController:  NewPublishController(services.VideoService),
		FavoriteController: NewFavoriteController(services.VideoService),
		CommentController:  NewCommentController(services.CommentService),
	}
}
