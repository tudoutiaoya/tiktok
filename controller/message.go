package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/errno"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/service"
)

type MessageController struct {
	messageService *service.MessageService
}

func (c *MessageController) MessageAction(context *gin.Context) {
	var MessageParam param.MessageAction
	if err := context.ShouldBind(&MessageParam); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	userID, _ := context.Get("id")
	err := c.messageService.MessageAction(userID.(int64), MessageParam.ToUserID, MessageParam.ActionType, MessageParam.Content)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, response.SuccessResponse)
}

func (c *MessageController) MessageChat(context *gin.Context) {
	var messageChat param.MessageChat
	if err := context.ShouldBind(&messageChat); err != nil {
		response.SendErrResponse(context, errno.ParamIllegal)
		return
	}
	userID, _ := context.Get("id")
	result, err := c.messageService.MessageChat(userID.(int64), messageChat.ToUserID)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	response.SendSuccessResponse(context, result)
}

func NewMessageController(messageService *service.MessageService) *MessageController {
	return &MessageController{messageService: messageService}
}
