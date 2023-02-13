package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/errno"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/service"
)

type CommentController struct {
	commentService *service.CommentService
}

func (c *CommentController) CommentAction(context *gin.Context) {
	var commentAction param.CommentActionParam
	if err := context.ShouldBind(&commentAction); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	userID, _ := context.Get("id")
	result, err := c.commentService.CommentAction(commentAction, userID.(int64))
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, result)
}

func (c *CommentController) CommentList(context *gin.Context) {
	var commentList param.CommentListParam
	if err := context.ShouldBind(&commentList); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	result, err := c.commentService.GetCommentList(commentList.VideoID)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, result)
}

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}
