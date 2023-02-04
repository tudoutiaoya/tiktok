package api

import (
	"github.com/gin-gonic/gin"
	"tiktok/config"
	"tiktok/controller"
	"tiktok/dao"
	"tiktok/service"
)

func InitRouter(r *gin.Engine) {
	// 初始化其他配置
	web := initOther()
	apiRouter := r.Group("/douyin")
	{
		apiRouter.GET("/create", web.UserController.Create)
	}
}

func initOther() *controller.Controllers {
	configuration := config.InitConfig()
	databases := dao.InitDao(configuration)
	services := service.InitService(databases)
	controllers := controller.InitController(services)
	return controllers
}
