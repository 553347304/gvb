package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
)

// go get -u github.com/gin-gonic/gin

type RouterGroup struct {
	// *gin.Engine
	*gin.RouterGroup
}

func InitRouter() {
	gin.SetMode(global.Config.System.Log) // 运行模式支持：gin.ReleaseMode|"release"|"test"|"debug"
	router := gin.Default()
	apiRouterGroup := router.Group("api")

	router.Static("file", "file") // 静态文件

	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingsRouter()
	routerGroupApp.ImagesRouter()
	routerGroupApp.AdvertiseRouter()
	routerGroupApp.MenuRouter()
	routerGroupApp.UserRouter()
	routerGroupApp.TagRouter()
	routerGroupApp.MessageRouter()
	routerGroupApp.ArticleRouter()
	routerGroupApp.CommentRouter()
	routerGroupApp.ChatRouter()
	routerGroupApp.DateRouter()
	routerGroupApp.BigModelRouter()

	err := router.Run("0.0.0.0:" + global.Config.System.Port)
	BY.LogE("router启动失败：%v" + err.Error())
}
