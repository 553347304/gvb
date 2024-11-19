package routers

import (
	"gvb_server/BAIYIN/JWT"
	"gvb_server/api"
)

func (router RouterGroup) CommentRouter() {
	app := api.GroupAppApi.CommentsApi
	const index = "comment"
	router.POST(index, JWT.Auth(), app.CreateView)
	router.GET(index, app.ListView)
	router.POST(index+"/:id", app.LikeView)
	router.DELETE(index, app.RemoveView)

}
