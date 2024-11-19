package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/models"
)

type ID struct {
	FK_Struct.Mysql
	UserID    uint `form:"user_id"`
	ReceiveID uint `form:"receive_id"`
}

// MessageListView 列表查询
func (MessageApi) MessageListView(c *gin.Context) {
	var cr ID
	if FK.Gin.BindQuery(c, &cr) {
		list := FK_SQL.GetList(models.MessageModel{}, cr.Mysql)
		uID := cr.UserID + cr.ReceiveID
		var List = make([]models.MessageModel, 0)
		for _, k := range list {
			// 两个人对话的ID总和
			id := k.SendUserID + k.RevUserID
			if uID != id {
				continue
			}
			List = append(List, k)
		}

		BY.JSONList(c, len(list), list)
	}
}
