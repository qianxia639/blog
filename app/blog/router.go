package app

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) *gin.Engine {
	/*blogHandler := NewBlogHandler()
	// e.Use(middleware.AuthorizationMiddlleware())
	r := e.Group("/blog")
	{
		r.POST("/save/:typeId", middleware.Auth(), blogHandler.Save)
		r.GET("/list", middleware.Auth(), blogHandler.List)
		//r.POST("/pageList", blogHandler.PageList)
		r.DELETE("/:id", blogHandler.Delete)
	}
	*/
	return e
}
