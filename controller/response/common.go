package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/common/errno"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

var SuccessResponse = Response{
	StatusCode: 0,
	StatusMsg:  "success",
}

// 返回前端错误信息
func SendErrResponse(context *gin.Context, errResponse *errno.ErrResponse) {
	context.JSON(http.StatusOK, errResponse)
	context.Abort()
}

// 成功返回消息
func SendSuccessResponse(context *gin.Context, response interface{}) {
	context.JSON(http.StatusOK, response)
}

type VideoVo struct {
	ID            int64  `json:"id,omitempty"`
	Author        UserVo `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type UserVo struct {
	ID            int64  `json:"id,omitempty"`
	UserName      string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
}

type CommentVo struct {
	ID        int64  `json:"id,omitempty"`
	UserVo    UserVo `json:"user"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"create_date,omitempty"`
}

type MessageVo struct {
	ID         int64  `json:"id,omitempty"`
	UserID     int64  `json:"from_user_id"`
	ToUserID   int64  `json:"to_user_id"`
	Content    string `json:"content,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}
