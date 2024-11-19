package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/ct"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

func (LikeApi) CollListRemoveView(c *gin.Context) {
	var cr FK_Struct.ID

	if FK.Gin.BindJSON(c, &cr) {
		claims := JWT.Token(c)

		var coll []models.UserCollectModel
		var articleIDList []string

		global.DB.Find(&coll, "user_id = ? and article_id in ?", claims.UserID, cr.IDList).
			Select("article_id").Scan(&articleIDList)

		if len(articleIDList) == 0 {
			BY.JSON(c, "请求非法")
			return
		}
		var IDList []interface{}
		for _, s := range articleIDList {
			IDList = append(IDList, s)
		}
		// 更新文章数
		result, err := FK.ES.FindIDList(ct.ArticleIndex, IDList)
		if err != nil {
			BY.JSON(c, err.Error())
			return
		}

		for _, hit := range result.Hits.Hits {
			var model models.ArticleModel
			if !BY.Json.Unmarshal(hit.Source, &model) {
				continue
			}
			count := model.CollectsCount - 1

			// 更新文章收藏数
			res := FK.ES.Update(ct.ArticleIndex, cr.ID, map[string]any{
				"collects_count": count,
			})
			if !res {
				continue
			}
		}
		FK_SQL.Delete(&coll)
		BY.JSON(c, BY.AT.Sp("成功取消收藏 %d 篇文章", len(articleIDList)))
	}
}
