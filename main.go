package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/api"
	"tiktok/config"
	"tiktok/controller"
	"tiktok/dao"
	"tiktok/service"
	"tiktok/util/jwtutil"
)

func main() {
	r := gin.Default()

	initApp(r)

	r.Run(":3333")
}

// 配置文件
var configuration = config.InitConfig()

// 初始化
func initApp(r *gin.Engine) {
	// 初始化其他配置
	initOther()
	// 初始化其他配置
	web := initMVC()
	// 初始化路由
	api.InitRouter(r, web)
}

// 初始化controller service dao
func initMVC() *controller.Controllers {
	databases := dao.InitDao(configuration)
	services := service.InitService(databases)
	controllers := controller.InitController(services)
	return controllers
}

func initOther() {
	// 初始化jwt
	jwtutil.InitJwtSecretKey(configuration)
}
