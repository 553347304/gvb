package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_ES"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
)

func (ArticleApi) RemoveView(c *gin.Context) {
	var cr FK_Struct.ID
	if FK.Gin.BindJSON(c, &cr) {
		R := FK_ES.Remove(ct.ArticleIndex, cr.IDList, FK_ES.ID)
		FK_ES.Remove(ct.SyncArticleIndex, cr.IDList, FK_ES.KEY)
		BY.JSONList(c, int(R.Count), R.List, "删除成功")
	}
}
