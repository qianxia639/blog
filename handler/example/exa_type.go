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

// @Summary      分类展示(amount降序)
// @Tags         Example/Type
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  []model.Type
// @Router       /type/listOrder [get]
func (th *TypeHandler) ListOrder(ctx *gin.Context) {
	types, err := th.typeService.ListOrderByAmountDesc()
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}

// @Summary      分类展示(不排序)
// @Tags         Example/Type
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  []model.Type
// @Router       /type/list [get]
func (th *TypeHandler) List(ctx *gin.Context) {
	types, err := th.typeService.List()
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"type": types})
}

// @Summary      分类查询分页
// @Tags         Example/Type
// @Accept       json
// @Produce      json
// @Param        id	  		query     int	true	"分类id"
// @Param        pageSize	query     int	false	"每页显示的"
// @Param        pageNo	  	query     int	false	"页码"
// @Success 	 200  {object}  response.PageList	{data=response.PageList}
// @Router       /type/page [get]
func (th *TypeHandler) TypeList(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Query("id"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "6"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "1"))

	typeList, err := th.typeService.TypeList(id, pageSize, pageNo)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", typeList)
}

// @Summary      新增分类
// @Tags         Example/Type
// @Accept       json
// @Produce      json
// @Param        blog body map[string]string  true  "Create Type"
// @Success 	 200  {object}  string
// @Security	 X-Token
// @Router       /type/save [post]
func (th *TypeHandler) CreateType(ctx *gin.Context) {
	var m map[string]string
	ctx.ShouldBindJSON(&m)
	if err := th.typeService.CreateType(m["typeName"]); err != nil {
		global.QX_LOG.Errorf("insert type err: %v", err)
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "添加成功", nil)

}
