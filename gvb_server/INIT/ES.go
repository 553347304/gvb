package INIT

import (
	"github.com/olivere/elastic/v7"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
)

func Es() *elastic.Client {
	c, err := elastic.NewClient(
		elastic.SetURL(global.Config.System.Es),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		BY.LogE("ES连接失败: " + err.Error())
	}

	return c
}
