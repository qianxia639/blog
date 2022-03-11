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
)

func main() {
	// 初始化路由
	router := routers.Init()

	// 加载配置信息
	initialize.Load()

	db, _ := global.QX_DB.DB()
	defer db.Close()
	defer global.QX_LOG.Sync()
	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", global.QX_YAML_CONFIG.Server.Host, global.QX_YAML_CONFIG.Server.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

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
