package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/qianxia/blog/docs"
	"github.com/qianxia/blog/handler"
	"github.com/qianxia/blog/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option func(*gin.Engine) *gin.Engine

var options = []Option{}

// 注册路由配置
func include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Router() *gin.Engine {
	// 加载路由配置
	include(handler.ExampleRouters, handler.SystemRouters)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	gin.SetMode(gin.ReleaseMode)
	r.Use(middleware.CORS())
	for _, opt := range options {
		opt(r)
	}
	return r
}
