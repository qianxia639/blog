package example

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
)

type TypeHandler struct{}

// @Summary      分类展示(amount降序)
// @Tags         Example/Type
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  []model.Type
// @Router       /type/listOrder [get]
func (th *TypeHandler) ListOrder(ctx *gin.Context) {
	types, _ := typeService.ListOrderByAmountDesc()
	command.Success(ctx, "查询成功", gin.H{"types": types})
}

// @Summary      分类展示(不排序)
// @Tags         Example/Type
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  []model.Type
// @Router       /type/list [get]
func (th *TypeHandler) TypeList(ctx *gin.Context) {
	types, _ := typeService.List()
	command.Success(ctx, "查询成功", gin.H{"types": types})
}

// @Summary      按分类查询博客分页
// @Tags         Example/Type
// @Accept       json
// @Produce      json
// @Param        id	  		query     int	true	"分类id"
// @Param        pageSize	query     int	false	"每页显示的"
// @Param        pageNo	  	query     int	false	"页码"
// @Success 	 200  {object}  response.PageList	{data=response.PageList}
// @Router       /type/page [get]
func (th *TypeHandler) TypePageList(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Query("id"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "6"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "1"))

	if pageNo <= 0 {
		pageNo = 1
	}

	if pageSize <= 0 || pageSize > 6 {
		pageSize = 6
	}

	typeList, err := typeService.TypePageList(id, pageSize, pageNo)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"typeList": typeList})
}

// @Summary      新增分类
// @Tags         Example/Type
// @Accept       mpfd
// @Produce      json
// @Param        typeName formData string true  "insert typeName"
// @Success 	 200  {object}  string
// @Security	 X-Token
// @Router       /type/save [post]
func (th *TypeHandler) CreateType(ctx *gin.Context) {

	typeName := ctx.PostForm("typeName")
	if typeName == "" {
		global.QX_LOG.Error("typeName cannot be empty")
		command.Failed(ctx, http.StatusBadRequest, "typeName cannot be empty")
		return
	}

	if err := typeService.CreateType(typeName); err != nil {
		global.QX_LOG.Errorf("insert type err: %v", err)
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "添加成功", nil)
}
