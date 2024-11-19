package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/ct"
	"gvb_server/models"
)

type CollResponse struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

func (LikeApi) CollListView(c *gin.Context) {
	var cr FK_Struct.Mysql

	if FK.Gin.BindQuery(c, &cr) {
		claims := JWT.Token(c)

		result := FK_SQL.GetList(models.UserCollectModel{UserID: claims.UserID}, cr)

		var collMap = map[string]string{}
		var IDList []interface{}
		for _, model := range result {
			IDList = append(IDList, model.ArticleID)
			collMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
		}

		idList, err := FK.ES.FindIDList(ct.ArticleIndex, IDList)
		if err != nil {
			BY.JSON(c, err.Error())
			return
		}
		var list = make([]CollResponse, 0)
		for _, hit := range idList.Hits.Hits {
			var model models.ArticleModel
			if !BY.Json.Unmarshal(hit.Source, &model) {
				continue
			}
			model.ID = hit.Id
			list = append(list, CollResponse{
				ArticleModel: model,
				CreatedAt:    collMap[hit.Id],
			})
		}
		BY.JSONList(c, len(list), list)
	}
}
