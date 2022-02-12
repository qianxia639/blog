package app

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/middleware"
)

func Routers(e *gin.Engine) *gin.Engine {
	blogHandler := NewBlogHandler()
	r := e.Group("/blog")
	{
		// 新增博客
		r.POST("/save", middleware.Auth(), blogHandler.createBlog)
		//个人博客展示
		r.GET("/list", middleware.Auth(), blogHandler.blogList)
		// 博客分页列表
		r.POST("/pageList", blogHandler.blogPageList)
		// 根据id删除博客
		r.DELETE("/:id", blogHandler.deleteBlog)
		// 最新推荐(按更新时间降序排列)
		r.GET("/latestList", blogHandler.latestList)
	}

	return e
}
