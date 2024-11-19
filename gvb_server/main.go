package main

import (
	"gvb_server/BAIYIN/System/Time"
	"gvb_server/INIT"
	sync2 "gvb_server/INIT/config/sync"
	"gvb_server/INIT/flag"
	"gvb_server/routers"
)

func TimeTask() {
	Cron := Time.Cron()
	Cron.AddFunc("0 0 0 * * *", sync2.ES)
	Cron.AddFunc("0 0 0 * * *", sync2.Mysql)
	Cron.Start()
}

// go run main.go -db 迁移表结构
// go run main.go -es c/d			ES创建/删除
func main() {
	INIT.INIT()      // 初始化
	flag.SwitchCmd() // 命令行参数绑定
	go TimeTask()    // 定时同步任务
	// test.Main()
	routers.InitRouter() // 初始化路由
}
