package ser_es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
)

func getHead(head string) string {
	head = BY.AT.Replace(head, "#", "")
	head = BY.AT.Trim(head)
	return head
}
func getBody(head string) string {
	unsafe := BY.Type.HtmlDoc(head)
	doc := BY.Type.DocHtml(unsafe)
	return doc
}

func GetHtml(_id, _title, _content string) SearchData {
	dataList := BY.AT.Split(_content, "\n")
	var isCode = false
	var title, content, body string
	title = getHead(_title)
	for _, s := range dataList {
		if BY.AT.IsHead(s, "```") {
			isCode = !isCode
		}
		if BY.AT.IsHead(s, "#") && !isCode {
			title = "\n" + getHead(s)
			content = "\n" + getBody(body)
			body = ""
			continue
		}
		body += s
	}
	return SearchData{
		Key:     _id,
		Title:   title,
		Content: "\n" + getBody(body),
		Slug:    _id + "#" + content,
	}
}

// DelID 删除文章数据
func DelID(index string, id string) int64 {
	boolSearch := elastic.NewTermQuery("key", id)
	res, _ := global.ES.DeleteByQuery().Index(index).Query(boolSearch).Do(context.Background())
	return res.Deleted
}
