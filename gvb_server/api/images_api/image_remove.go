package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/models"
)

// ImageRemoveView 批量删除
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr FK_Struct.UID
	if FK.Gin.BindJSON(c, &cr) {
		R := FK_SQL.Remove(&[]models.BannerModel{}, cr.IDList)
		BY.JSON(c, R.Code, R.Message)
	}
}
