package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/handler/system"
	"github.com/qianxia/blog/middleware"
)

func SystemRouters(e *gin.Engine) *gin.Engine {
	// ========== user router group ==========
	ug := e.Group("/user")
	{
		// 注册
		ug.POST("/register", system.GetInstance().Register)
		// 登录
		ug.POST("/login", system.GetInstance().Login)
		// 生成验证码
		ug.POST("/captcha", system.GetInstance().Captcha)

		ug = ug.Group("/")
		ug.Use(middleware.Auth())
		{
			// 用户信息
			ug.GET("/info", system.GetInstance().Info)
			// 修改名称
			ug.PUT("/updateName", system.GetInstance().UpdateUsername)
			// 修改密码
			ug.PUT("/updatePwd", system.GetInstance().UpdatePwd)
			// 修改头像
			ug.PUT("/updateAvatar", system.GetInstance().UpdateAvatar)
		}
	}

	//  ========== search router group ==========
	sg := e.Group("/search")
	{
		// 搜索所有博客
		sg.GET("/blog", system.GetInstance().SearchBlog)
		// 搜索个人博客列表
		sg.GET("/priblog", system.GetInstance().SearchPriBlog, middleware.Auth())
	}

	// ========== upload router group ==========
	fg := e.Group("/upload")
	fg.Use(middleware.Auth())
	{
		// markdown文件上传
		fg.POST("/mdFile", system.GetInstance().UploadMdFile)
	}

	// ========== comment router group ==========
	cg := e.Group("/comment")
	{
		// 获取comment列表
		cg.GET("/list", system.GetInstance().CommentList)
		// 发布评论
		cg.POST("/save", system.GetInstance().Save)
	}

	// ========== leave router group ==========
	lg := e.Group("/leave")
	{
		// 获取所有留言记录
		lg.GET("/all", system.GetInstance().All)
		// 添加留言
		lg.POST("/insert", system.GetInstance().Insert)
		// 删除留言
		lg.DELETE("/delete", system.GetInstance().Delete)
	}

	return e
}
