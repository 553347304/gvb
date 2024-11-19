package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuModel
	if FK.Gin.BindJSON(c, &cr) {
		is := FK_SQL.Get(models.MenuModel{}, "title", cr.Title, "title")
		if is != "" {
			BY.JSON(c, 1000, "重复菜单")
			return
		}

		BY.JSON(c, "创建菜单", create(cr))
	}
}
