package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	app := api.GroupAppApi.ImagesApi
	const index = "image"
	router.POST(index, app.ImageUploadView)
	router.DELETE(index, app.ImageRemoveView)
	router.PUT(index, app.ImageUpdateView)
	router.GET(index, app.ImageListView)
}
