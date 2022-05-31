package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/handler/system"
	"github.com/qianxia/blog/middleware"
)

func SystemRouters(e *gin.Engine) *gin.Engine {

	// ========== system router group ==========
	captchaRouterApi := system.SystemRouterGroups.CaptchaHandler
	emailRouterApi := system.SystemRouterGroups.EmailHandler
	sysg := e.Group("/system")
	{
		sysg.POST("/captcha", captchaRouterApi.Captcha)     // 生成验证码
		sysg.GET("/email", emailRouterApi.SendMail)         // 发送邮箱验证码
		sysg.POST("/verifyMail", emailRouterApi.VerifyMail) // 校验邮箱验证码
	}

	// ========== user router group ==========

	userRouterWithoutRecord := e.Group("/user")
	userRouter := e.Group("/user").Use(middleware.Auth())
	userRouterApi := system.SystemRouterGroups.UserHandler
	{
		userRouterWithoutRecord.POST("/register", userRouterApi.Register)     // 注册
		userRouterWithoutRecord.POST("/login", userRouterApi.Login)           // 登录
		userRouterWithoutRecord.POST("/emailLogin", userRouterApi.EmailLogin) // 邮箱登录
	}
	{
		userRouter.GET("/info", userRouterApi.UserInfo)       // 用户信息
		userRouter.PUT("/name", userRouterApi.UpdateUsername) // 修改名称
		userRouter.PUT("/pwd", userRouterApi.UpdatePwd)       // 修改密码
		userRouter.PUT("/avatar", userRouterApi.UpdateAvatar) // 修改头像
		userRouter.PUT("/email", userRouterApi.UpdateEmail)   // 修改邮箱
	}

	//  ========== search router group ==========
	searchRouterWithoutRecord := e.Group("/search")
	searchRouter := e.Group("/search").Use(middleware.Auth())
	searchRouterApi := system.SystemRouterGroups.SearchHandler
	{
		searchRouterWithoutRecord.GET("/blog", searchRouterApi.SearchBlog) // 搜索所有博客
	}
	{
		searchRouter.GET("/priblog", searchRouterApi.SearchPriBlog) // 搜索个人博客列表
	}

	// ========== upload router group ==========
	// uploadRouterWithoutRecord := e.Group("/upload")
	uploadRouter := e.Group("/upload").Use(middleware.Auth())
	uploadRouterApi := system.SystemRouterGroups.UploadHandler
	{
		uploadRouter.POST("/mdFile", uploadRouterApi.UploadMdFile) // markdown文件上传
	}

	// ========== comment router group ==========

	return e
}
