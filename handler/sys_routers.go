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
	sysearchGroup := e.Group("/system")
	{
		sysearchGroup.POST("/captcha", captchaRouterApi.Captcha)     // 生成验证码
		sysearchGroup.GET("/email", emailRouterApi.SendMail)         // 发送邮箱验证码
		sysearchGroup.POST("/verifyMail", emailRouterApi.VerifyMail) // 校验邮箱验证码
	}

	// ========== user router group ==========
	userGroup := e.Group("/user")
	userRouter := e.Group("/user").Use(middleware.Authorization())
	userRouterApi := system.SystemRouterGroups.UserHandler
	{
		userGroup.POST("/register", userRouterApi.Register)     // 注册
		userGroup.POST("/login", userRouterApi.Login)           // 登录
		userGroup.POST("/emailLogin", userRouterApi.EmailLogin) // 邮箱登录
		userGroup.GET("/logout", userRouterApi.Logout)          // 登出
	}
	{
		userRouter.GET("/info", userRouterApi.UserInfo)       // 用户信息
		userRouter.PUT("/name", userRouterApi.UpdateUsername) // 修改名称
		userRouter.PUT("/pwd", userRouterApi.UpdatePwd)       // 修改密码
		userRouter.PUT("/avatar", userRouterApi.UpdateAvatar) // 修改头像
		userRouter.PUT("/email", userRouterApi.UpdateEmail)   // 修改邮箱
	}

	//  ========== search router group ==========
	searchGroup := e.Group("/search")
	searchRouter := e.Group("/search").Use(middleware.Authorization())
	searchRouterApi := system.SystemRouterGroups.SearchHandler
	{
		searchGroup.GET("/blog", searchRouterApi.SearchBlog) // 搜索所有博客
	}
	{
		searchRouter.GET("/priblog", searchRouterApi.SearchPriBlog) // 搜索个人博客列表
	}

	// ========== upload router group ==========
	uploadRouter := e.Group("/upload").Use(middleware.Authorization())
	uploadRouterApi := system.SystemRouterGroups.UploadHandler
	{
		uploadRouter.POST("/mdFile", uploadRouterApi.UploadMdFile) // markdown文件上传
	}

	// ========== comment router group ==========

	return e
}
