package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/models"
)

// AdvertListView 列表查询
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr FK_Struct.Mysql
	if FK.Gin.BindQuery(c, &cr) {
		// 判断 Referer 是否包含 admin  返回: 全部/false
		IsShow := BY.AT.Find(FK.Gin.GetHead(c, "Referer"), "admin")
		list := FK_SQL.GetList(models.AdvertModel{IsShow: IsShow}, cr)
		BY.JSONList(c, len(list), list)
	}
}
