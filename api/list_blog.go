package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	db "Blog/db/sqlc"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type listBlogsRequest struct {
	PageNo   int32 `form:"page_no" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1"`
}

type pageResponse struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

func (server *Server) listBlogs(ctx *gin.Context) {

	var req listBlogsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logs.Logs.Error(err)
		ctx.JSON(http.StatusBadRequest, errors.ParamErr)
		return
	}

	var resp pageResponse
	offset := (req.PageNo - 1) * req.PageSize

	var wg sync.WaitGroup
	ctxCopy := ctx.Copy()

	wg.Add(2)
	go func() {
		defer wg.Done()
		data, err := server.store.ListBlogs(ctxCopy, db.ListBlogsParams{
			Limit:  req.PageSize,
			Offset: offset,
		})
		if err != nil {
			logs.Logs.Error(err)
			ctxCopy.JSON(http.StatusInternalServerError, errors.ServerErr)
			return
		}
		resp.Data = data
	}()

	go func() {
		defer wg.Done()
		total, err := server.store.CountBlog(ctxCopy)
		if err != nil {
			logs.Logs.Error(err)
			ctxCopy.JSON(http.StatusInternalServerError, errors.ServerErr)
			return
		}
		resp.Total = total
	}()
	wg.Wait()

	ctx.JSON(http.StatusOK, resp)
}
