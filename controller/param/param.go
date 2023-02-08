package param

type FeedParam struct {
	LatestTime int64  // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string // 用户登录状态下设置
}
