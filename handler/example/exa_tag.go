package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/service/example"
)

type TagHandler struct {
	tagService example.TagService
}

func (th TagHandler) TagList(ctx *gin.Context) {
	tags, err := th.tagService.List()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"tag": tags})
}
