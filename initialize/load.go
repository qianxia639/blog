package initialize

import (
	"time"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/utils"
	"go.uber.org/zap"
)

func Load() {
	// 读取yaml配置文件
	// global.QX_CONFIG = utils.ParseConfig()
	global.QX_CONFIG = utils.DeCode()

	// 加载log日志
	global.QX_LOG = utils.InitLogger("./log/info."+time.Now().Format("2006-01-02")+".log", global.QX_WARN_PATH, zap.InfoLevel)

	// 加载MySQL配置信息
	global.QX_DB = utils.InitDb(global.QX_CONFIG)
	if global.QX_DB != nil {
		// 初始化表
		RegisterTables(global.QX_DB)
	}
}
