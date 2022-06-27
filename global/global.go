package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/qianxia/blog/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	LOG    *zap.SugaredLogger
	CONFIG *config.Config
	REDIS  *redis.Client
)
