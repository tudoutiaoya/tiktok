package test

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"tiktok/middleware/mredis"
)

func TestLua(t *testing.T) {
	var client = mredis.RedisClient
	var luaScript = redis.NewScript(`
		local key = KEYS[1];
		local list = redis.call('lrange',key,'0','-1');
		
		if(#list > 0) then
			-- 清空list
			redis.call('ltrim',key,'1','0');
			return list;
		end
		return;
    `)
	result, _ := luaScript.Run(client, []string{"list"}, 1).Result()
	fmt.Println("结果为:", result)
}

func TestLuaDemo(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "43.143.224.79:6379", // redis地址
		Password: "123456",             // redis密码，没有则留空
		DB:       0,                    // 默认数据库，默认是0
	})
	result, _ := client.LRange("list", 0, -1).Result()
	fmt.Println(result[0])
}
