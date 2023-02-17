package mredis

import (
	"fmt"
	"github.com/go-redis/redis"
	"tiktok/config"
)

// 声明一个全局的redisDb变量
var RedisClient *redis.Client

// 根据redis配置初始化一个客户端
func InitClient(config *config.Configuration) (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisSettings.Host, config.RedisSettings.Port), // redis地址
		Password: fmt.Sprintf("%s", config.RedisSettings.PassWord),                           // redis密码，没有则留空
		DB:       config.RedisSettings.DB,                                                    // 默认数据库，默认是0
	})

	//通过 *mredis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
