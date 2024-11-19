package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/models"
)

// TagListView 列表查询
func (TagApi) TagListView(c *gin.Context) {
	var cr FK_Struct.Mysql
	if FK.Gin.BindJSON(c, &cr) {
		list := FK_SQL.GetList(models.TagModel{}, cr)
		BY.JSONList(c, len(list), list)
	}
}
