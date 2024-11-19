package routers

import "gvb_server/api"

func (router RouterGroup) AdvertiseRouter() {
	app := api.GroupAppApi.AdvertiseApi
	const index = "advert"
	router.POST(index, app.AdvertCreateView)
	router.GET(index, app.AdvertListView)
	router.PUT(index+"/:id", app.AdvertUpdateView)
	router.DELETE(index, app.AdvertRemoveView)
}
