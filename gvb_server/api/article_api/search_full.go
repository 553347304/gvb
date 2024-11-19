package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_ES"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
	"gvb_server/service/ser_es"
)

func (Search) FullView(c *gin.Context) {
	var o FK_Struct.PageInfo
	if FK.Gin.BindQuery(c, &o) {
		const light = "body"

		boolQuery := elastic.NewBoolQuery()
		if o.Key != "" {
			boolQuery.Must(elastic.NewMultiMatchQuery(o.Key, "title", light))
		}

		result, err := FK.ES.FindLight(ct.SyncArticleIndex, boolQuery, light)
		if err != nil {
			return
		}

		total := int(result.Hits.TotalHits.Value)
		list := make([]ser_es.SearchData, 0)
		for _, hit := range result.Hits.Hits {
			var model ser_es.SearchData
			if !BY.Json.Unmarshal(hit.Source, &model) {
				continue
			}

			body, ok := hit.Highlight[light]
			if ok {
				model.Content = body[0]
			}

			list = append(list, model)
		}

		for i := range list {
			list[i].Title = FK_ES.Highlight(list[i].Title, "title", o)
		}
		BY.JSONList(c, total, list)
	}
}
