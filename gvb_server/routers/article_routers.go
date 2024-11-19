package routers

import (
	"gvb_server/BAIYIN/JWT"
	"gvb_server/api"
)

func (router RouterGroup) ArticleRouter() {
	app := api.GroupAppApi.ArticleApi
	const index = "article"
	router.POST(index, JWT.Auth(), app.CreateView)
	router.DELETE(index, app.RemoveView)
	router.PUT(index+"/:id", app.UpdateView)
	router.GET(index, app.ListView)

	router.GET("article_search", app.ArticleSearchView)
	router.GET("articles/calendar", app.CalendarView)
	router.GET("articles/:id", app.DetailView)

	like := api.GroupAppApi.LikeApi
	router.POST("likes", like.LikeView)
	router.POST("likes/coll", JWT.Auth(), like.CollView)
	router.GET("likes/coll/list", JWT.Auth(), like.CollListView)
	router.DELETE("likes/coll/list", JWT.Auth(), like.CollListRemoveView)
	search := api.GroupAppApi.Search
	router.GET("search", search.FullView)
	router.DELETE("search", search.RemoveView)
	router.PUT("search", search.UpdateView)

}
