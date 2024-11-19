package ser_es

type SearchData struct {
	Key     string `json:"key"`
	Title   string `json:"title"`   // 标题
	Slug    string `json:"slug"`    // 包含文章的id 的跳转地址
	Content string `json:"content"` // 正文
}
