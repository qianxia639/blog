package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getBlogRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBlog(ctx *gin.Context) {
	var req getBlogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	blog, err := server.store.GetBlog(ctx, req.Id)
	if err != nil {
		if err == ErrNoRows {
			ctx.SecureJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, blog)
}
