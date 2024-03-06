package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Config struct {
	Wechat *config.Wechat
	Rdb    *redis.Client
	Log    *zap.Logger
}

var GlobalConfig = new(Config)
