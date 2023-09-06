package core

import (
	"fmt"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/Ghostbb-io/g-api/pkg/redisx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"os"
)

func InitRedis() *redisx.Redis {
	_redis, err := redisx.NewRedis(&redis.Options{
		Addr:     global.GB_CONFIG.Redis.Addr,
		Password: global.GB_CONFIG.Redis.Password,
		DB:       global.GB_CONFIG.Redis.DB, // use default DB
	})
	if err != nil {
		global.GB_LOG.Error("Redis connect ping failed, err:", zap.Error(err))
		if !global.GB_CONFIG.Zap.LogInConsole {
			fmt.Println("Redis connect ping failed, err:", zap.Error(err))
		}
		os.Exit(0)
	}
	global.GB_LOG.Info("Redis connect success")
	return _redis
}
