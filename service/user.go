package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"tiktok/controller/response"
	"tiktok/dao"
	"tiktok/domain"
	"tiktok/util/encryptutil"
	"tiktok/util/jwtutil"
)

const SALT = "wokanguoliuxingzhuiruodimian"

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{userDao: userDao}
}

func (s *UserService) Register(username string, password string) (*domain.User, error) {
	// 验证参数
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("用户名或密码不能为空")
	}
	if len(username) > 32 {
		return nil, errors.New("用户名太长")
	}
	if len(password) > 32 {
		return nil, errors.New("密码太长")
	}
	// 判断是否存在
	count := s.userDao.SelectCount(username)
	if count > 0 {
		return nil, errors.New("用户已经存在")
	}
	// 加密保存到数据库
	hashPassword, _ := encryptutil.HashPassword(password)
	user := &domain.User{
		UserName: username,
		PassWord: hashPassword,
	}
	err := s.userDao.CreatUse(user)
	if err != nil {
		return nil, err
	}
	// 返回用户信息
	return user, nil
}

func (s *UserService) Login(username string, password string) (*response.UserVo, error) {
	// 验证参数
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("用户名或密码不能为空")
	}
	user, err := s.userDao.GetUserByUserName(username)
	if err != nil {
		return nil, err
	}
	// 判断用户是否存在
	if user == (domain.User{}) {
		return nil, errors.New("用户不存在")
	}
	if !encryptutil.CheckPasswordHash(password, user.PassWord) {
		return nil, errors.New("密码不正确，请重新输入")
	}
	var userVo response.UserVo
	copier.Copy(&userVo, &user)
	return &userVo, nil
}

func (s *UserService) GetCurrentUser(id int64) (*response.UserVo, error) {
	user, err := s.userDao.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if user == (domain.User{}) {
		return nil, errors.New("用户不存在")
	}
	var userVo response.UserVo
	copier.Copy(&userVo, &user)
	return &userVo, nil
}

func (s *UserService) IsLogin(token string) (bool, *jwtutil.CustomClaims) {
	parseToken, err := jwtutil.ParseToken(token)
	if err != nil {
		return false, nil
	}
	return true, parseToken
}
