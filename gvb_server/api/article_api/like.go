package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
)

func SetLike(c *gin.Context, index string, cr FK_Struct.ID) bool {
	if !FK.ES.ID(ct.ArticleIndex, cr.ID).Exist {
		BY.JSON(c, msg+"ID不存在")
		return false
	}
	num := FK.Redis.HGet(index, cr.ID)
	FK.Redis.HSet(index, cr.ID, num+1)
	BY.Log(index + "成功: " + cr.ID)
	return true
}
