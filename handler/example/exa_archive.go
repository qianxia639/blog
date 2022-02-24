package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/service/example"
)

type ArchiveHandler struct {
	archiveService example.ArchiveService
}

// 按年份显示全部博客信息并分页
func (ah ArchiveHandler) ArchivePageList(ctx *gin.Context) {

	pageMap := make(map[string]int, 3)
	if err := ctx.ShouldBindJSON(&pageMap); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		global.RY_LOG.Warnf("%s-{%v}", "数据绑定失败", err)
		return
	}

	switch {
	case pageMap["pageSize"] == 0:
		pageMap["pageSize"] = 10
	case pageMap["pageNo"] == 0:
		pageMap["pageNo"] = 1
	}

	skipCount := (pageMap["pageNo"] - 1) * pageMap["pageSize"]
	pageMap["skipCount"] = skipCount

	if m, err := ah.archiveService.ArchivePageList(pageMap); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		global.RY_LOG.Warn(err)
		return
	} else {
		command.Success(ctx, "成功", gin.H{"archiveList": m})
	}

}
