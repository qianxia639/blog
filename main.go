package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/routers"
)

func main() {
	// 初始化路由
	router := routers.Init()
	// 加载配置信息
	db := routers.Load()
	defer global.RY_LOG.Sync()
	defer db.Close()

	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", global.RY_YAML_CONFIG.Server.Host, global.RY_YAML_CONFIG.Server.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.RY_LOG.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号关闭服务器(设置5秒超时时间)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.RY_LOG.Fatal("Server Shutdown: ", err)
	}
}
