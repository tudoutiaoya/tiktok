package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/errno"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/service"
)

type FavoriteController struct {
	videoService *service.VideoService
}

// TODO 只点赞一次没做 查看是否已经点赞也有bug
func (c FavoriteController) Action(context *gin.Context) {
	var favoriteAction param.FavoriteActionParam
	if err := context.ShouldBind(&favoriteAction); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	userID, _ := context.Get("id")
	err := c.videoService.Action(userID.(int64), favoriteAction.VideoID, favoriteAction.ActionType)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, response.SuccessResponse)
}

func (c FavoriteController) LikeList(context *gin.Context) {
	var currentUserParam param.CurrentUserParam
	if err := context.ShouldBind(&currentUserParam); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	result, err := c.videoService.LikeList(currentUserParam.UserID, currentUserParam.Token)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, result)
}

func NewFavoriteController(videoService *service.VideoService) *FavoriteController {
	return &FavoriteController{videoService: videoService}
}
