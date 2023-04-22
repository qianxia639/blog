package cache

import (
	"Blog/core/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis(conf config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Address,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
}
