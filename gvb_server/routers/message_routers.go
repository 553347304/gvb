package routers

import (
	"gvb_server/BAIYIN/JWT"
	"gvb_server/api"
)

func (router RouterGroup) MessageRouter() {
	app := api.GroupAppApi.MessageApi
	const index = "message"
	router.GET(index+"_user_list/:id", JWT.Auth(), app.MessageUserListView)
	router.GET(index, JWT.Auth(), app.MessageListView)
	router.POST(index, JWT.Auth(), app.MessageCreateView)
}
