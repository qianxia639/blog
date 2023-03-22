package api

import (
	db "Blog/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getCommentRequest struct {
	Id int64 `form:"id"`
}

func (server *Server) getComments(ctx *gin.Context) {
	var req getCommentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	comments, err := server.store.GetComments(ctx, req.Id)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	cs := make([]db.Comment, 0)
	for _, comment := range comments {
		arg := db.GetChildCommentsParams{
			OwnerID:  comment.OwnerID,
			ParentID: comment.ID,
		}
		cs, err = server.store.GetChildComments(ctx, arg)
		if err != nil {
			ctx.SecureJSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	ctx.JSON(http.StatusOK, cs)
}
