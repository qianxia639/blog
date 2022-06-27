package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/qianxia/blog/docs"
	"github.com/qianxia/blog/handler/system"
	"github.com/qianxia/blog/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SystemRouters(e *gin.Engine) *gin.Engine {

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ========== system router group ==========
	captchaRouterApi := system.SystemRouterGroups.CaptchaHandler
	sysearchGroup := e.Group("/system")
	{
		sysearchGroup.POST("/captcha", captchaRouterApi.Captcha) // 生成验证码
	}

	// ========== user router group ==========
	userGroup := e.Group("/user")
	userRouter := e.Group("/user").Use(middleware.Authorization())
	userRouterApi := system.SystemRouterGroups.UserHandler
	{
		userGroup.POST("/register", userRouterApi.Register)   // 注册
		userGroup.POST("/login", userRouterApi.Login)         // 登录
		userGroup.GET("/logout", userRouterApi.Logout)        // 登出
		userGroup.POST("/forgetPwd", userRouterApi.ForgetPwd) // 找回密码
	}
	{
		userRouter.GET("/info", userRouterApi.UserInfo)       // 用户信息
		userRouter.PUT("/name", userRouterApi.UpdateNickname) // 修改名称
		userRouter.PUT("/pwd", userRouterApi.UpdatePwd)       // 修改密码
		userRouter.PUT("/avatar", userRouterApi.UpdateAvatar) // 修改头像
	}

	//  ========== search router group ==========
	searchGroup := e.Group("/search")
	searchRouterApi := system.SystemRouterGroups.SearchHandler
	{
		searchGroup.GET("/blog", searchRouterApi.SearchBlog) // 搜索所有博客
	}

	// ========== upload router group ==========
	uploadRouter := e.Group("/upload").Use(middleware.Authorization())
	uploadRouterApi := system.SystemRouterGroups.UploadHandler
	{
		uploadRouter.POST("/mdFile", uploadRouterApi.UploadMdFile) // markdown文件上传
	}

	// ========== comment router group ==========
	commentGroup := e.Group("/comment")
	commentRouterApi := system.SystemRouterGroups.CommentHandler
	{
		commentGroup.POST("/save", commentRouterApi.SaveComment)             // 添加评论
		commentGroup.DELETE("/parent", commentRouterApi.DeleteParentComment) // 删除父级评论
		commentGroup.DELETE("/child", commentRouterApi.DeleteChildComment)   // 删除子级平论
		commentGroup.GET("/list", commentRouterApi.CommentList)              // 评论列表
	}

	return e
}
