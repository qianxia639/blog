package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/service/system"
	"github.com/qianxia/blog/utils"
)

type LeaveHandler struct {
	leaveService system.LeaveService
}

// 显示所有留言记录
func (lh *LeaveHandler) All(ctx *gin.Context) {
	if l, err := lh.leaveService.All(); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	} else {
		command.Success(ctx, "查询成功", gin.H{"leave": l})
	}
}

// 新增留言记录
func (lh *LeaveHandler) Insert(ctx *gin.Context) {
	var l request.Leave
	_ = ctx.ShouldBindJSON(&l)

	if err := utils.Verify(l); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := lh.leaveService.Insert(model.Leave{Name: l.Name, Content: l.Content}); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "添加成功", nil)
}

// 删除留言
func (lh *LeaveHandler) Delete(ctx *gin.Context) {
	// id, _ := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	if err := lh.leaveService.Delete(id); err != nil {
		global.QX_LOG.Errorf("delete leave err: %v", err)
		command.Failed(ctx, http.StatusInternalServerError, "操作失败")
		return
	}
	command.Success(ctx, "操作成功", nil)
}
