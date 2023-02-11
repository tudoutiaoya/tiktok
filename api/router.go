package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/controller"
	"tiktok/controller/middleware"
	"tiktok/controller/response"
)

func InitRouter(r *gin.Engine, web *controller.Controllers) {
	apiRouter := r.Group("/douyin")
	// 基础接口
	// 视频流
	apiRouter.GET("/feed", web.FeedController.GetFeed)

	// 获取登录用户信息
	apiRouter.GET("/user/", web.UserController.CurrentUser)

	// 用户注册
	apiRouter.POST("/user/register/", web.UserController.Register)

	// 用户登录
	apiRouter.POST("/user/login/", web.UserController.Login)

	// 视频投稿
	apiRouter.POST("/publish/action/", middleware.JWTAuthMiddleware(), web.PublishController.Publish)

	// 发布列表
	apiRouter.GET("/publish/list/", middleware.JWTAuthMiddleware(), web.PublishController.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
		})
	})
	// 打开app请求这个
	apiRouter.GET("/favorite/list/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.VideoListResponse{
			Response: response.Response{
				StatusCode: 0,
			},
			VideoList: []response.VideoVo{},
		})
	})
	apiRouter.POST("/comment/action/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
		})
	})
	apiRouter.GET("/comment/list/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
		})
	})

	//extra apis - II
	apiRouter.POST("/relation/action/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
		})
	})
	apiRouter.GET("/relation/follow/list/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
		})
	})
	apiRouter.GET("/relation/follower/list/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
		})
	})
	apiRouter.GET("/relation/friend/list/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
		})
	})
	apiRouter.GET("/message/chat/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
		})
	})
	apiRouter.POST("/message/action/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
		})
	})

}

var DemoUser = response.UserVo{
	ID:            1,
	UserName:      "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var DemoVideos = []response.VideoVo{
	{
		ID:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}, {
		ID:            2,
		Author:        DemoUser,
		PlayUrl:       "http://tiktok.tudoutiao.pro/video/13cdca069e040698246df2346541e656.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}, {
		ID:            3,
		Author:        DemoUser,
		PlayUrl:       "http://tiktok.tudoutiao.pro/video/20bab060d3a6b70dfbbc4dbab869efe6.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}, {
		ID:            4,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}, {
		ID:            5,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}, {
		ID:            6,
		Author:        DemoUser,
		PlayUrl:       "http://tiktok.tudoutiao.pro/video/24ce64ccfc93b4a44681de669216561d.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}, {
		ID:            7,
		Author:        DemoUser,
		PlayUrl:       "http://tiktok.tudoutiao.pro/video/67d84e3c3984800c0d8f3c075ad97e31.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}, {
		ID:            8,
		Author:        DemoUser,
		PlayUrl:       "http://tiktok.tudoutiao.pro/video/6c3b5c9e0acd8c1c751a1d757249998d.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}, {
		ID:            9,
		Author:        DemoUser,
		PlayUrl:       "http://tiktok.tudoutiao.pro/video/70934a9e7e661341d0237d50aa94d980.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}
