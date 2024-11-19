package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

// TagUpdateView 更新
func (TagApi) TagUpdateView(c *gin.Context) {
	var cr TagRequest
	if FK.Gin.BindJSON(c, &cr) {
		var update models.TagModel
		if !FK_SQL.Is(&update, FK.Gin.Param(c, "id")) {
			BY.JSON(c, "id不存在")
			return
		}

		res := FK_SQL.Update(&update, BY.Map.StructMapDelEmpty(cr))
		BY.JSON(c, "更新"+msg, res)
	}
}
