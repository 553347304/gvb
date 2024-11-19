package routers

import "gvb_server/api"

func (router RouterGroup) TagRouter() {
	app := api.GroupAppApi.TagApi
	const index = "tag"
	router.POST(index, app.TagCreateView)
	router.GET(index, app.TagListView)
	router.PUT(index+"/:id", app.TagUpdateView)
	router.DELETE(index, app.TagRemoveView)
}
