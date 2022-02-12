package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
)

type ITagHandler interface {
	tagList(ctx *gin.Context)
}

type TagHandler struct {
	Service TagService
}

func NewTagHandler() ITagHandler {
	var tagService TagService

	return TagHandler{Service: tagService}
}

func (t TagHandler) tagList(ctx *gin.Context) {
	tags, err := t.Service.List()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"tag": tags})
}
