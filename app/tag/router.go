package tag

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) *gin.Engine {
	tagHandler := NewTagHandler()

	r := e.Group("/tag")
	{
		r.GET("/list", tagHandler.List)
	}

	return e
}
