package FK_ES

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"strings"
)

func _Query(o FK_Struct.PageInfo) elastic.Query {
	if o.Search != "" && o.Key != "" {
		switch o.Source {
		case "like":
			return elastic.NewWildcardQuery(o.Search, fmt.Sprintf("*%v*", o.Key)) // 模糊匹配
		case "match":
			return elastic.NewMatchQuery(o.Search, o.Key).Operator("and") // 精确匹配
		}
		return elastic.NewTermQuery(o.Search, o.Key) // 精确匹配
	}
	return elastic.NewMatchAllQuery() // 查询全部
}

// 分页查询
func _Line(es *elastic.SearchService, o FK_Struct.PageInfo) *elastic.SearchService {
	o.Param()
	es = es.Query(_Query(o)).From(o.Page).Size(o.Limit)
	// 排序
	if o.Sort != "" {
		split := strings.Split(o.Sort, " ")
		return es.Sort(split[0], split[1] == "asc")
	}
	return es
}

// Parse 解析ES		返回全部  map[string]any{}
func Parse[T any](m T, result *elastic.SearchResult) (R TypeResponseList[T]) {
	R.Count = result.TotalHits() // 总数
	R.List = make([]T, 0)
	for _, hit := range result.Hits.Hits {
		var m T
		// BY.Log(fmt.Sprintf("%s", hit.Source))
		BY.Json.Unmarshal(hit.Source, &m)
		R.List = append(R.List, m)
	}
	return R
}
