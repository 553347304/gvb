package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

type ImageUpdate struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

// ImageUpdateView 更新
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdate
	if FK.Gin.BindJSON(c, &cr) {
		var update models.BannerModel
		if !FK_SQL.Is(&update, cr.ID) {
			BY.JSON(c, "id不存在")
			return
		}
		ok := FK_SQL.Update(&update, map[string]any{"name": cr.Name})
		BY.JSON(c, "更新"+msg, ok)
	}
}
