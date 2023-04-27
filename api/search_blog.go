package api

import (
	"Blog/core/errors"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"fmt"

	"github.com/gin-gonic/gin"
)

const wildcard = "%%%s%%"

type searchBlogRequest struct {
	Query    string `form:"query"`
	PageNo   int32  `form:"page_no" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=1"`
}

func (server *Server) searchBlog(ctx *gin.Context) {
	var req searchBlogRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		result.BadRequestError(ctx, errors.ParamErr.Error())
		return
	}

	if req.PageNo < 1 {
		req.PageNo = 0
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	if req.PageSize > 100 {
		req.PageSize = 100
	}

	arg := db.SearchBlogParams{
		Title:  fmt.Sprintf(wildcard, req.Query),
		Limit:  req.PageSize,
		Offset: (req.PageNo - 1) * req.PageSize,
	}

	blogs, err := server.store.SearchBlog(ctx, arg)
	if err != nil {
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.Obj(ctx, blogs)
}
