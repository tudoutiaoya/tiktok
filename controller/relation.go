package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/errno"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/service"
)

type RelationController struct {
	relationService *service.RelationService
}

func (c *RelationController) RelationAction(context *gin.Context) {
	var relationAction param.RelationActionParam
	if err := context.ShouldBind(&relationAction); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	userID, _ := context.Get("id")
	err := c.relationService.RelationAction(userID.(int64), relationAction.ToUserID, relationAction.ActionType)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, response.SuccessResponse)
}

func (c *RelationController) FollowList(context *gin.Context) {
	var currentUserParam param.CurrentUserParam
	if err := context.ShouldBind(&currentUserParam); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	result, err := c.relationService.FollowList(currentUserParam.UserID)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, result)
}

func (c *RelationController) FollowerList(context *gin.Context) {
	var currentUserParam param.CurrentUserParam
	if err := context.ShouldBind(&currentUserParam); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	result, err := c.relationService.FollowerList(currentUserParam.UserID)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, result)
}

func (c *RelationController) FriendList(context *gin.Context) {
	var currentUserParam param.CurrentUserParam
	if err := context.ShouldBind(&currentUserParam); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	result, err := c.relationService.FriendList(currentUserParam.UserID)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, result)
}

func NewRelationController(relationService *service.RelationService) *RelationController {
	return &RelationController{relationService: relationService}
}
