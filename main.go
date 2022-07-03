package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
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
	utils.Viper() // 初始化配置文件信息

	global.LOG = utils.Zap() // 初始化zap日志
	if global.LOG != nil {
		defer global.LOG.Sync()
	}

	global.REDIS = utils.Redis() // 初始化redis

	if err := global.REDIS.Ping(context.Background()).Err(); err != nil {
		global.LOG.Fatal(err)
	}

	global.DB = utils.Mysql(global.CONFIG) // 初始化mysql
	if global.DB != nil {
		err := global.DB.AutoMigrate(&model.User{}, &model.Type{}, &model.Comment{}, &model.Blog{}, &model.Tag{})
		if err != nil {
			global.LOG.Fatalf("Error AutoMigrate Table: %v\n", err)
		}
		db, _ := global.DB.DB()
		defer db.Close()
	}

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", global.CONFIG.Http.Host, global.CONFIG.Http.Port),
		Handler:        routers.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOG.Fatalf("Error Listen Server: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		global.LOG.Fatalf("Error Server Shutdown: %v\n", err)
	}

}
