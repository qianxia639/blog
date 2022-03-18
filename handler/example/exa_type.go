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
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}

// 不排序只显示列表
func (th TypeHandler) List(ctx *gin.Context) {
	types, err := th.typeService.List()
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}

//	点击分类进行查询并分页
func (th TypeHandler) TypeList(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Query("id"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "6"))
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))

	typeList, err := th.typeService.TypeList(id, pageSize, pageNum)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", typeList)
}
