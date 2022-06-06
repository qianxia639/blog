package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/handler/example"
	"github.com/qianxia/blog/middleware"
)

func ExampleRouters(e *gin.Engine) *gin.Engine {
	//  ========== blog router group ==========
	blogGroup := e.Group("/blog")
	blogRouter := e.Group("/blog").Use(middleware.Authorization())
	blogRouterApi := example.ExampleRouterGroups.BlogHandler
	{
		blogGroup.GET("/pageList", blogRouterApi.BlogPageList) // 博客分页列表
		blogGroup.GET("/latestList", blogRouterApi.LatestList) // 最新推荐(按更新时间降序排列)
		blogGroup.GET("/:id", blogRouterApi.GetBlogInfo)       // 获取博客信息
	}
	{

		blogRouter.POST("/save", blogRouterApi.CreateBlog)  // 新增博客
		blogRouter.PUT("/update", blogRouterApi.UpdateBlog) // 编辑博客
		blogRouter.GET("/list", blogRouterApi.BlogList)     //个人博客展示
		blogRouter.DELETE("/:id", blogRouterApi.DeleteBlog) // 根据id删除博客
	}

	//  ========== type router group ==========
	typeGroup := e.Group("/type")
	typeRouter := e.Group("/type").Use(middleware.Authorization())
	typeRouterApi := example.ExampleRouterGroups.TypeHandler
	{
		typeGroup.GET("/listOrder", typeRouterApi.ListOrder) // 分类列表(按amount降序排列)
		typeGroup.GET("/list", typeRouterApi.TypeList)       // 分类列表(不排序)
		typeGroup.GET("/page", typeRouterApi.TypePageList)   // 按分类进行博客的展示并分页
	}
	{
		typeRouter.POST("/save", typeRouterApi.CreateType)
	}

	//  ========== tag router group ==========
	tagRouterApi := example.ExampleRouterGroups.TagHandler
	tagRouterWithoutRecord := e.Group("/tag")
	{
		tagRouterWithoutRecord.GET("/list", tagRouterApi.TagList) // 标签列表(不分页)
	}

	// ========== archive router group ==========
	archiveRouterApi := example.ExampleRouterGroups.ArchiveHandler
	archiveRouterWithoutRecord := e.Group("/archive")
	{
		archiveRouterWithoutRecord.GET("/list", archiveRouterApi.ArchiveList) // 归档列表
	}
	return e
}
