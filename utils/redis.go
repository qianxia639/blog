package utils

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/qianxia/blog/global"
)

func Redis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", global.QX_CONFIG.Redis.Host, global.QX_CONFIG.Redis.Port),
		Password: global.QX_CONFIG.Redis.Password,
		DB:       global.QX_CONFIG.Redis.DB,
	})
}
