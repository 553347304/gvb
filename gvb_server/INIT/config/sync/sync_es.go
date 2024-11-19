package sync

import (
	"context"
	"github.com/olivere/elastic/v7"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/INIT/ct"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

func ES() {
	result, err := global.ES.Search(ct.ArticleIndex).Query(elastic.NewMatchAllQuery()).Size(10000).Do(context.Background())

	if err != nil {
		BY.LogE("查询失败： " + err.Error())
		return
	}

	likeInfo := FK.Redis.HGetAll(ct.ArticleLike) // 点赞量
	lookInfo := FK.Redis.HGetAll(ct.ArticleLook) // 浏览量
	commInfo := FK.Redis.HGetAll(ct.CommentLike) // 浏览量

	for _, hit := range result.Hits.Hits {
		var model models.ArticleModel
		BY.Json.Unmarshal(hit.Source, &model)

		Map := map[string]any{
			"digg_count":    model.DiggCount + likeInfo[hit.Id],
			"look_count":    model.LookCount + lookInfo[hit.Id],
			"comment_count": model.CommentCount + commInfo[hit.Id],
		}
		FK.ES.Update(ct.ArticleIndex, hit.Id, Map)
		BY.Log(BY.AT.Sp("%v  %v", model.Title, Map))
	}
	FK.Redis.Del(ct.ArticleLike) // 删除
	FK.Redis.Del(ct.ArticleLook) // 删除
	FK.Redis.Del(ct.CommentLike) // 删除
	BY.Log("ES同步成功")
}
