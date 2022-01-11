package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
)

type ITypeHandler interface {
	List(ctx *gin.Context)
}

type TypeHandler struct {
	TypeService
}

func NewTypeHandler() ITypeHandler {
	var typeService TypeService

	return TypeHandler{TypeService: typeService}
}

func (t TypeHandler) List(ctx *gin.Context) {
	types, err := t.TypeService.List()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}
