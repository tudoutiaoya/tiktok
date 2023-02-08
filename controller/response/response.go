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
