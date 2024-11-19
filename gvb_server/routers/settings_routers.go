package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.GroupAppApi.SettingsApi
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}
