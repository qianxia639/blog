package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/handler/example"
	"github.com/qianxia/blog/middleware"
)

func ExampleRouters(e *gin.Engine) *gin.Engine {
	//  ========== blog router group ==========
	bg := e.Group("/blog")
	{
		// 博客分页列表
		bg.GET("/pageList", example.GetInstance().BlogPageList)
		// 最新推荐(按更新时间降序排列)
		bg.GET("/latestList", example.GetInstance().LatestList)
		// 获取博客信息
		bg.GET("/:id", example.GetInstance().GetBlog)

		bg = bg.Group("/")
		bg.Use(middleware.Auth())
		{
			// 新增博客
			bg.POST("/save", example.GetInstance().CreateBlog)
			// 获取要编辑的博客的信息
			bg.GET("/update/:id", example.GetInstance().GetUpdateBlog)
			// 编辑博客
			bg.PUT("/update", example.GetInstance().UpdateBlog)
			//个人博客展示
			bg.GET("/list", example.GetInstance().BlogList)
			// 根据id删除博客
			bg.DELETE("/:id", example.GetInstance().DeleteBlog)
		}
	}

	//  ========== type router group ==========
	tg := e.Group("/type")
	{
		// 分类列表(按amount降序排列)
		tg.GET("/listOrder", example.GetInstance().ListOrder)
		// 分类列表(不排序)
		tg.GET("/list", example.GetInstance().List)
		// 点击分类进行博客的展示并分页
		tg.GET("/page", example.GetInstance().TypeList)

		tg = tg.Group("")
		tg.Use(middleware.Auth())
		{
			tg.POST("/save", example.GetInstance().CreateType)
		}

	}

	//  ========== tag router group ==========
	tgg := e.Group("/tag")
	{
		// 标签列表(不分页)
		tgg.GET("/list", example.GetInstance().TagList)
	}

	// ========== archive router group ==========
	ag := e.Group("/archive")
	{
		// 归档列表
		ag.GET("/list", example.GetInstance().ArchiveList)
	}
	return e
}
