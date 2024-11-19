package INIT

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
	"time"
)

// const name = "127.0.0.1:3306"

func MySQL() *gorm.DB {
	dsn := global.Config.System.Mysql + "?charset=utf8mb4&parseTime=True&loc=Local"
	var mysqlLogger = logger.Default.LogMode(logger.Error) // 日志等级
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		BY.LogE("MySQL连接失败: " + dsn)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout

	return db
}
