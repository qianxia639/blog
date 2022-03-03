package global

import (
	"time"

	"github.com/qianxia/blog/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	RY_DB          *gorm.DB
	RY_LOG         *zap.SugaredLogger
	RY_YAML_CONFIG *config.Config
)

// ============== log path ==============
var (
	RY_INFO_PATH = "./log/info." + time.Now().Format("2006-01-02") + ".log"
	RY_WARN_PATH = "./log/error." + time.Now().Format("2006-01-02") + ".log"
)
