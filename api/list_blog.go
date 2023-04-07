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
	var err error

	offset := (req.PageNo - 1) * req.PageSize

	server.wg.Add(2)
	go func() {
		defer server.wg.Done()
		resp.Data, err = server.store.ListBlogs(ctx, db.ListBlogsParams{
			Limit:  req.PageSize,
			Offset: offset,
		})

		// if err != nil {
		// 	ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		// 	return
		// }
	}()

	go func() {
		defer server.wg.Done()
		resp.Total, err = server.store.CountBlog(ctx)
		// if err != nil {
		// 	ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		// 	return
		// }
	}()

	server.wg.Wait()

	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// resp := pageResponse{
	// 	Total: total,
	// 	Data:  blogs,
	// }

	ctx.JSON(http.StatusOK, resp)
}
