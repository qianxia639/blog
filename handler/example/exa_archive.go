package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/service/example"
)

type ArchiveHandler struct {
	archiveService example.ArchiveService
}

// 按年份显示全部博客信息并分页
func (ah *ArchiveHandler) ArchivePageList(ctx *gin.Context) {

	pageMap := make(map[string]int, 3)
	ctx.ShouldBindJSON(&pageMap)

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
		return
	} else {
		command.Success(ctx, "成功", gin.H{"archiveList": m})
	}

}
