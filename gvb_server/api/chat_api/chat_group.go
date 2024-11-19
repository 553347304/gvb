package chat_api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/System/Time"
	"gvb_server/INIT/global"
	"gvb_server/models"
	"net/http"
)

func (ChatApi) GroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 鉴权 true表示放行，false表示拦截
		},
	}
	// 将http升级至websocket
	conn, _ := upGrader.Upgrade(c.Writer, c.Request, nil)
	addr := conn.RemoteAddr().String() // 连接对象

	user := User{
		Conn:     conn,
		NickName: BY.Random.AnimeName(), // 随机名称
		Avatar:   "/file/head/init.png", // 随机头像
	}
	ConnGroupMap[addr] = user

	chatRoom(inChatRoom, user)

	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		var request Group
		BY.Json.Unmarshal(p, &request)
		if err != nil {
			chatRoom(outChatRoom, user)
			break
		}
		chatRoom(request.Type, user, request.Message)
	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

func chatRoom(Type int, user User, _Message ...string) {
	addr := ""
	Addr := user.Conn.RemoteAddr().String()
	Message := user.NickName
	switch Type {
	case outChatRoom:
		Message = Message + "离开聊天室"
	case inChatRoom:
		Message = Message + "进入聊天室"
	case textMessage:
		Message = _Message[0]
		if Message == "" {
			addr = Addr
			Message = "消息不能为空"
			break
		}
	case imageMessage:
	case voiceMessage:
	case videoMessage:
	case systemMessage:
	default:
		addr = Addr
		Message = "消息类型错误"
	}
	// 发送消息
	SendMessage(addr, BY.Json.Marshal(Group{
		NickName:  user.NickName,
		Avatar:    user.Avatar,
		Message:   Message,
		Type:      Type,
		Online:    len(ConnGroupMap),
		CreatedAt: Time.Now(),
	}))
	global.DB.Create(&models.ChatModel{
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Message:  Message,
		Addr:     Addr,
		IsGroup:  addr == "",
		Type:     Type,
	})
}

// SendMessage 发送消息
func SendMessage(addr string, data []byte) {
	// 给当前用户发送消息
	if addr != "" {
		user := ConnGroupMap[addr]
		user.Conn.WriteMessage(websocket.TextMessage, data)
		return
	}
	// 给所有人发送消息
	for _, user := range ConnGroupMap {
		user.Conn.WriteMessage(websocket.TextMessage, data)
	}
}
