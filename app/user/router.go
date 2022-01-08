package user

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/middleware"
)

func Routers(e *gin.Engine) *gin.Engine {
	// userHandler := NewUserHandler()
	var userHandler UserHandler
	r := e.Group("/user")
	{
		// 注册
		r.POST("/register", userHandler.registerHandler)
		// 登录
		r.POST("/login", userHandler.loginHandler)
		// 用户信息
		r.GET("/info", middleware.AuthorizationMiddlleware(), userHandler.infoHandler)
	}
	return e
}
