package blog

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) *gin.Engine {
	blogHandler := NewBlogHandler()
	r := e.Group("/blog")
	{
		r.POST("/save/:userId/:typeId", blogHandler.Save)
	}

	return e
}
