package service

import "tiktok/dao"

type Services struct {
	*UserService
	*VideoService
	*CommentService
}

func InitService(databases *dao.Databases) *Services {
	userService := NewUserService(databases.UserDao)
	videoService := NewVideoService(databases.VideoDao, databases.UserDao, userService)
	commentService := NewCommentService(databases.CommentDao, userService)
	return &Services{
		UserService:    userService,
		VideoService:   videoService,
		CommentService: commentService,
	}
}
