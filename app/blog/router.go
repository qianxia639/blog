package app

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/middleware"
)

func Routers(e *gin.Engine) *gin.Engine {
	blogHandler := NewBlogHandler()
	// e.Use(middleware.AuthorizationMiddlleware())
	r := e.Group("/blog")
	{
		r.POST("/save/:typeId", middleware.AuthorizationMiddlleware(), blogHandler.Save)
		r.GET("/list", middleware.AuthorizationMiddlleware(), blogHandler.List)
		r.GET("/pageList/:pageNo/:pageSize", blogHandler.PageList)
		r.DELETE("/:id", blogHandler.Delete)
	}

	return e
}
