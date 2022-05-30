package system

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/service/system"
)

type CommentHandler struct {
	commentService system.CommentService
}

func (ch *CommentHandler) CommentList(ctx *gin.Context) {

}

func (ch *CommentHandler) Save(ctx *gin.Context) {
	var m map[string]interface{}
	if err := ctx.ShouldBindJSON(&m); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, 500, err.Error())
		return
	}

	if err := ch.commentService.Save(m); err != nil {
		command.Failed(ctx, 500, err.Error())
		return
	} else {
		command.Success(ctx, "成功", nil)
	}
}
