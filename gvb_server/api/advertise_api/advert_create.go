package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr models.AdvertModel
	if FK.Gin.BindJSON(c, &cr) {
		if FK_SQL.Is(&models.AdvertModel{IsShow: false}, map[string]any{"title": cr.Title}) {
			BY.JSON(c, msg+"已存在")
			return
		}

		res := FK_SQL.Create(&models.AdvertModel{
			Title:  cr.Title,
			Href:   cr.Href,
			Images: cr.Images,
			IsShow: cr.IsShow,
		})
		BY.JSON(c, "创建"+msg, res)
	}
}
