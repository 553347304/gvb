package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/INIT"
	"gvb_server/INIT/config"
	"gvb_server/INIT/global"
	"path"
)

const docsPath = "file/docs"

func (BigModelApi) ModelSettingView(c *gin.Context) {
	filePath := path.Join(docsPath, global.Config.BigModel.Setting.Name+".md")
	global.Config.BigModel.Setting.Help = BY.File.Read(filePath)
	BY.JSON(c, global.Config.BigModel)
}

func (BigModelApi) ModelSettingUpdateView(c *gin.Context) {
	var cr config.BigModel
	if FK.Gin.BindJSON(c, &cr) {
		if cr.Session != nil {
			global.Config.BigModel.Session = cr.Session
		}
		if cr.Setting != nil {
			global.Config.BigModel.Setting = cr.Setting
		}
		if cr.ModelList != nil && len(*cr.ModelList) != 0 {
			global.Config.BigModel.ModelList = cr.ModelList
		}
		INIT.SetYaml()
		BY.JSON(c, "更新成功")
	}
}
