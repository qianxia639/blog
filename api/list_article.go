package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

const wildcard = "%%%s%%"

type listArticlesRequest struct {
	Query    string `form:"query"`
	PageNo   int32  `form:"page_no" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=1"`
}

type pageResponse struct {
	PageNo   int32       `json:"page_no"`
	PageSize int32       `json:"page_size"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}

func (server *Server) listArticles(ctx *gin.Context) {

	var req listArticlesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logs.Logs.Error(err)
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	if req.PageNo < 1 {
		req.PageNo = 1
	}

	if req.PageSize < 1 {
		req.PageSize = 10
	}

	if req.PageSize > 100 {
		req.PageSize = 100
	}

	var resp pageResponse
	resp.PageNo = req.PageNo
	resp.PageSize = req.PageSize

	offset := (req.PageNo - 1) * req.PageSize

	var wg sync.WaitGroup
	ctxCopy := ctx.Copy()

	wg.Add(2)
	go func() {
		defer wg.Done()
		data, err := server.store.ListArticles(ctxCopy, &db.ListArticlesParams{
			Title:  fmt.Sprintf(wildcard, req.Query),
			Limit:  req.PageSize,
			Offset: offset,
		})
		if err != nil {
			logs.Logs.Error(err)
			result.ServerError(ctxCopy, errors.ServerErr.Error())
			return
		}
		resp.Data = data
	}()

	go func() {
		defer wg.Done()
		total, err := server.store.CountArticle(ctxCopy)
		if err != nil {
			logs.Logs.Error(err)
			result.ServerError(ctxCopy, errors.ServerErr.Error())
			return
		}
		resp.Total = total
	}()
	wg.Wait()

	result.Obj(ctx, resp)
}
