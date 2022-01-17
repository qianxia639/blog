package routers

import (
	"github.com/gin-gonic/gin"

	blog "github.com/qianxia/blog/app/blog"
	tag "github.com/qianxia/blog/app/tag"
	types "github.com/qianxia/blog/app/type"
	user "github.com/qianxia/blog/app/user"
	"github.com/qianxia/blog/middleware"
)

type Option func(*gin.Engine) *gin.Engine

var options = []Option{}

// 注册app的路由配置
func include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	// 加载app的配置路由
	include(user.Routers, types.Routers, blog.Routers, tag.Routers)

	r := gin.Default()
	r.Use(middleware.CORS())
	for _, opt := range options {
		opt(r)
	}
	return r
}
