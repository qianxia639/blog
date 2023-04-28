package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createCommentRequest struct {
	OwnerId  int64  `json:"owner_id" binding:"required"`
	ParentId int64  `json:"parent_id" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

func (server *Server) createComment(ctx *gin.Context) {
	var req createCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error(err)
		result.BadRequestError(ctx, errors.ParamErr.Error())
		return
	}

	arg := &db.CreateCommentParams{
		OwnerID:  req.OwnerId,
		ParentID: req.ParentId,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Content:  req.Content,
	}

	comment, err := server.store.CreateComment(ctx, arg)
	if err != nil {
		logs.Logs.Error(err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}
	result.Obj(ctx, comment)
}
