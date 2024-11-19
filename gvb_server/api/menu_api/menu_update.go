package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuModel
	if FK.Gin.BindJSON(c, &cr) {
		id := FK.Gin.Param(c, "id")
		uid := BY.Type.StringUint(id)
		if !Remove(c, []uint{uid}) {
			return
		}
		BY.JSON(c, "更新菜单", create(cr))
	}
}
