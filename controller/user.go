package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tiktok/common/errno"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/service"
	"tiktok/util/jwtutil"
	"tiktok/util/strutil"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// Register 用户注册
func (c *UserController) Register(context *gin.Context) {
	var userParam param.UserParam
	if err := context.ShouldBind(&userParam); err != nil {
		response.SendErrResponse(context, errno.LoginParamIllegal)
		return
	}
	// 思考？用户名和密码中间有空格怎么处理？？？
	// 参数处理
	username := strutil.StringStrip(userParam.Username)
	password := strutil.StringStrip(userParam.Password)

	// 注册用户
	user, err := c.userService.Register(username, password)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	// 成功则返回token
	token, _ := jwtutil.GenToken(int64(user.ID))
	response.SendSuccessResponse(context, response.UserLoginResponse{
		Response: response.SuccessResponse,
		UserId:   int64(user.ID),
		Token:    token,
	})
}

// Login 用户登录
func (c *UserController) Login(context *gin.Context) {
	var userParam param.UserParam
	if err := context.ShouldBind(&userParam); err != nil {
		response.SendErrResponse(context, errno.LoginParamIllegal)
		return
	}
	userVo, err := c.userService.Login(userParam.Username, userParam.Password)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
		return
	}
	// 获取token并返回
	token, _ := jwtutil.GenToken(userVo.ID)
	response.SendSuccessResponse(context, response.UserLoginResponse{
		Response: response.SuccessResponse,
		UserId:   userVo.ID,
		Token:    token,
	})
}

func (c *UserController) CurrentUser(context *gin.Context) {
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
	userVo, err := c.userService.GetCurrentUser(userID)
	if err != nil {
		response.SendErrResponse(context, errno.HandleServiceErrRes(err))
	}
	response.SendSuccessResponse(context, response.UserResponse{
		Response: response.SuccessResponse,
		User:     *userVo,
	})
}
