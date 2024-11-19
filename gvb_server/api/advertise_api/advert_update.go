package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

type AdvertModel struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`       // 显示的标题
	Href   string `json:"href" binding:"required"`                    // 跳转链接
	Images string `json:"images" binding:"required,uri" msg:"非法图片地址"` // 图片
	IsShow bool   `json:"is_show"`                                    // 是否展示
}

// AdvertUpdateView 更新
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	var cr AdvertModel
	if FK.Gin.BindJSON(c, &cr) {
		Map := BY.Type.StructMap(cr)
		var update models.AdvertModel
		if !FK_SQL.Is(&update, FK.Gin.Param(c, "id")) {
			BY.JSON(c, "id不存在")
			return
		}
		res := FK_SQL.Update(&update, Map)
		FK.JSON(c, res, "更新"+msg, map[string]any{})
	}
}
