package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"path"
)

func (BigModelApi) IconView(c *gin.Context) {
	const _path = "file/icon/role"
	dir := BY.File.Dir(_path)
	if dir == nil {
		return
	}

	var list = make([]FK_Struct.Options[string], 0)
	for _, entry := range dir {
		key := "/" + path.Join(_path, entry.Name())
		list = append(list, FK_Struct.Options[string]{
			Label: key,
			Value: key,
		})
	}
	BY.JSONList(c, len(list), list)
}
