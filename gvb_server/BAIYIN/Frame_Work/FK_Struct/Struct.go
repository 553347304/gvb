package FK_Struct

type UID struct {
	ID     uint   `json:"id" form:"id" uri:"id"`
	IDList []uint `json:"id_list"`
}

type ID struct {
	ID     string   `json:"id" form:"id" uri:"id"`
	IDList []string `json:"id_list"`
}

type PageInfo struct {
	Search string `form:"search"` // 查询字段
	Key    string `form:"key" `   // 关键字
	Page   int    `form:"page"`   // 当前页数
	Limit  int    `form:"limit"`  // 每页条数
	Sort   string `form:"sort"`   // 排序---从后往前:"_ desc"   从前往后:"_ asc"
	Source string `form:"source"` // 来源
}

func (o *PageInfo) Param() {
	// 默认条数
	if o.Limit == 0 {
		o.Limit = 100
	}
	// 默认页数
	if o.Page != 0 {
		o.Page = (o.Page - 1) * o.Limit
	}

	if o.Sort == "" {
		o.Sort = "created_at desc"
	}
}

type Mysql struct {
	PageInfo
	Preload []string
	Where   []interface{}
}

type EsPageInfo struct {
	PageInfo
	SortField string `json:"sort_field"`
	Fields    []string
	Tag       string
	SortBool  bool
}

func (p *EsPageInfo) EsParam() {
	p.Param()
	p.SortBool = false
	if p.Sort == "created_at asc" {
		p.SortBool = true
	}
	if p.SortField == "" {
		p.SortField = "created_at"
	}
}

type Addr struct {
	IP   string
	Addr string
}

type Options[T any] struct {
	Label string `json:"label"`
	Value T      `json:"value"`
}
