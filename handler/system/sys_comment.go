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
		command.RFailed(ctx, 500, err.Error())
	}

	if err := ch.commentService.Save(m); err != nil {
		command.RFailed(ctx, 500, err.Error())
	} else {
		command.Success(ctx, "成功", nil)
	}
}
