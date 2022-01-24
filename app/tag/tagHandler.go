package app

import (
	"github.com/gin-gonic/gin"
)

type ITagHandler interface {
	List(ctx *gin.Context)
}

type TagHandler struct {
	Service TagService
}

/*
func NewTagHandler() ITagHandler {
	var tagService TagService

	return TagHandler{Service: tagService}
}

func (t TagHandler) List(ctx *gin.Context) {
	tags, err := t.Service.List()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"tag": tags})
}
*/
