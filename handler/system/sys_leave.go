package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/service/system"
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
	var l model.Leave
	if err := ctx.ShouldBindJSON(&l); err != nil {
		global.QX_LOG.Errorf("bind parame err: %v", err)
		command.Failed(ctx, http.StatusBadRequest, "缺少必要的参数")
		return
	}

	if l.Content == "" {
		command.Failed(ctx, http.StatusBadRequest, "缺少必要的参数")
		return
	}

	if err := lh.leaveService.Insert(l); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "添加成功", nil)

}
