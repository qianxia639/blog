package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
)

type ArchiveHandler struct{}

// @Summary      按年份显示全部博客信息
// @Tags         Example/Archive
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  map[string]interface{}
// @Router       /archive/list [get]
func (ah *ArchiveHandler) ArchiveList(ctx *gin.Context) {

	if m, total, err := archiveService.GetArchiveGroupByYear(); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	} else {
		command.Success(ctx, "查询成功", gin.H{
			"archiveList": m,
			"total":       total,
		})
	}

}
