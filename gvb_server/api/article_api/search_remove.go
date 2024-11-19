package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
	"gvb_server/service/ser_es"
)

func (Search) RemoveView(c *gin.Context) {
	var cr FK_Struct.ID
	if FK.Gin.BindJSON(c, &cr) {
		count := ser_es.DelID(ct.SyncArticleIndex, cr.ID)
		BY.Log(BY.AT.Sp("成功删除 %d 条记录", count))
	}
}
