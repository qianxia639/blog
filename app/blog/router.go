package app

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/middleware"
)

func Routers(e *gin.Engine) *gin.Engine {
	blogHandler := NewBlogHandler()
	e.Use(middleware.AuthorizationMiddlleware())
	r := e.Group("/blog")
	{
		r.POST("/save/:typeId", blogHandler.Save)
		r.GET("/list", blogHandler.List)
		r.GET("/show", blogHandler.Show)
		r.GET("/latest", blogHandler.Latest)
		r.DELETE("/:id", blogHandler.Delete)
	}

	return e
}
