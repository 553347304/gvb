package INIT

import (
	"gvb_server/INIT/global"
)

const (
	SystemLog  = "release" // 运行模式支持：gin.ReleaseMode|"release"|"test"|"debug"
	ConfigFile = "settings.yaml"
)

func INIT() {
	Yaml()                 // 初始化配置
	global.Log = Logger()  // 初始化日志
	global.DB = MySQL()    // 初始化数据库
	global.Redis = Redis() // 初始化Redis
	global.ES = Es()       // 初始化ES
}
