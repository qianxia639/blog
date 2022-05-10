package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/handler/system"
	"github.com/qianxia/blog/middleware"
)

func SystemRouters(e *gin.Engine) *gin.Engine {

	// ========== system router group ==========
	sysg := e.Group("/system")
	{
		// 注册
		sysg.POST("/register", system.GetInstance().Register)
		// 登录
		sysg.POST("/login", system.GetInstance().Login)
		// 生成验证码
		sysg.POST("/captcha", system.GetInstance().Captcha)

		sysg = sysg.Group("/")
		sysg.Use(middleware.Auth())
		{
			// 发送邮箱验证码
			sysg.GET("/email", system.GetInstance().SendMail)
			// 校验邮箱验证码
			sysg.POST("/verifyMail", system.GetInstance().VerifyMail)
		}
	}

	// ========== user router group ==========
	ug := e.Group("/user")
	ug.Use(middleware.Auth())
	{
		// 用户信息
		ug.GET("/info", system.GetInstance().Info)
		// 修改名称
		ug.PUT("/name", system.GetInstance().UpdateUsername)
		// 修改密码
		ug.PUT("/pwd", system.GetInstance().UpdatePwd)
		// 修改头像
		ug.PUT("/avatar", system.GetInstance().UpdateAvatar)
		// 修改邮箱
		ug.PUT("/email", system.GetInstance().UpdateEmail)
	}

	//  ========== search router group ==========
	sg := e.Group("/search")
	{
		// 搜索所有博客
		sg.GET("/blog", system.GetInstance().SearchBlog)
		// 搜索个人博客列表
		sg.GET("/priblog", middleware.Auth(), system.GetInstance().SearchPriBlog)
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
