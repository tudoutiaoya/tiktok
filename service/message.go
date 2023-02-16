package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"tiktok/controller/response"
	"tiktok/dao"
	"tiktok/domain"
)

const SEND_ACTION = 1

type MessageService struct {
	messageDao *dao.MessageDao
}

func (s *MessageService) MessageAction(userID int64, toUserID int64, actionType int, content string) error {
	if actionType == SEND_ACTION {
		var message = domain.Message{
			Content:  content,
			UserID:   userID,
			ToUserID: toUserID,
		}
		err := s.messageDao.MessageAction(message)
		return err
	}
	return nil
}

func (s *MessageService) MessageChat(userID int64, toUserID int64) (*response.MessageChatResponse, error) {
	messageChatList, err := s.messageDao.MessageChat(userID, toUserID)
	if err != nil {
		return nil, errors.New("查询记录失败")
	}
	var messageChatVoList = ChatListToVo(messageChatList)
	var result = &response.MessageChatResponse{
		Response:    response.SuccessResponse,
		MessageList: messageChatVoList,
	}
	return result, nil
}

func ChatListToVo(chatList []domain.Message) []response.MessageVo {
	var chatVoList []response.MessageVo
	for _, message := range chatList {
		var chatVo response.MessageVo
		copier.Copy(&chatVo, message)
		chatVo.CreateTime = message.CreatedAt.UnixMilli()
		chatVoList = append(chatVoList, chatVo)
	}
	return chatVoList
}

func NewMessageService(messageDao *dao.MessageDao) *MessageService {
	return &MessageService{messageDao: messageDao}
}
