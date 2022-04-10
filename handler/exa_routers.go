package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/handler/example"
	"github.com/qianxia/blog/middleware"
)

func ExampleRouters(e *gin.Engine) *gin.Engine {
	//  ========== blog router group ==========
	blogGroup := e.Group("/blog")
	{
		// 博客分页列表
		blogGroup.GET("/pageList", example.GetInstance().BlogPageList)
		// 最新推荐(按更新时间降序排列)
		blogGroup.GET("/latestList", example.GetInstance().LatestList)
		// 获取博客信息
		blogGroup.GET("/:id", example.GetInstance().GetBlog)

		blogGroup = blogGroup.Group("/")
		blogGroup.Use(middleware.Auth())
		{
			// 新增博客
			blogGroup.POST("/save", example.GetInstance().CreateBlog)
			// 获取要编辑的博客的信息
			blogGroup.GET("/update/:id", example.GetInstance().GetUpdateBlog)
			// 编辑博客
			blogGroup.PUT("/update", example.GetInstance().UpdateBlog)
			//个人博客展示
			blogGroup.GET("/list", example.GetInstance().BlogList)
			// 根据id删除博客
			blogGroup.DELETE("/:id", example.GetInstance().DeleteBlog)
		}
	}

	//  ========== type router group ==========
	typeGroup := e.Group("/type")
	{
		// 分类列表(按amount降序排列)
		typeGroup.GET("/listOrder", example.GetInstance().ListOrder)
		// 分类列表(不排序)
		typeGroup.GET("/list", example.GetInstance().List)
		// 点击分类进行博客的展示并分页
		typeGroup.GET("/page", example.GetInstance().TypeList)

		typeGroup = typeGroup.Group("")
		typeGroup.Use(middleware.Auth())
		{
			typeGroup.POST("/save", example.GetInstance().CreateType)
		}

	}

	//  ========== tag router group ==========
	tagGroup := e.Group("/tag")
	{
		// 标签列表(不分页)
		tagGroup.GET("/list", example.GetInstance().TagList)
	}

	// ========== archive router group ==========
	archiveGroup := e.Group("/archive")
	{
		// 归档列表
		archiveGroup.GET("/list", example.GetInstance().ArchiveList)
	}
	return e
}
