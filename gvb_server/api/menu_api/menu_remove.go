package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/models"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	if FK.Gin.BindJSON(c, &cr) {
		if !Remove(c, cr.IDList) {
			return
		}
		BY.JSON(c, "删除成功")
	}
}
