package service

import (
	"github.com/jinzhu/copier"
	"tiktok/controller/response"
	"tiktok/dao"
	"time"
)

// LIMIT 视频限制条数
const LIMIT = 2

type FeedService struct {
	feedDao     *dao.FeedDao
	userDao     *dao.UserDao
	userService *UserService
}

func NewFeedService(feedDao *dao.FeedDao, userDao *dao.UserDao, userService *UserService) *FeedService {
	return &FeedService{feedDao: feedDao, userDao: userDao, userService: userService}
}

func (s *FeedService) GetFeed(latestTime int64, token string) (*response.FeedResponse, error) {
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
		// 添加视频作者信息
		user, _ := s.userDao.GetUserById(video.AuthorID)
		var userVo = response.UserVo{}
		copier.Copy(&userVo, &user)
		videoVo.Author = userVo
		// 添加是否喜欢
		isLogin, _ := s.userService.IsLogin(token)
		if isLogin {
			// TODO 等点赞接口
			videoVo.IsFavorite = true
		}
		videoVos = append(videoVos, videoVo)
	}

	result := &response.FeedResponse{
		VideoList: videoVos,
		NextTime:  videos[len(videos)-1].UpdatedAt.UnixMilli(),
	}
	return result, nil
}

func (s *FeedService) GetUserVoByVideoId(id int64) (*response.UserVo, error) {
	user, err := s.userDao.GetUserByVideoId(id)
	if err != nil {
		return nil, err
	}
	var userVo = response.UserVo{}
	copier.Copy(&user, &userVo)
	return &userVo, nil
}
