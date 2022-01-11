package user

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/middleware"
)

func Routers(e *gin.Engine) *gin.Engine {
	userHandler := NewUserHandler()
	// var userHandler UserHandler
	r := e.Group("/user")
	{
		// 注册
		r.POST("/register", userHandler.Register)
		// 登录
		r.POST("/login", userHandler.Login)
		// 用户信息
		r.GET("/info", middleware.AuthorizationMiddlleware(), userHandler.Info)
	}
	return e
}
