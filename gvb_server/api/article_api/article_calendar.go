package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/ct"
	"gvb_server/INIT/global"
	"time"
)

type BucketsType struct {
	Buckets []struct {
		Day      string `json:"key_as_string"`
		Key      int64  `json:"key"`
		DocCount int    `json:"doc_count"`
	} `json:"buckets"`
}

var DateCount = map[string]int{}

func (ArticleApi) CalendarView(c *gin.Context) {
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0) // 一年前
	format := "2006-01-02 15:04:05"

	// 时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")

	// 查找大于xx小于xx之间的数据
	query := elastic.NewRangeQuery("created_at").Gte(aYearAgo.Format(format)).Lte(now.Format(format))

	result, err := global.ES.Search(ct.ArticleIndex).Query(query).
		Aggregation("calendar", agg).Size(0).Do(context.Background())
	if err != nil {
		BY.JSON(c, "查询失败")
		return
	}

	// 存储数据
	var data BucketsType
	_ = json.Unmarshal(result.Aggregations["calendar"], &data)
	for _, bucket := range data.Buckets {
		day := BY.AT.Split(bucket.Day, " ")
		DateCount[day[0]] = bucket.DocCount
	}

	// 遍历日期
	days := int(now.Sub(aYearAgo).Hours() / 24)
	var list = make([]interface{}, 0)
	for i := 0; i <= days; i++ {
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")
		list = append(list, []interface{}{day, DateCount[day]})
	}

	BY.JSONList(c, len(list), list)
}
