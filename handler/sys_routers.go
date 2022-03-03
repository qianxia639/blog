package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/handler/system"
	"github.com/qianxia/blog/middleware"
)

func SystemRouters(e *gin.Engine) *gin.Engine {
	// ========== user router group ==========
	userGroup := e.Group("/user")
	{
		// 注册
		userGroup.POST("/register", system.GetInstance().Register)
		// 登录
		userGroup.POST("/login", system.GetInstance().Login)
		ug := userGroup.Group("/", middleware.Auth())
		{
			// 用户信息
			ug.GET("/info", system.GetInstance().Info)
			// 修改名称
			ug.PUT("/updateName", system.GetInstance().UpdateUsername)
		}

		// 修改密码
		// userGroup.PUT("/updatePwd", system.GetInstance().UpdatePassword)
	}
	//  ========== search router group ==========

	return e
}
