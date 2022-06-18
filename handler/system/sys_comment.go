package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/utils"
)

type CommentHandler struct{}

// @Summary      父级评论
// @Tags         System/Comment
// @Accept       json
// @Produce      json
// @Param        ParentConment body request.ParentConment true "Create ParentComment"
// @Success 	 200  {object}  model.Comment
// @Router       /comment/parent [post]
func (*CommentHandler) ParentComment(ctx *gin.Context) {
	var pc request.ParentConment
	_ = ctx.ShouldBindJSON(&pc)
	if err := utils.Verify(pc); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := commentService.ParentComment(pc)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
	}

	command.Success(ctx, "评论成功", comment)
}

// @Summary      子级评论
// @Tags         System/Comment
// @Accept       json
// @Produce      json
// @Param        ChildComment body request.ChildComment true "Create ChildComment"
// @Success 	 200  {object}  model.Comment
// @Router       /comment/child [post]
func (*CommentHandler) ChildComment(ctx *gin.Context) {
	var cc request.ChildComment
	_ = ctx.ShouldBindJSON(&cc)
	if err := utils.Verify(cc); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := commentService.ChildComment(cc)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "评论成功", comment)
}

// @Summary      删除父级评论
// @Tags         System/Comment
// @Accept       json
// @Produce      json
// @Param        id query int true "Delete ParentComment"
// @Success 	 200  {object}  string
// @Router       /comment/parent [delete]
func (*CommentHandler) DeleteParentComment(ctx *gin.Context) {
	commentId, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	err := commentService.DeleteParentComment(commentId)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "评论删除成功", nil)
}

// @Summary      删除子级评论
// @Tags         System/Comment
// @Accept       json
// @Produce      json
// @Param        parentId query int true "Delete ChildComment"
// @Success 	 200  {object}  string
// @Router       /comment/child [delete]
func (*CommentHandler) DeleteChildComment(ctx *gin.Context) {
	parentId, _ := strconv.ParseUint(ctx.Query("parentId"), 10, 64)
	err := commentService.DeleteChildComment(parentId)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "评论删除成功", nil)
}
