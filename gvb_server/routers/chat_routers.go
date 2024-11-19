package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) ChatRouter() {
	app := api.GroupAppApi.ChatApi
	const index = "chat"
	router.GET("chat_group", app.GroupView)
	router.GET(index, app.GroupListView)
}
