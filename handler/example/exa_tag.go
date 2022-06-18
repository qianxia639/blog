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
// @Success 	 200  {object}  []model.Tag
// @Router       /tag/list [get]
func (th *TagHandler) TagList(ctx *gin.Context) {
	tags, _ := tagService.List()
	command.Success(ctx, "查询成功", tags)
}

// @Summary      新增标签
// @Tags         Example/Tag
// @Accept       mpfd
// @Produce      json
// @Param        tagName formData string true  "insert tagName"
// @Success 	 200  {object}  string
// @Security	 X-Token
// @Router       /tag/save [post]
func (th *TagHandler) CreateTag(ctx *gin.Context) {

	tagName := ctx.PostForm("tagName")
	if tagName == "" {
		global.QX_LOG.Error("tagName cannot be empty")
		command.Failed(ctx, http.StatusBadRequest, "tagName cannot be empty")
		return
	}

	if err := tagService.CreateTag(tagName); err != nil {
		global.QX_LOG.Errorf("insert tag err: %v", err)
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "添加成功", nil)
}
