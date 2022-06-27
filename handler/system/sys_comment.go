package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model/request"
)

type CommentHandler struct{}

// @Summary      添加评论
// @Tags         System/Comment
// @Accept       json
// @Produce      json
// @Param        Comment body request.Comment true "Create Comment"
// @Success 	 200  {object}  model.Comment
// @Router       /comment/save [post]
func (*CommentHandler) SaveComment(ctx *gin.Context) {

	var comment request.Comment
	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusBadRequest, "缺少必要的参数")
		return
	}
	c, err := commentService.Save(comment)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "评论成功", gin.H{"comment": c})
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
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "评论删除成功", nil)
}

// @Summary      删除子级评论
// @Tags         System/Comment
// @Accept       json
// @Produce      json
// @Param        id query int true "Delete ChildComment"
// @Success 	 200  {object}  string
// @Router       /comment/child [delete]
func (*CommentHandler) DeleteChildComment(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	err := commentService.DeleteChildComment(id)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "评论删除成功", nil)
}

// @Summary      评论列表
// @Tags         System/Comment
// @Accept       json
// @Produce      json
// @Param        id query int true "Get Comment List"
// @Success 	 200  {object}  []model.Comment
// @Router       /comment/list [get]
func (*CommentHandler) CommentList(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	comments, err := commentService.List(id)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"comments": comments})
}
