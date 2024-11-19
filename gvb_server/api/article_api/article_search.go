package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_ES"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
)

// ArticleSearchView 聚合搜索 o.Search
func (ArticleApi) ArticleSearchView(c *gin.Context) {
	var o FK_Struct.PageInfo
	if FK.Gin.BindQuery(c, &o) {
		R := FK_ES.FindGroup(ct.ArticleIndex, o)
		BY.JSONList(c, int(R.Count), R.List, R.Message)
	}
}
