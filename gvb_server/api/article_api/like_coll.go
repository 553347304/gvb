package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/BAIYIN/System/SS"
	"gvb_server/INIT/ct"
	"gvb_server/models"
)

func (LikeApi) CollView(c *gin.Context) {
	var cr FK_Struct.ID

	if FK.Gin.BindJSON(c, &cr) {
		claims := JWT.Token(c)

		if !FK.ES.ID(ct.ArticleIndex, cr.ID).Exist {
			BY.JSON(c, msg+"ID不存在")
			return
		}

		var coll models.UserCollectModel

		is := FK_SQL.Is(&coll, "user_id = ? and article_id = ?", claims.UserID, cr.ID)
		BY.Log(is)
		var num = -1
		if !is {
			FK_SQL.Create(&models.UserCollectModel{
				UserID:    claims.UserID,
				ArticleID: cr.ID,
			})
			num = 1
		}

		FK_SQL.Delete(&coll)

		id := FK.ES.ID(ct.ArticleIndex, cr.ID)
		if !id.Exist {
			BY.JSON(c, msg+"ID不存在")
			return
		}
		// 更新文章收藏数
		FK.ES.Update(ct.ArticleIndex, cr.ID, map[string]any{
			"collects_count": id.Model.CollectsCount + num,
		})

		message := SS.IF(num == 1, msg+"收藏成功", "取消收藏成功")
		BY.JSON(c, message)
	}
}
