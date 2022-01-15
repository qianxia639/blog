package blog

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
	}

	return e
}
