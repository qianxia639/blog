package app

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) *gin.Engine {
	tagHandler := NewTagHandler()

	r := e.Group("/tag")
	{
		// 标签列表(不分页)
		r.GET("/list", tagHandler.tagList)
	}

	return e
}
