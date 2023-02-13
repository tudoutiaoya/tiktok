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
	feedService *service.VideoService
}

func NewFeedController(feedService *service.VideoService) *FeedController {
	return &FeedController{feedService: feedService}
}

func (c *FeedController) GetFeed(context *gin.Context) {
	var feedParam param.FeedParam
	var latestTime int64
	var token string
	if err := context.ShouldBind(&feedParam); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}

	// 校验参数合法性 为 "" 转化出错
	// 参数不合法 直接返回， 还存在一种为 0
	if len(feedParam.LatestTime) != 0 {
		latestTimeInt64, err := strconv.ParseInt(feedParam.LatestTime, 10, 64)
		if err != nil {
			response.SendErrResponse(context, errno.ParamIllegal)
			return
		}
		latestTime = latestTimeInt64
	}

	// 不传或者为0，为当前时间
	if feedParam.LatestTime == "" || latestTime == 0 {
		latestTime = time.Now().UnixMilli()
	}

	token = feedParam.Token

	// 查询Feed流
	result, err := c.feedService.GetFeed(latestTime, token)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	// 成功响应
	response.SendSuccessResponse(context, result)
}
