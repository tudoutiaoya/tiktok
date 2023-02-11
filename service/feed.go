package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"mime/multipart"
	"tiktok/controller/response"
	"tiktok/dao"
	"tiktok/domain"
	"tiktok/util/qiniuutil"
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

func (s *FeedService) SaveVideo(id int64, title string, file *multipart.FileHeader) error {
	playUrl, err := qiniuutil.Upload(file)
	if err != nil {
		return errors.New("文件上传失败")
	}

	coverUrl := "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
	var video = domain.Video{
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
		AuthorID: id,
	}
	return s.feedDao.SaveVideo(video)
}

func (s *FeedService) GetUserPublishList(id int64) ([]response.VideoVo, error) {
	user, _ := s.userDao.GetUserById(id)
	var useVo response.UserVo
	if user != (domain.User{}) {
		copier.Copy(&useVo, &user)
	}
	videos, err := s.feedDao.GetUserPublishList(id)
	if err != nil {
		return nil, err
	}
	var videoVos []response.VideoVo
	for _, video := range videos {
		var videoVo response.VideoVo
		copier.Copy(&videoVo, &video)
		videoVo.Author = useVo
		// TODO 用户是否点赞
		videoVos = append(videoVos, videoVo)
	}
	return videoVos, nil
}
