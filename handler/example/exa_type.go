package example

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/service/example"
)

type TypeHandler struct {
	typeService example.TypeService
}

// 按amount降序排列
func (th *TypeHandler) ListOrder(ctx *gin.Context) {
	types, err := th.typeService.ListOrderByAmountDesc()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}

// 不排序只显示列表
func (th *TypeHandler) List(ctx *gin.Context) {
	types, err := th.typeService.List()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}

//
func (th *TypeHandler) TypeList(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	typeList, err := th.typeService.TypeList(id)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", gin.H{"typeList": typeList})
}
