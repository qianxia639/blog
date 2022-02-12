package app

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) *gin.Engine {
	typeHandler := NewTypeHandler()
	r := e.Group("/type")
	{
		// 分类列表(按amount降序排列)
		r.GET("/listOrder", typeHandler.listOrder)
		// 分类列表(不排序)
		r.GET("/list", typeHandler.list)
		r.GET("/:id", typeHandler.typeList)
	}

	return e
}
