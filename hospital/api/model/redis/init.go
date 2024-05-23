package redis

import (
	"demo/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

// 初始化redis
func init() {
	Client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.GlobalConfig.Redis.Host, config.GlobalConfig.Redis.Port),
		DB:   config.GlobalConfig.Redis.Db,
	})
}
