package global

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/qianxia/blog/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	QX_DB     *gorm.DB
	QX_LOG    *zap.SugaredLogger
	QX_CONFIG *config.Config
	QX_ES     *elasticsearch.Client
	QX_REDIS  *redis.Client
)
