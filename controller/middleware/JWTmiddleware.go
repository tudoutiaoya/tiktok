package middleware

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/errno"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/util/jwtutil"
)

// JWTAuthMiddleware 基于JWT的认证中间件 用于验证token是否合法
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里是请求体中
		var tokenParam param.TokenParam
		err := c.ShouldBind(&tokenParam)
		if err != nil {
			response.SendErrResponse(c, errno.NoLoginErr)
			return
		}
		parseToken, err := jwtutil.ParseToken(tokenParam.Token)
		if err != nil {
			response.SendErrResponse(c, errno.NoLoginErr)
			return
		}
		id := parseToken.ID
		// 将当前请求的id信息保存到请求的上下文c上
		c.Set("id", id)
		c.Next() // 后续的处理函数可以用过c.Get("id")来获取当前请求的用户信息
	}
}
