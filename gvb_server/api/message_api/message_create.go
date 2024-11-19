package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

type MessageModel struct {
	SendUserID uint   `json:"send_user_id" binding:"required" msg:"不能为空"` // 发送人id
	RevUserID  uint   `json:"rev_user_id" binding:"required" msg:"不能为空"`  // 接收人id
	Content    string `json:"content" binding:"required" msg:"不能为空"`      // 消息内容
}

func (MessageApi) MessageCreateView(c *gin.Context) {
	var cr MessageModel
	if FK.Gin.BindJSON(c, &cr) {
		var send, receive models.UserModel
		if !FK_SQL.Is(&send, cr.SendUserID) {
			BY.JSON(c, "发送人不存在")
			return
		}
		if !FK_SQL.Is(&receive, cr.RevUserID) {
			BY.JSON(c, "接收人不存在")
			return
		}
		res := FK_SQL.Create(&models.MessageModel{
			SendUserID:       cr.SendUserID,
			SendUserNickName: send.NickName,
			SendUserAvatar:   send.Avatar,
			RevUserID:        cr.RevUserID,
			RevUserNickName:  receive.NickName,
			RevUserAvatar:    receive.Avatar,
			IsRead:           false,
			Content:          cr.Content,
		})
		BY.JSON(c, "创建"+msg, res)
	}
}
