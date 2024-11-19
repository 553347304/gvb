package flag

import (
	sys_flag "flag"
	SM "gvb_server/BAIYIN/System"
)

// SwitchCmd 根据命令执行不同的函数
func SwitchCmd() {
	DB := sys_flag.Bool("db", false, "初始化数据库")
	ES := sys_flag.String("es", "", "es操作")
	sys_flag.Parse() // 解析命令行参数写入注册的flag里

	if *DB {
		MigrationTable()
		SM.SYS.Exit()
	}
	if *ES != "" {
		Es(*ES)
		SM.SYS.Exit()
	}
}
