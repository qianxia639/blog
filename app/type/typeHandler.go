package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
)

type ITypeHandler interface {
	listOrder(ctx *gin.Context)
	list(ctx *gin.Context)
	typeList(ctx *gin.Context)
}

type TypeHandler struct {
	Service TypeService
}

func NewTypeHandler() ITypeHandler {
	var typeService TypeService

	return TypeHandler{Service: typeService}
}

// 按amount降序排列
func (t TypeHandler) listOrder(ctx *gin.Context) {
	types, err := t.Service.ListOrderByAmountDesc()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}

// 不排序只显示列表
func (t TypeHandler) list(ctx *gin.Context) {
	types, err := t.Service.List()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}

//
func (t TypeHandler) typeList(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	typeList, err := t.Service.typeList(id)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", gin.H{"typeList": typeList})
}
