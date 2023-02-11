package response

// FeedResponse 视频Feed流
type FeedResponse struct {
	Response
	VideoList []VideoVo `json:"video_list,omitempty"`
	NextTime  int64     `json:"next_time"`
	//NextTime  int64     `json:"next_time,omitempty"`
}

// VideoListResponse 发表-喜欢视频列表
type VideoListResponse struct {
	Response
	VideoList []VideoVo `json:"video_list"`
}

// UserLoginResponse 用户登录注册
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// UserResponse 当前用户信息
type UserResponse struct {
	Response
	User UserVo `json:"user"`
}
