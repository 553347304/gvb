package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/models"
)

type userSendList struct {
	UserID   uint   `json:"user_id"`
	NikeName string `json:"nike_name"`
	Count    int64  `json:"total"`
	Avatar   string `json:"avatar"`
}

// MessageUserListView 查询发送过消息的用户
func (MessageApi) MessageUserListView(c *gin.Context) {
	var o FK_Struct.Mysql
	if FK.Gin.BindQuery(c, &o) {
		var group = make([]userSendList, 0)
		FK_SQL.FindGroup(models.MessageModel{}, o, &group,
			"send_user_id as user_id,send_user_nick_name as nike_name,send_user_avatar as avatar,count(*) as count")

		// 去除传参的用户
		id := FK.Gin.Param(c, "id")
		uid := BY.Type.StringUint(id)
		var list = make([]userSendList, 0)
		for _, k := range group {
			if k.UserID == uid {
				continue
			}
			list = append(list, k)
		}
		BY.JSONList(c, len(group), list)
	}
}
