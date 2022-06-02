package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/initialize"
	"github.com/qianxia/blog/routers"
	"github.com/qianxia/blog/utils"
)

// @title           Blog API Swagger
// @version         1.0
// @description     This is blog server.
// @securitydefinitions.apikey  X-Token
// @in  heaader
// @name X-Token
// @BasePath  /
func main() {
	utils.Viper()               // 初始化配置文件信息
	global.QX_LOG = utils.Zap() // 初始化zap日志
	// global.QX_ES = utils.ElasticSearch()                                              // 初始化elasticsearch
	// if err := system.SystemGroups.ElasticSearchService.IndicesMapping(); err != nil { // 初始化索引
	// 	global.QX_LOG.Fatal(err)
	// 	return
	// }

	global.QX_REDIS = utils.Redis() // 初始化redis

	if err := global.QX_REDIS.Ping(context.Background()).Err(); err != nil {
		global.QX_LOG.Fatal(err)
	}

	global.QX_DB = utils.Mysql(global.QX_CONFIG) // 初始化mysql
	if global.QX_DB != nil {
		initialize.RegisterTables(global.QX_DB) // 初始化表
		db, _ := global.QX_DB.DB()
		defer db.Close() // 关闭连接
	}
	defer global.QX_LOG.Sync()

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", global.QX_CONFIG.Http.Host, global.QX_CONFIG.Http.Port),
		Handler:        routers.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		global.QX_LOG.Error(server.ListenAndServe().Error())
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	global.QX_LOG.Error(server.Shutdown(ctx))
}
