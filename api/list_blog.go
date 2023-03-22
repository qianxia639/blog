package api

import (
	db "Blog/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listBlogsRequest struct {
	PageNo   int32 `form:"page_no" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1"`
}

func (server *Server) listBlogs(ctx *gin.Context) {
	var req listBlogsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}
	blogs, err := server.store.ListBlogs(ctx, db.ListBlogsParams{
		Limit:  req.PageSize,
		Offset: (req.PageNo - 1) * req.PageSize,
	})
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, blogs)
}
