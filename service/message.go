package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"tiktok/controller/response"
	"tiktok/dao"
	"tiktok/domain"
	"tiktok/middleware/mrabbitmq"
	"tiktok/middleware/mredis"
	"tiktok/middleware/msnowflake"
	"time"
)

const SEND_ACTION = 1

type MessageService struct {
	messageDao  *dao.MessageDao
	redisClient *redis.Client
}

func NewMessageService(messageDao *dao.MessageDao, redisClient *redis.Client) *MessageService {
	return &MessageService{messageDao: messageDao, redisClient: redisClient}
}

func (s *MessageService) MessageAction(userID int64, toUserID int64, actionType int, content string) error {
	if actionType == SEND_ACTION {
		var messageVo = response.MessageVo{
			ID:         msnowflake.GenerateID.NextVal(),
			Content:    content,
			UserID:     userID,
			ToUserID:   toUserID,
			CreateTime: time.Now().Unix(),
		}
		// 发送消息到对方的邮箱
		key := mredis.GetToUserChatEmailKey(userID, toUserID)
		mes, _ := json.Marshal(messageVo)
		fmt.Println("发送给对方的消息:", messageVo)
		s.redisClient.RPush(key, mes)
		// todo mq异步保存到数据库
		var message domain.Message
		copier.Copy(&message, &messageVo)
		message.CreatedAt = time.Unix(messageVo.CreateTime, 0)
		theSaveMessage, _ := json.Marshal(message)
		mrabbitmq.MQ.PublishSimple(theSaveMessage)
		//err := s.messageDao.MessageAction(message)
		return nil
	}
	return nil
}

func (s *MessageService) MessageChat(userID int64, toUserID int64) (*response.MessageChatResponse, error) {
	//messageChatList, err := s.messageDao.MessageChat(userID, toUserID)
	//if err != nil {
	//	return nil, errors.New("查询记录失败")
	//}
	// TODO 改进为lua
	// 获取我自己的邮箱
	key := mredis.GetMyChatEmailKey(userID, toUserID)
	list, err := s.redisClient.LRange(key, 0, -1).Result()
	var messages []response.MessageVo
	for _, str := range list {
		var mes response.MessageVo
		json.Unmarshal([]byte(str), &mes)
		messages = append(messages, mes)
	}
	if err != nil {
		return nil, errors.New("获取消息列表错误")
	}
	// 清空所有
	s.redisClient.LTrim(key, 1, 0)
	fmt.Println("接收到的消息", messages)
	//redisClient.Eval()
	//var messageChatVoList = ChatListToVo(messageChatList)
	var result = &response.MessageChatResponse{
		Response:    response.SuccessResponse,
		MessageList: messages,
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
