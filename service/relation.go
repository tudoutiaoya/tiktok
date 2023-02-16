package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"tiktok/controller/response"
	"tiktok/dao"
	"tiktok/domain"
)

const FOLLOW_ACTION = 1
const UNFOLLOW_ACTION = 2

type RelationService struct {
	relationDao *dao.RelationDao
	userService *UserService
}

func NewRelationService(relationDao *dao.RelationDao, userService *UserService) *RelationService {
	return &RelationService{relationDao: relationDao, userService: userService}
}

func (s *RelationService) RelationAction(userID int64, toUserID int64, actionType int64) error {
	if toUserID == userID {
		return errors.New("不允许关注自己")
	}
	toUser, _ := s.userService.GetCurrentUser(toUserID)
	if toUser == nil {
		return errors.New("关注用户不存在")
	}
	if actionType == FOLLOW_ACTION {
		isHas, _ := s.relationDao.IsHas(userID, toUserID)
		if isHas {
			return errors.New("已经关注过了")
		}
		err := s.relationDao.Follow(userID, toUserID)
		if err != nil {
			return errors.New("关注失败")
		}
	} else if actionType == UNFOLLOW_ACTION {
		err := s.relationDao.UnFollow(userID, toUserID)
		if err != nil {
			return errors.New("取消关注失败")
		}
	}
	return nil
}

// FollowList 关注列表
func (s *RelationService) FollowList(userID int64) (*response.UserListResponse, error) {
	follows, _ := s.relationDao.FollowList(userID)

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

	userVos := s.UserListToVoWithFollow(userID, fans)

	var result = &response.UserListResponse{
		Response: response.SuccessResponse,
		UserList: userVos,
	}
	return result, nil
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
