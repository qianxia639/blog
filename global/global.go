package global

import (
	"time"

	"github.com/qianxia/blog/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	QX_DB          *gorm.DB
	QX_LOG         *zap.SugaredLogger
	QX_YAML_CONFIG *config.Config
)

// ============== log path ==============
var (
	QX_INFO_PATH = "./log/info." + time.Now().Format("2006-01-02") + ".log"
	QX_WARN_PATH = "./log/error." + time.Now().Format("2006-01-02") + ".log"
)
