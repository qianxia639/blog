package example

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/service/example"
)

type TypeHandler struct {
	typeService example.TypeService
}

// 按amount降序排列
func (th TypeHandler) ListOrder(ctx *gin.Context) {
	types, err := th.typeService.ListOrderByAmountDesc()
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, command.FindFailed)
		return
	}
	command.Success(ctx, command.FindSuccess, gin.H{"type": types})
}

// 不排序只显示列表
func (th TypeHandler) List(ctx *gin.Context) {
	types, err := th.typeService.List()
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, command.FindFailed)
		return
	}
	command.Success(ctx, command.FindSuccess, gin.H{"type": types})
}

//	点击分类进行查询并分页
func (th TypeHandler) TypeList(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Query("id"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("paginate", "6"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	typeList, err := th.typeService.TypeList(id, pageSize, pageNo)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, command.OperationFailed)
		return
	}
	command.Success(ctx, command.OperationSuccess, typeList)
}
