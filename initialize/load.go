package initialize

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/service/system"
	"github.com/qianxia/blog/utils"
)

func Load(path string) {
	// 读取配置文件
	utils.Viper(path)
	// 加载log日志
	global.QX_LOG = utils.Zap()
	global.QX_ES = utils.ElasticSearch()

	system.ElasticSearch.IndicesAndMapping()

	// 加载MySQL配置信息
	global.QX_DB = utils.InitDb(global.QX_CONFIG)
	if global.QX_DB != nil {
		// 初始化表
		RegisterTables(global.QX_DB)
	}
}
