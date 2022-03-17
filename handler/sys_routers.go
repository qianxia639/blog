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

		userGroup = userGroup.Group("/", middleware.Auth())
		{
			// 用户信息
			userGroup.GET("/info", system.GetInstance().Info)
			// 修改名称
			userGroup.PUT("/updateName", system.GetInstance().UpdateUsername)
		}

		// 修改密码
		// userGroup.PUT("/updatePwd", system.GetInstance().UpdatePassword)
	}

	//  ========== search router group ==========
	searchGroup := e.Group("/search")
	{
		// 搜索所有博客
		searchGroup.GET("/blog", system.GetInstance().SearchBlog)
		// 搜索个人博客列表
		searchGroup.GET("/priblog", system.GetInstance().SearchPriBlog)
	}

	// ========== upload router group ==========
	fileGroup := e.Group("/upload")
	{
		// markdown文件上传
		fileGroup.POST("/mdFile", system.GetInstance().UploadMdFile)
	}

	return e
}
