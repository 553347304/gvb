package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
)

type SettingsApi struct {
}

type SettingsUri struct {
	Name string `uri:"name" binding:"required"`
}

func SettingsBindUriMap(c *gin.Context) (any, bool) {
	Map := map[string]any{
		"site":  &global.Config.SiteInfo,
		"email": &global.Config.Email,
		"qq":    &global.Config.QQ,
	}

	var cr SettingsUri
	// 查询URI参数
	err := c.ShouldBindUri(&cr)
	if err != nil {
		BY.JSON(c, 1001, err)
		return nil, false
	}
	// 查询配置
	setting, ok := Map[cr.Name]
	if !ok {
		BY.JSON(c, 1000, "没有对应的配置信息")
		return nil, false
	}
	return setting, true
}
