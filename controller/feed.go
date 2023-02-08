package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tiktok/common/errno"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/service"
	"time"
)

type FeedController struct {
	feedService *service.FeedService
}

func NewFeedController(feedService *service.FeedService) *FeedController {
	return &FeedController{feedService: feedService}
}

func (c *FeedController) GetFeed(g *gin.Context) {
	var feedParam param.FeedParam
	var latestTime int64
	var token string
	latest_time := g.Query("latest_time")
	token = g.Query("token")

	// 校验参数合法性 为 "" 转化出错
	// 参数不合法 直接返回， 还存在一种为 0
	if len(latest_time) != 0 {
		latestTimeInt64, err := strconv.ParseInt(latest_time, 10, 64)
		if err != nil {
			response.SendResponse(g, errno.ParamIllegalErr)
			return
		}
		latestTime = latestTimeInt64
	}

	// 不传或者为0，为当前时间
	if latest_time == "" || latestTime == 0 {
		latestTime = time.Now().UnixMilli()
	}

	feedParam.LatestTime = latestTime
	feedParam.Token = token

	// 查询Feed流
	result, err := c.feedService.GetFeed(feedParam)
	if err != nil {
		response.SendResponse(g, errno.HandleServiceErrRes(err))
	}

	response.SendResponse(g, result)
}
