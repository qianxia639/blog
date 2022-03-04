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

/**
* 按年份显示全部博客信息
 */
func (ah ArchiveHandler) ArchiveList(ctx *gin.Context) {

	if m, total, err := ah.archiveService.GetArchiveGroupByYear(); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	} else {
		command.Success(ctx, "查询成功", gin.H{
			"archiveList": m,
			"total":       total,
		})
	}

}
