package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

// TagRemoveView 批量删除
func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	if FK.Gin.BindJSON(c, &cr) {
		R := FK_SQL.Remove(&[]models.TagModel{}, cr.IDList)
		BY.JSON(c, R.Code, R.Message)
	}
}
