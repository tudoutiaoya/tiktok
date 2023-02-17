package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/api"
	"tiktok/config"
	"tiktok/controller"
	"tiktok/dao"
	"tiktok/middleware/mrabbitmq"
	"tiktok/middleware/mredis"
	"tiktok/service"
	"tiktok/util/jwtutil"
	"tiktok/util/qiniuutil"
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
	// 初始化rabbitmq
	mrabbitmq.NewRabbitMQSimple("saveMessage", databases.MessageDao)
	// 开启两个监听 可以根据数量添加，但是单体项目中，可能会造成性能损耗
	go mrabbitmq.MQ.ConsumeSimple()
	go mrabbitmq.MQ.ConsumeSimple()
	return controllers
}

func initOther() {
	// 初始化jwt
	jwtutil.InitJwtSecretKey(configuration)
	// 初始化七牛云
	qiniuutil.InitQiniu(configuration)
	// 初始化Redis
	err := mredis.InitClient(configuration)
	if err != nil {
		//redis连接错误
		panic(err)
	}
}
