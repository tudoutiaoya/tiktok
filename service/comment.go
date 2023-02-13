package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"tiktok/controller/param"
	"tiktok/controller/response"
	"tiktok/dao"
	"tiktok/domain"
	"tiktok/util/timeutil"
)

const ACTION_PUBLISH = 1
const ACTION_DELETE = 2

type CommentService struct {
	commentDao  *dao.CommentDao
	userService *UserService
}

func NewCommentService(commentDao *dao.CommentDao, userService *UserService) *CommentService {
	return &CommentService{commentDao: commentDao, userService: userService}
}

// CommentAction TODO 评论内容是否合法
func (s *CommentService) CommentAction(commentAction param.CommentActionParam, userID int64) (*response.CommentActionResponse, error) {
	comment := &domain.Comment{}
	if commentAction.ActionType == ACTION_PUBLISH {
		commentText := commentAction.CommentText
		if commentText == "" {
			return nil, errors.New("评论内容不能为空")
		}
		videoID := commentAction.VideoID
		comment = &domain.Comment{
			VideoID: videoID,
			UserID:  userID,
			Content: commentText,
		}
		err := s.commentDao.CreatComment(comment)
		if err != nil {
			return nil, errors.New("创建评论失败")
		}
	} else if commentAction.ActionType == ACTION_DELETE {
		commentID := commentAction.CommentID
		if commentID == "" {
			return nil, errors.New("删除的评论id不能为空")
		}
		err := s.commentDao.DeleteComment(commentID)
		if err != nil {
			return nil, errors.New("删除评论失败")
		}
	}
	commentVo := s.copyComment2Vo(*comment)
	commentActionResponse := &response.CommentActionResponse{
		Response:  response.SuccessResponse,
		CommentVo: commentVo,
	}
	return commentActionResponse, nil
}

func (s *CommentService) GetCommentList(videoID int64) (*response.CommentListResponse, error) {
	comments, err := s.commentDao.GetCommentList(videoID)
	if err != nil {
		return nil, errors.New("查询评论列表失败")
	}
	commentVos := s.copyCommentList(comments)
	result := &response.CommentListResponse{
		Response:    response.SuccessResponse,
		CommentList: commentVos,
	}
	return result, nil
}

// comments 转换为 commentVoList
func (s *CommentService) copyCommentList(comments []domain.Comment) []response.CommentVo {
	var commentVoList []response.CommentVo
	for _, comment := range comments {
		commentVo := s.copyComment2Vo(comment)
		commentVoList = append(commentVoList, commentVo)
	}
	return commentVoList
}

func (s *CommentService) copyComment2Vo(comment domain.Comment) response.CommentVo {
	var commentVo = response.CommentVo{}
	copier.Copy(&commentVo, &comment)
	userVo, _ := s.userService.GetCurrentUser(comment.UserID)
	// TODO 有关注操作吗？虽然暂时还没做
	commentVo.UserVo = *userVo
	commentVo.CreatedAt = timeutil.Time2mm_dd(comment.CreatedAt)
	return commentVo
}
