package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/initialize"
	"github.com/qianxia/blog/routers"
	"github.com/qianxia/blog/server"
)

var (
	confPath string
	fileType string
)

func init() {
	if runtime.GOOS == "windows" {
		flag.StringVar(&confPath, "conf-path", "./config/application.toml", "配置文件路径")
	} else if runtime.GOOS == "linux" {
		flag.StringVar(&confPath, "conf-path", "/opt/conf/application.toml", "配置文件路径")
	}
	flag.StringVar(&fileType, "type", "toml", "配置文件类型(支持toml和yaml)")
}

func main() {
	flag.Parse()
	// 初始化路由
	router := routers.InitRouter()
	// 加载配置信息
	initialize.Load(confPath, fileType)

	db, _ := global.QX_DB.DB()
	defer db.Close()
	defer global.QX_LOG.Sync()

	srv := server.Server(router)

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.QX_LOG.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号关闭服务器(设置5秒超时时间)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.QX_LOG.Fatal("Server Shutdown: ", err)
	}
}
