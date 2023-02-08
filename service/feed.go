package service

import (
	"github.com/jinzhu/copier"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/dao"
	"time"
)

var LIMIT = 2

type FeedService struct {
	feedDao *dao.FeedDao
}

func NewFeedService(feedDao *dao.FeedDao) *FeedService {
	return &FeedService{feedDao: feedDao}
}

func (s *FeedService) GetFeed(feedParam param.FeedParam) (result *response.FeedResponse, err error) {

	latestTime := feedParam.LatestTime
	videos, err := s.feedDao.GetFeed(LIMIT, latestTime)
	if err != nil {
		return nil, err
	}
	// 为空，返回当前时间
	if len(videos) == 0 {
		return &response.FeedResponse{
			VideoList: nil,
			NextTime:  time.Now().UnixMilli(),
		}, nil
	}

	var videoVos = make([]response.VideoVo, 0)
	for _, video := range videos {
		videoVo := response.VideoVo{}
		copier.Copy(&videoVo, &video)
		// 查询作者 之后添加
		// 是否喜欢 之后添加
		videoVo.Author = response.UserVo{
			ID:            1,
			Name:          "TestUser",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}

		videoVos = append(videoVos, videoVo)
	}

	res := &response.FeedResponse{
		VideoList: videoVos,
		NextTime:  videos[len(videos)-1].UpdatedAt.UnixMilli(),
	}
	return res, nil
}
