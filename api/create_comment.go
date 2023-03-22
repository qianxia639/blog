package api

import (
	db "Blog/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCommentRequest struct {
	OwnerId  int64  `json:"owner_id"`
	ParentId int64  `json:"parent_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Content  string `json:"content"`
}

func (server *Server) createComment(ctx *gin.Context) {
	var req createCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error)
		return
	}

	arg := db.CreateCommentParams{
		OwnerID:  req.OwnerId,
		ParentID: req.ParentId,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Content:  req.Content,
	}

	comment, err := server.store.CreateComment(ctx, arg)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.SecureJSON(http.StatusOK, comment)
}
