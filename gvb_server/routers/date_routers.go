package routers

import "gvb_server/api"

func (router RouterGroup) DateRouter() {
	app := api.GroupAppApi.DateApi
	const index = "date"
	router.GET(index+"_list", app.DateView)
	router.GET(index+"_count", app.CountView)
}
