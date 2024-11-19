package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
)

func (Search) UpdateView(c *gin.Context) {
	var id FK_Struct.ID
	var cr ArticleRequest
	FK.Gin.BindQuery(c, &id)
	FK.Gin.BindJSON(c, &cr)

	res := FK.ES.ID(ct.ArticleIndex, id.ID)
	if !res.Exist {
		BY.LogE("ID不存在")
		return
	}

	S := BY.Map.StructMapDelEmpty(cr)
	BY.Log(S)

	// IsMap(Now, New)

	// s := BY.Type.StructMap(res.Model)

	// BY.Log(s)

	// FK.ES.Del(ArticleIndex, []string{id.ID}) // 删除文章

}
