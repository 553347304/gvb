package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
)

func (LikeApi) LikeView(c *gin.Context) {
	var cr FK_Struct.ID
	if FK.Gin.BindJSON(c, &cr) {
		SetLike(c, ct.ArticleLike, cr)
		BY.JSON(c, "点赞成功", []string{cr.ID})
	}
}
