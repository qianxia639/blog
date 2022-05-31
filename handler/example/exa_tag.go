package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
)

type TagHandler struct{}

// @Summary      标签列表
// @Tags         Example/Tag
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  []string
// @Router       /tag/list [get]
func (th *TagHandler) TagList(ctx *gin.Context) {
	tags, err := tagService.List()
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"tag": tags})
}
