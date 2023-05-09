package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) getArticle(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		logs.Logs.Errorf("err: %s, param: %s", err.Error(), ctx.Param("id"))
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	if id < 1 {
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	article, err := server.store.GetArticle(ctx, id)
	switch err {
	case nil:
		result.Obj(ctx, article)
	case errors.NoRowsErr:
		logs.Logs.Error("Article Not Found err: ", err)
		result.Error(ctx, http.StatusNotFound, errors.NotExistsAtricleErr.Error())
	default:
		logs.Logs.Error("Get Article err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
	}
}
