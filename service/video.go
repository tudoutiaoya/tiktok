package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"log"
	"mime/multipart"
	"tiktok/controller/response"
	"tiktok/dao"
	"tiktok/domain"
	"tiktok/util/jwtutil"
	"tiktok/util/qiniuutil"
	"time"
)

// LIMIT 视频限制条数
const LIMIT = 2

// LIKE_ACTION 点赞
const LIKE_ACTION = 1

// DISLIKE_ACTION 取消点赞
const DISLIKE_ACTION = 2

type VideoService struct {
	videoDao    *dao.VideoDao
	userDao     *dao.UserDao
	userService *UserService
	relationDao *dao.RelationDao
}

func NewVideoService(videoDao *dao.VideoDao, userDao *dao.UserDao, userService *UserService, relationDao *dao.RelationDao) *VideoService {
	return &VideoService{videoDao: videoDao, userDao: userDao, userService: userService, relationDao: relationDao}
}

// GetFeed 获取视频流
func (s *VideoService) GetFeed(latestTime int64, token string) (*response.FeedResponse, error) {
	videos, err := s.videoDao.GetFeed(LIMIT, latestTime)
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

	videoVos := s.videosToVideoVosNoKnownLogin(videos, token)

	result := &response.FeedResponse{
		VideoList: videoVos,
		NextTime:  videos[len(videos)-1].UpdatedAt.UnixMilli(),
	}
	return result, nil
}

// 视频列表封装为响应对象，包含作者信息
func (s *VideoService) videosToVideoVos(videos []domain.Video, knownLogin bool, token string) []response.VideoVo {
	var videoVos = make([]response.VideoVo, 0)
	for _, video := range videos {
		videoVo := response.VideoVo{}
		copier.Copy(&videoVo, &video)
		// 添加视频作者信息
		author, _ := s.userDao.GetUserById(video.AuthorID)
		var userVo = response.UserVo{}
		copier.Copy(&userVo, &author)
		videoVo.Author = userVo
		// 添加是否关注
		// 添加是否喜欢
		isLogin, _ := s.userService.IsLogin(token)
		if isLogin || knownLogin {
			// TODO 等点赞接口
			parseToken, _ := jwtutil.ParseToken(token)
			userID := parseToken.ID
			isLike, _ := s.videoDao.IsLike(int64(author.ID), int64(video.ID))
			videoVo.IsFavorite = isLike
			isHas, _ := s.relationDao.IsHas(userID, int64(author.ID))
			videoVo.Author.IsFollow = isHas
		}
		videoVos = append(videoVos, videoVo)
	}
	return videoVos
}

// 已经登录--视频列表封装为响应对象，包含作者信息
func (s *VideoService) videosToVideoVosHasLogin(videos []domain.Video, token string) []response.VideoVo {
	return s.videosToVideoVos(videos, true, token)
}

// 不知道是否登录--视频列表封装为响应对象，包含作者信息
func (s *VideoService) videosToVideoVosNoKnownLogin(videos []domain.Video, token string) []response.VideoVo {
	return s.videosToVideoVos(videos, false, token)
}

// GetUserVoByVideoId 获取用户通过视频id
func (s *VideoService) GetUserVoByVideoId(id int64) (*response.UserVo, error) {
	user, err := s.userDao.GetUserByVideoId(id)
	if err != nil {
		return nil, err
	}
	var userVo = response.UserVo{}
	copier.Copy(&user, &userVo)
	return &userVo, nil
}

// SaveVideo 保存视频
func (s *VideoService) SaveVideo(id int64, title string, file *multipart.FileHeader) error {
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
	return s.videoDao.SaveVideo(video)
}

// GetUserPublishList 获取用户发布视频列表
func (s *VideoService) GetUserPublishList(id int64) ([]response.VideoVo, error) {
	user, _ := s.userDao.GetUserById(id)
	var useVo response.UserVo
	if user != (domain.User{}) {
		copier.Copy(&useVo, &user)
	}
	videos, err := s.videoDao.GetUserPublishList(id)
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

// Action 视频点赞
func (s *VideoService) Action(userID int64, videoID int64, actionType int64) error {
	var uLikeV domain.ULikeV
	uLikeV.UserID = userID
	uLikeV.VideoID = videoID
	video, _ := s.videoDao.GetVideoById(videoID)
	if video == (domain.Video{}) {
		return errors.New("视频不存在")
	}
	var err error
	if actionType == LIKE_ACTION {
		isLike, _ := s.videoDao.IsLike(userID, videoID)
		if isLike {
			return errors.New("已经点过赞了")
		}
		err = s.videoDao.Like(uLikeV)
	} else if actionType == DISLIKE_ACTION {
		err = s.videoDao.DisLike(uLikeV)
	}
	if err != nil {
		log.Println(err)
		return errors.New("点赞失败")
	}
	return nil
}

// LikeList 喜欢视频列表
func (s *VideoService) LikeList(userID int64, token string) (*response.VideoListResponse, error) {
	videoIDs, err := s.videoDao.LikeListVideoIDs(userID)
	if err != nil {
		log.Println(err)
		return nil, errors.New("查询失败")
	}
	videos, err := s.videoDao.GetVideoListByIDs(videoIDs)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	videoVos := s.videosToVideoVosHasLogin(videos, token)
	result := &response.VideoListResponse{
		Response:  response.SuccessResponse,
		VideoList: videoVos,
	}
	return result, nil
}
