package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	setting, ok := SettingsBindUriMap(c)
	if !ok {
		return
	}
	BY.JSON(c, 0, setting, "查询成功")
}
