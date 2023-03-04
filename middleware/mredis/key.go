package mredis

import "fmt"

var CHAT_ROOM_MESSAGE = "douyin:chatroom:"

// GetToUserChatEmailKey 给对方发消息，发送消息到对方的邮箱
// a->b b:a userID=a toUserID=b
func GetToUserChatEmailKey(userID int64, toUserID int64) string {
	return fmt.Sprintf("%s%d:%d", CHAT_ROOM_MESSAGE, toUserID, userID)
}

// GetMyChatEmailKey 我拉取对方给我发送的消息
// b获取 b:a userID=b toUserID=a
func GetMyChatEmailKey(userID int64, toUserID int64) string {
	return fmt.Sprintf("%s%d:%d", CHAT_ROOM_MESSAGE, userID, toUserID)
}

// FOLLOW_LIST 用户关注列表
var FOLLOW_LIST = "douyin:follow:"

// FOLLOWER_LIST 用户粉丝列表
var FOLLOWER_LIST = "douyin:follower:"
