package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_ES"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
	"gvb_server/models"
)

func (ArticleApi) ListView(c *gin.Context) {
	var o FK_Struct.PageInfo
	if FK.Gin.BindQuery(c, &o) {
		R := FK_ES.Find(ct.ArticleIndex, models.ArticleModel{}, o)
		for i := range R.List {
			R.List[i].Title = FK_ES.Highlight(R.List[i].Title, "title", o)
		}
		BY.JSONList(c, int(R.Count), R.List)
	}
}
