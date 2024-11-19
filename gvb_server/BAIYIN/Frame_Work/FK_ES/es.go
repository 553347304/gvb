package FK_ES

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/BAIYIN/System/SS"
	"gvb_server/INIT/global"
	"strings"
)

// Highlight 搜索高亮	  for i := range List { List[i].Title = FK_ES.Highlight(List[i].Title, "title", o) }
func Highlight(s string, field string, o FK_Struct.PageInfo) string {
	if o.Key == "" || o.Search != field {
		return s
	}
	return strings.ReplaceAll(s,
		fmt.Sprintf("%s", o.Key),
		fmt.Sprintf("<em>%s</em>", o.Key))
}

func Find[T any](Index string, m T, o FK_Struct.PageInfo) (R TypeResponseList[T]) {
	es := global.ES.Search().Index(Index)
	result, _ := _Line(es, o).Do(context.Background()) // 执行请求
	return Parse(m, result)
}

// Get 获取一个字段值
func Get(Index, search, value, res string) string {
	var o = FK_Struct.PageInfo{Search: search, Key: value}
	result, _ := global.ES.Search().Index(Index).Query(_Query(o)).Do(context.Background()) // 执行请求
	var m map[string]any
	for _, hit := range result.Hits.Hits {
		// BY.Log(fmt.Sprintf("%s", hit.Id))
		BY.Json.Unmarshal(hit.Source, &m)

		if res == "" {
			res = search
		}

		switch res {
		case ID:
			return hit.Id
		}
		if m[res] != nil {
			return m[res].(string)
		}
	}
	return ""
}

// FindGroup 分组查询
func FindGroup(index string, o FK_Struct.PageInfo) (R TypeResponseList[interface{}]) {
	R.List = make([]interface{}, 0)
	query := elastic.NewTermsAggregation().Field(o.Search)
	result, err := global.ES.Search(index).Aggregation(o.Search, query).Size(0).Do(context.Background())
	if err != nil {
		R.Message = SS.IF(o.Search == "", "搜索条件不能为空", "ok")
		BY.LogE(err)
		return
	}
	term, _ := result.Aggregations.Terms(o.Search)
	for _, b := range term.Buckets {
		R.List = append(R.List, b.Key)
		// b.DocCount
	}
	R.Count = int64(len(R.List))
	return
}

// Create 创建
func Create(index string, obj any) string {
	result, err := global.ES.Index().Index(index).BodyJson(obj).Refresh("true").Do(context.Background())
	if err != nil {
		BY.LogE(err)
		return ""
	}
	return result.Id
}

func Remove[T any](Index string, List []T, name string) (R TypeResponseList[T]) {
	for _, v := range List {
		query := elastic.NewTermQuery(name, v) // 匹配一个字段 _id = ""
		global.ES.DeleteByQuery().Index(Index).Query(query).Do(context.Background())
	}
	R.Count = int64(len(R.List))
	R.List = List
	return R
}

func Update(Index string, Id string, data map[string]any) bool {
	_, err := global.ES.Update().Index(Index).Id(Id).Doc(data).Do(context.Background())
	return err == nil
}
