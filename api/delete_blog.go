package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type deleteBlogRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteBlog(ctx *gin.Context) {
	var req deleteBlogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := server.store.GetBlog(ctx, req.Id)
	if err != nil {
		if err == ErrNoRows {
			ctx.SecureJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = server.store.DeleteBlog(ctx, req.Id)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.SecureJSON(http.StatusOK, "Delete Blog Successfully")
}
