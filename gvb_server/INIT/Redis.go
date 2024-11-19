package INIT

import (
	"context"
	"github.com/go-redis/redis"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
	"time"
)

func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.System.Redis,
		Password: "baiyin", // no password set
		DB:       0,        // use default DB
		PoolSize: 100,      // 连接池大小
	})

	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		BY.LogE("Redis连接失败: " + err.Error())
	}
	return rdb
}
