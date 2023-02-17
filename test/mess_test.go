package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"tiktok/controller/response"
)

func TestMess(t *testing.T) {
	var messageVo = response.MessageVo{
		Content:  "hahhah",
		UserID:   5,
		ToUserID: 2,
	}
	fmt.Println(messageVo)
	// 发送消息到对方的邮箱
	str, _ := json.Marshal(messageVo)
	fmt.Println(string(str))
	//接收到消息
	var mes response.MessageVo
	json.Unmarshal(str, &mes)
	fmt.Println(mes)
}
