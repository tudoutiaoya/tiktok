package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tiktok/common/errno"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/service"
	"tiktok/util/jwtutil"
)

type PublishController struct {
	*service.FeedService
}

func NewPublishController(feedService *service.FeedService) *PublishController {
	return &PublishController{FeedService: feedService}
}

// Publish 视频发稿
// 思考？是否文件秒传、断点续传？？？
func (c *PublishController) Publish(context *gin.Context) {
	var publishParam param.PublishParam
	if err := context.ShouldBind(&publishParam); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	file, _ := context.FormFile("data")
	userId, _ := context.Get("id")
	err := c.FeedService.SaveVideo(userId.(int64), publishParam.Title, file)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
	}
	response.SendSuccessResponse(context, response.SuccessResponse)
}

func (c *PublishController) PublishList(context *gin.Context) {
	var currentUser param.CurrentUserParam
	if err := context.ShouldBind(&currentUser); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	// 验证token
	userID, _ := strconv.ParseInt(currentUser.UserID, 10, 64)
	token := currentUser.Token
	parseToken, err := jwtutil.ParseToken(token)
	if err != nil {
		response.SendErrResponse(context, errno.TokenIllegal)
		return
	}
	// 签名不一样
	if parseToken.ID != userID {
		response.SendErrResponse(context, errno.TokenIllegal)
		return
	}
	videoVos, err := c.FeedService.GetUserPublishList(userID)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, response.VideoListResponse{
		Response:  response.SuccessResponse,
		VideoList: videoVos,
	})
}
