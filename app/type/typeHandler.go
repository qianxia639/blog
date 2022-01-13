package types

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
)

type ITypeHandler interface {
	List(ctx *gin.Context)
}

type TypeHandler struct {
	Service TypeService
}

func NewTypeHandler() ITypeHandler {
	typeService := NewTypeService()

	return TypeHandler{Service: typeService}
}

func (t TypeHandler) List(ctx *gin.Context) {
	types, err := t.Service.List()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}
