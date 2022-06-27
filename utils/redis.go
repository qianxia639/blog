package utils

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/qianxia/blog/global"
)

func Redis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", global.CONFIG.Redis.Host, global.CONFIG.Redis.Port),
		Password: global.CONFIG.Redis.Password,
		DB:       global.CONFIG.Redis.DB,
	})
}
