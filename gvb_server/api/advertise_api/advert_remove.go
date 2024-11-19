package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

// AdvertRemoveView 批量删除
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	if FK.Gin.BindJSON(c, &cr) {
		R := FK_SQL.Remove(&[]models.AdvertModel{}, cr.IDList)
		BY.JSON(c, R.Code, R.Message)
	}
}
