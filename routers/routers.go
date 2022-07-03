package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/handler"
	"github.com/qianxia/blog/middleware"
)

type Option func(*gin.Engine) *gin.Engine

// 注册路由配置
func include(opts ...Option) (options []Option) {
	options = append(options, opts...)
	return
}

// 初始化
func Router() *gin.Engine {
	// 加载路由配置
	options := include(handler.ExampleRouters, handler.SystemRouters)

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	// r.Use(middleware.Recovery())
	r.Use(middleware.CORS())
	// r.Use(middleware.Authorization(system.CasbinServices.Casbin()))
	for _, opt := range options {
		opt(r)
	}
	return r
}
