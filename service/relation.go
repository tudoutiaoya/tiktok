package service

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"strconv"
	"tiktok/controller/response"
	"tiktok/dao"
	"tiktok/domain"
	"tiktok/middleware/mredis"
)

const FOLLOW_ACTION = 1
const UNFOLLOW_ACTION = 2

type RelationService struct {
	relationDao *dao.RelationDao
	userService *UserService
	messageDao  *dao.MessageDao
	redisClient *redis.Client
	userDao     *dao.UserDao
}

func NewRelationService(relationDao *dao.RelationDao, userService *UserService, messageDao *dao.MessageDao, redisClient *redis.Client, userDao *dao.UserDao) *RelationService {
	return &RelationService{relationDao: relationDao, userService: userService, messageDao: messageDao, redisClient: redisClient, userDao: userDao}
}

func (s *RelationService) RelationAction(userID int64, toUserID int64, actionType int64) error {
	//if toUserID == userID {
	//	return errors.New("不允许关注自己")
	//}
	//toUser, _ := s.userService.GetCurrentUser(toUserID)
	//if toUser == nil {
	//	return errors.New("关注用户不存在")
	//}
	//if actionType == FOLLOW_ACTION {
	//	isHas, _ := s.relationDao.IsHas(userID, toUserID)
	//	if isHas {
	//		return errors.New("已经关注过了")
	//	}
	//	err := s.relationDao.Follow(userID, toUserID)
	//	if err != nil {
	//		return errors.New("关注失败")
	//	}
	//} else if actionType == UNFOLLOW_ACTION {
	//	err := s.relationDao.UnFollow(userID, toUserID)
	//	if err != nil {
	//		return errors.New("取消关注失败")
	//	}
	//}

	// 改造为Redis  不设置过期时间
	if toUserID == userID {
		return errors.New("不允许关注自己")
	}
	followKey := mredis.FOLLOW_LIST + strconv.FormatInt(userID, 10)
	// 判断关注用户存在不存在
	toUser, _ := s.userService.GetCurrentUser(toUserID)
	if toUser == nil {
		return errors.New("关注用户不存在")
	}
	// 判断是否关注
	isMember := s.redisClient.SIsMember(followKey, toUserID)
	if isMember.Val() {
		return errors.New("已经关注过了")
	}
	// 关注
	// 添加自己的关注列表
	s.redisClient.SAdd(followKey, toUserID)
	// 对方粉丝列表添加我
	followerKey := mredis.FOLLOWER_LIST + strconv.FormatInt(toUserID, 10)
	s.redisClient.SAdd(followerKey, userID)
	return nil
}

// FollowList 关注列表
func (s *RelationService) FollowList(userID int64) (*response.UserListResponse, error) {
	follows, _ := s.relationDao.FollowList(userID)
	// todo 改造为redis
	//followKey := mredis.FOLLOW_LIST + strconv.FormatInt(userID, 10)
	//fanIDs := s.redisClient.SMembers(followKey).Val()
	//follows, _ := s.userDao.GetUserByIds(fanIDs)
	userVos := s.UserListToVoWithFollow(userID, follows)

	var result = &response.UserListResponse{
		Response: response.SuccessResponse,
		UserList: userVos,
	}
	return result, nil
}

// FollowerList 粉丝列表
func (s *RelationService) FollowerList(userID int64) (*response.UserListResponse, error) {
	fans, _ := s.relationDao.FansList(userID)

	// todo 改造为redis
	//followKey := mredis.FOLLOWER_LIST + strconv.FormatInt(userID, 10)
	//fanIDs := s.redisClient.SMembers(followKey).Val()
	//fans, _ := s.userDao.GetUserByIds(fanIDs)
	userVos := s.UserListToVoWithFollow(userID, fans)
	var result = &response.UserListResponse{
		Response: response.SuccessResponse,
		UserList: userVos,
	}
	return result, nil
}

// FriendList 朋友列表
func (s *RelationService) FriendList(userID int64) (*response.FriendListResponse, error) {
	fans, _ := s.relationDao.FriendList(userID)
	// todo 改造为redis, 朋友列表取交集

	//s.redisClient.SInter()
	userVos := s.UserListToVoWithFollow(userID, fans)

	var friendList []response.FriendUser

	for _, userVo := range userVos {
		msg := s.messageDao.GetNewestMsg(userVo.ID)
		var friend = response.FriendUser{
			UserVo:  userVo,
			Message: msg.Content,
			// 根据当前用户的id来判断消息
			MsgType: getMsgType(userID, msg),
		}
		friendList = append(friendList, friend)
	}

	var result = &response.FriendListResponse{
		Response:   response.SuccessResponse,
		FriendList: friendList,
	}
	return result, nil
}

// 获取消息的类型
func getMsgType(userID int64, msg domain.Message) int64 {
	if userID == msg.UserID {
		// 当前请求用户发送的消息
		return 1
	}
	// 当前请求用户接收的消息
	return 0
}

// UserListToVoWithFollow 用户列表转换为vo
func (s *RelationService) UserListToVoWithFollow(userID int64, followsOrFans []domain.User) []response.UserVo {
	var userVoList []response.UserVo
	for _, followOrfan := range followsOrFans {
		var userVo response.UserVo
		copier.Copy(&userVo, &followOrfan)
		isHas, _ := s.relationDao.IsHas(userID, int64(followOrfan.ID))
		userVo.IsFollow = isHas
		userVoList = append(userVoList, userVo)
	}
	return userVoList
}
