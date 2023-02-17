package api

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
	"tiktok/middleware/mjwt"
)

func InitRouter(r *gin.Engine, web *controller.Controllers) {
	apiRouter := r.Group("/douyin")
	// 基础接口
	// 视频流
	apiRouter.GET("/feed", web.FeedController.GetFeed)

	// 获取登录用户信息
	apiRouter.GET("/user/", mjwt.JWTAuthMiddleware(), web.UserController.CurrentUser)

	// 用户注册
	apiRouter.POST("/user/register/", web.UserController.Register)

	// 用户登录
	apiRouter.POST("/user/login/", web.UserController.Login)

	// 视频投稿
	apiRouter.POST("/publish/action/", mjwt.JWTAuthMiddleware(), web.PublishController.Publish)

	// 发布列表
	apiRouter.GET("/publish/list/", mjwt.JWTAuthMiddleware(), web.PublishController.PublishList)

	// 点赞/取消
	apiRouter.POST("/favorite/action/", mjwt.JWTAuthMiddleware(), web.FavoriteController.Action)

	// 喜欢列表
	apiRouter.GET("/favorite/list/", mjwt.JWTAuthMiddleware(), web.FavoriteController.LikeList)

	// 评论
	apiRouter.POST("/comment/action/", mjwt.JWTAuthMiddleware(), web.CommentController.CommentAction)

	// 评论列表
	apiRouter.GET("/comment/list/", web.CommentController.CommentList)

	// 关注操作
	apiRouter.POST("/relation/action/", mjwt.JWTAuthMiddleware(), web.RelationController.RelationAction)

	// 关注列表
	apiRouter.GET("/relation/follow/list/", mjwt.JWTAuthMiddleware(), web.RelationController.FollowList)

	// 粉丝列表
	apiRouter.GET("/relation/follower/list/", mjwt.JWTAuthMiddleware(), web.RelationController.FollowerList)

	// 朋友列表
	apiRouter.GET("/relation/friend/list/", mjwt.JWTAuthMiddleware(), web.RelationController.FriendList)

	// 发送消息
	apiRouter.POST("/message/action/", mjwt.JWTAuthMiddleware(), web.MessageController.MessageAction)

	// 聊天记录
	apiRouter.GET("/message/chat/", mjwt.JWTAuthMiddleware(), web.MessageController.MessageChat)

}
