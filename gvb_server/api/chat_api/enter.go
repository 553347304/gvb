package chat_api

import "github.com/gorilla/websocket"

type ChatApi struct{}

const (
	outChatRoom   = -1 // 离开聊天室
	inChatRoom    = 0  // 进入聊天室
	textMessage   = 1  // 文本消息
	imageMessage  = 2  // 图片消息
	voiceMessage  = 3  // 语音消息
	videoMessage  = 4  // 视频消息
	systemMessage = 5  // 系统消息
)

var ConnGroupMap = map[string]User{}

type User struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

type Group struct {
	NickName  string `json:"nick_name"`  // 前端自己生成
	Avatar    string `json:"avatar"`     // 头像
	Message   string `json:"message"`    // 聊天的内容
	Type      int    `json:"type"`       // 聊天类型
	Online    int    `json:"online"`     // 在线人数
	CreatedAt string `json:"created_at"` // 消息的时间
}
