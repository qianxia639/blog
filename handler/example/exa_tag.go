package example

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
)

type TagHandler struct{}

// @Summary      标签列表
// @Tags         Example/Tag
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  []model.Tag
// @Router       /tag/list [get]
func (th *TagHandler) TagList(ctx *gin.Context) {
	tags, _ := tagService.List()
	command.Success(ctx, "查询成功", gin.H{"tag": tags})
}
