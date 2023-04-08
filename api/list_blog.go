package api

import (
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
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	var resp pageResponse

	offset := (req.PageNo - 1) * req.PageSize

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		data, err := server.store.ListBlogs(ctx, db.ListBlogsParams{
			Limit:  req.PageSize,
			Offset: offset,
		})
		if err != nil {
			ctx.SecureJSON(http.StatusInternalServerError, err.Error())
			return
		}
		resp.Data = data
	}()

	go func() {
		defer wg.Done()
		total, err := server.store.CountBlog(ctx)
		if err != nil {
			ctx.SecureJSON(http.StatusInternalServerError, err.Error())
			return
		}
		resp.Total = total
	}()

	wg.Wait()

	ctx.JSON(http.StatusOK, resp)
}
