package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createCritiqueRequest struct {
	OwnerId  int64  `json:"owner_id" binding:"required"`
	ParentId int64  `json:"parent_id" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

func (server *Server) createCritique(ctx *gin.Context) {
	var req createCritiqueRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error(err)
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	arg := &db.CreateCritiqueParams{
		OwnerID:  req.OwnerId,
		ParentID: req.ParentId,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Content:  req.Content,
	}

	critique, err := server.store.CreateCritique(ctx, arg)
	if err != nil {
		logs.Logs.Error(err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}
	result.Obj(ctx, critique)
}
