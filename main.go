package main

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/initialize"
	"github.com/qianxia/blog/server"
	"github.com/qianxia/blog/service/system"
	"github.com/qianxia/blog/utils"
)

func main() {
	utils.Viper()                                                 // 初始化配置文件信息
	global.QX_LOG = utils.Zap()                                   // 初始化zap日志
	global.QX_ES = utils.ElasticSearch()                          // 初始化elasticsearch
	if err := system.ElasticSearch.IndicesMapping(); err != nil { // 初始化索引
		global.QX_LOG.Fatal(err)
		return
	}

	global.QX_REDIS = utils.Redis() // 初始化redis

	global.QX_DB = utils.Mysql(global.QX_CONFIG) // 初始化mysql
	if global.QX_DB != nil {
		initialize.RegisterTables(global.QX_DB) // 初始化表
		db, _ := global.QX_DB.DB()
		defer db.Close() // 关闭连接
	}
	defer global.QX_LOG.Sync()

	server.Run()
}
