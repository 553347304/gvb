package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT"
	"gvb_server/INIT/global"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	setting, ok := SettingsBindUriMap(c)
	if !ok {
		return
	}

	// 处理 JSON 绑定到配置变量上
	err := c.ShouldBindJSON(setting)
	if err != nil {
		global.Log.Error(err)
		BY.JSON(c, 1000, err.Error())
		return
	}

	BY.JSON(c, 0, setting, "配置文件修改成功")
	INIT.SetYaml() // 修改 YAML 文件
}
