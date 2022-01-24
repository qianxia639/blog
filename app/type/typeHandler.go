package app

import (
	"github.com/gin-gonic/gin"
)

type ITypeHandler interface {
	List(ctx *gin.Context)
	typeList(ctx *gin.Context)
}

type TypeHandler struct {
	Service TypeService
}

/*
func NewTypeHandler() ITypeHandler {
	var typeService TypeService

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

func (t TypeHandler) typeList(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	typeList, err := t.Service.typeList(id)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", gin.H{"typeList": typeList})
}
*/
