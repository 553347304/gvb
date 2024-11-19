package global

import (
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb_server/INIT/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	Redis  *redis.Client
	ES     *elastic.Client
)
