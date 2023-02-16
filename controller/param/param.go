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

// CurrentUserParam 获取当前用户信息参数 视频发布列表参数
type CurrentUserParam struct {
	UserID int64  `form:"user_id" binding:"required"` // 用户id
	Token  string `form:"token" binding:"required"`   // 用户鉴权token
}

// PublishParam 视频投稿
type PublishParam struct {
	Token string `form:"token" binding:"required"` // 用户鉴权token
	Title string `form:"title" binding:"required"` // 视频标题
}

// FavoriteActionParam 点赞/取消
type FavoriteActionParam struct {
	Token      string `form:"token" binding:"required"`
	VideoID    int64  `form:"video_id" binding:"required"`
	ActionType int64  `form:"action_type" binding:"required,oneof=1 2"`
}

// CommentActionParam 评论操作
type CommentActionParam struct {
	Token       string `form:"token" binding:"required"`                 // 用户鉴权token
	VideoID     int64  `form:"video_id" binding:"required"`              // 视频id
	ActionType  int    `form:"action_type" binding:"required,oneof=1 2"` // 1-发表评论 2-删除评论
	CommentText string `form:"comment_text"`                             // 评论内容 action_type=1时候用
	CommentID   string `form:"comment_id"`                               // 要删除的评论id,在action_type=2时使用
}

// 评论列表
type CommentListParam struct {
	Token   string `form:"token"`                       // 用户鉴权token
	VideoID int64  `form:"video_id" binding:"required"` // 视频id
}

// RelationActionParam 关注操作
type RelationActionParam struct {
	Token      string `form:"token"`                                    // 用户鉴权token
	ToUserID   int64  `form:"to_user_id" binding:"required"`            // 视频id
	ActionType int64  `form:"action_type" binding:"required,oneof=1 2"` // 1-关注 2-取消关注
}

// MessageAction 发送消息
type MessageAction struct {
	Token      string `form:"token"`       // 用户鉴权token
	ToUserID   int64  `form:"to_user_id"`  // 对方用户id
	ActionType int    `form:"action_type"` // 1-发送消息
	Content    string `form:"content"`     // 消息内容
}

// MessageChat 聊天记录
type MessageChat struct {
	Token    string `form:"token"`      // 用户鉴权token
	ToUserID int64  `form:"to_user_id"` // 对方用户id
}

// TokenParam JWT中间件判断token是否合法
type TokenParam struct {
	UserID string `form:"user_id"` // 用户id
	Token  string `form:"token" binding:"required"`
}
