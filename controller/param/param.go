package param

// FeedParam 获取视频流参数
type FeedParam struct {
	LatestTime string `form:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `form:"token"`       // 用户登录状态下设置
}

// UserParam 用户登录注册参数
type UserParam struct {
	Username string `form:"username" binding:"required"` // 注册用户名，最长32个字符
	Password string `form:"password" binding:"required"` // 密码，最长32个字符
}

// CurrentUserParam 获取当前用户信息参数
type CurrentUserParam struct {
	UserID string `form:"user_id" binding:"required"` // 用户id
	Token  string `form:"token" binding:"required"`   // 用户鉴权token
}
