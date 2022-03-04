package utils

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/qianxia/blog/config"
)

func InitRedis(y *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", y.Redis.Host, y.Redis.Port),
		DB:   0,
	})
}
