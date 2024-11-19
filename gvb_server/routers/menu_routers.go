package routers

import "gvb_server/api"

func (router RouterGroup) MenuRouter() {
	app := api.GroupAppApi.MenuApi
	const index = "menu"
	router.POST(index, app.MenuCreateView)
	router.GET(index, app.MenuListView)
	router.PUT(index+"/:id", app.MenuUpdateView)
	router.DELETE(index, app.MenuRemoveView)
}
