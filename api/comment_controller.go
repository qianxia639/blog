package api

import (
	db "Blog/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type insertCommentRequest struct {
	BlogId    int64  `json:"blog_id"`
	CommentId int64  `json:"comment_id"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Content   string `json:"content"`
}

func (server *Server) insertComment(ctx *gin.Context) {
	var req insertCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error)
		return
	}

	arg := db.InsertCommentParams{
		BlogID:    req.BlogId,
		CommentID: req.CommentId,
		Nickname:  req.Nickname,
		Avatar:    req.Avatar,
		Content:   req.Content,
	}

	comment, err := server.store.InsertComment(ctx, arg)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.SecureJSON(http.StatusOK, comment)
}
