package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
)

func (ArticleApi) DetailView(c *gin.Context) {
	var cr FK_Struct.ID
	if FK.Gin.BindURI(c, &cr) {
		id := FK.ES.ID(ct.ArticleIndex, cr.ID)
		if !id.Exist {
			BY.JSON(c, msg+"ID不存在")
			return
		}
		BY.JSON(c, msg+"详情", id.Model)
		SetLike(c, ct.ArticleLook, cr) // 浏览量
	}
}
