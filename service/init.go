package service

import (
	"tiktok/dao"
	"tiktok/middleware/mredis"
)

type Services struct {
	*UserService
	*VideoService
	*CommentService
	*RelationService
	*MessageService
}

func InitService(databases *dao.Databases) *Services {
	userService := NewUserService(databases.UserDao)
	videoService := NewVideoService(databases.VideoDao, databases.UserDao, userService, databases.RelationDao)
	commentService := NewCommentService(databases.CommentDao, userService)
	relationService := NewRelationService(databases.RelationDao, userService, databases.MessageDao, mredis.RedisClient, databases.UserDao)
	messageService := NewMessageService(databases.MessageDao, mredis.RedisClient)
	return &Services{
		UserService:     userService,
		VideoService:    videoService,
		CommentService:  commentService,
		RelationService: relationService,
		MessageService:  messageService,
	}
}
