package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_ES"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
	"gvb_server/models"
	"gvb_server/service/ser_es"
)

func (ArticleApi) UpdateView(c *gin.Context) {

	var id FK_Struct.ID
	if !FK.Gin.BindURI(c, &id) {
		return
	}

	var cr ArticleRequest
	if FK.Gin.BindJSON(c, &cr) {
		cr.BannerURL = FK_SQL.Get(models.BannerModel{}, "id", cr.BannerID, "path")
		R := FK_ES.Update(ct.ArticleIndex, id.ID, BY.Map.StructMapDelEmpty(cr))
		// 同步数据
		data := ser_es.GetHtml(id.ID, cr.Title, cr.Content)
		ID := FK_ES.Get(ct.SyncArticleIndex, "key", id.ID, FK_ES.ID)
		FK_ES.Update(ct.SyncArticleIndex, ID, BY.Map.StructMapDelEmpty(data))
		BY.JSON(c, msg+"更新", R)
	}
}
