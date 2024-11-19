package comment_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/models"
)

func (CommentApi) ListView(c *gin.Context) {
	var o FK_Struct.Mysql
	if FK.Gin.BindQuery(c, &o) {
		list := FK_SQL.GetList(models.CommentModel{}, o)
		BY.JSONList(c, len(list), list)
	}
	return
}