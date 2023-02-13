package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/errno"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/service"
)

type PublishController struct {
	*service.VideoService
}

func NewPublishController(feedService *service.VideoService) *PublishController {
	return &PublishController{VideoService: feedService}
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
	err := c.VideoService.SaveVideo(userId.(int64), publishParam.Title, file)
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
	// 验证token中的id和用户的id是否相等
	tokenUserId, _ := context.Get("id")
	userID := currentUser.UserID
	// 签名不一样
	if tokenUserId != userID {
		response.SendErrResponse(context, errno.TokenIllegal)
		return
	}

	videoVos, err := c.VideoService.GetUserPublishList(userID)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, response.VideoListResponse{
		Response:  response.SuccessResponse,
		VideoList: videoVos,
	})
}
