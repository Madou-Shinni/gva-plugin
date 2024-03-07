package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Config struct {
	Rdb *redis.Client
	Log *zap.Logger
}

var GlobalConfig = new(Config)
