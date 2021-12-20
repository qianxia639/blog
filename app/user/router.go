package user

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/middleware"
)

func Routers(e *gin.Engine) *gin.Engine {
	r := e.Group("/user")
	{
		// 注册
		r.POST("/register", registerHandler)
		// 登录
		r.POST("/login", loginHandler)
		// 用户信息
		r.GET("/info", middleware.AuthorizationMiddlleware(), infoHandler)
	}
	return e
}
