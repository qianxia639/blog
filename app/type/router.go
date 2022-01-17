package app

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) *gin.Engine {
	typeHandler := NewTypeHandler()
	r := e.Group("/type")
	{
		r.GET("/list", typeHandler.List)
	}

	return e
}
