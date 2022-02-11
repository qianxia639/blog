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
		r.POST("/save", middleware.Auth(), blogHandler.Save)
		r.GET("/list", middleware.Auth(), blogHandler.List)
		r.POST("/pageList", blogHandler.pageList)
		r.DELETE("/:id", blogHandler.Delete)
	}

	return e
}
