package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

type UserResponse struct {
	Response
	Content string `json:"content"`
}

func (c *UserController) Hello(g *gin.Context) {
	g.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: http.StatusOK},
		Content:  "hello gin",
	})
}

func (c *UserController) Create(context *gin.Context) {
	err := c.userService.Create()
	if err != nil {

	}
	context.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: http.StatusOK},
		Content:  "创建成功",
	})
}
