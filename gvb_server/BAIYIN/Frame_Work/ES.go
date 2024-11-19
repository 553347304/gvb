package FK

import (
	"context"
	"github.com/olivere/elastic/v7"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

const index = "index" // ES索引名
const param = `
{
 "settings": { "index":{ "max_result_window": "100000" } },
 "mappings": {
   "properties": {
     "id": { "type": "integer" },
     "title": { "type": "text" },
     "created_at":{
       "type": "date",
       "null_value": "null",
       "format": "[yyyy-MM-dd HH:mm:ss]"
     }
   }
 }
}
`

// existIndex 索引是否存在
func existIndex(index string) bool {
	is, _ := global.ES.IndexExists(index).Do(context.Background())
	return is
}

// CreateIndex 创建索引
func (fk _ES) CreateIndex(index, param string) bool {
	if existIndex(index) {
		fk.DelIndex(index)
	}
	r, _ := global.ES.CreateIndex(index).BodyString(param).Do(context.Background())
	BY.Log("创建索引: " + index)
	return r.Acknowledged
}

// DelIndex 删除索引
func (_ES) DelIndex(index string) bool {
	BY.Log("删除索引: " + index)
	if existIndex(index) {
		r, _ := global.ES.DeleteIndex(index).Do(context.Background())
		return r.Acknowledged
	}
	return false
}

type _ID struct {
	Model models.ArticleModel
	Exist bool
}

// ID 根据ID查找
func (_ES) ID(index string, id string) _ID {
	var cr models.ArticleModel
	res, err := global.ES.Get().Index(index).Id(id).Do(context.Background())
	if err != nil {
		return _ID{Model: cr, Exist: false}
	}
	BY.Json.Unmarshal(res.Source, &cr)
	cr.ID = res.Id
	return _ID{Model: cr, Exist: true}
}

func (_ES) FindLight(index string, query *elastic.BoolQuery, light string) (*elastic.SearchResult, error) {
	return global.ES.Search(index).Query(query).
		Highlight(elastic.NewHighlight().Field(light)). // 高亮字段
		Size(1000).Do(context.Background())
}

// Find 查找
func (fk _ES) Find(index string, cr FK_Struct.EsPageInfo) (List []models.ArticleModel, count int) {
	cr.EsParam()
	List = make([]models.ArticleModel, 0)
	search := elastic.NewBoolQuery()
	if cr.Key != "" {
		search.Must(elastic.NewMultiMatchQuery(cr.Key, cr.Fields...))
	}
	if cr.Tag != "" {
		search.Must(elastic.NewMultiMatchQuery(cr.Tag, "tags"))
	}
	// 关键词高亮   NewHighlight
	res, err := global.ES.Search(index).Query(search).Highlight(elastic.NewHighlight().Field("title")).
		From(cr.Page).Sort(cr.SortField, cr.SortBool).Size(cr.Limit).Do(context.Background())
	if err != nil {
		return
	}
	// 找到条数
	count = int(res.Hits.TotalHits.Value)
	for _, hit := range res.Hits.Hits {
		var model models.ArticleModel
		if !BY.Json.Unmarshal(hit.Source, &model) {
			continue
		}
		title, ok := hit.Highlight["title"]
		if ok {
			model.Title = title[0]
		}
		model.ID = hit.Id // ES ID

		List = append(List, model)
	}

	return
}

func (_ES) Search(index string) (*elastic.SearchResult, error) {
	Search := elastic.NewMatchAllQuery()
	return global.ES.Search(index).Query(Search).Size(1000).Do(context.Background())
	// 总数
	// int(result.Hits.TotalHits.Value)
}

// Count 查询总数
func (fk _ES) Count(index string) int {
	result, _ := fk.Search(index)
	return int(result.Hits.TotalHits.Value)
}

func (_ES) FindIDList(index string, IDList []interface{}) (*elastic.SearchResult, error) {
	Search := elastic.NewTermsQuery("_id", IDList...)
	return global.ES.Search(index).Query(Search).Size(1000).Do(context.Background())
}

// GetCount 获取总数
func (_ES) GetCount(index string, name string) int64 {
	// NewValueCountAggregation	  不去重
	// NewCardinalityAggregation  去重
	result, _ := global.ES.Search(index).
		Aggregation(name, elastic.NewValueCountAggregation().
			Field(name)).Size(0).Do(context.Background())
	cTag, _ := result.Aggregations.Cardinality(name)

	return int64(*cTag.Value)
}

// Create 创建
func (_ES) Create(index string, obj any) (*elastic.IndexResponse, bool) {
	res, err := global.ES.Index().Index(index).BodyJson(obj).Refresh("true").Do(context.Background())
	return res, err == nil
}

// Update 更新
func (_ES) Update(index string, id string, data map[string]any) bool {
	_, err := global.ES.Update().Index(index).Id(id).Doc(data).Do(context.Background())
	return err == nil
}

// Del 批量删除
func (fk _ES) Del(index string, IDList []string) int {
	bulkService := global.ES.Bulk().Index(index).Refresh("true")
	for _, id := range IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	result, _ := bulkService.Do(context.Background()) // 执行
	return len(result.Succeeded())                    // 删除个数
}
