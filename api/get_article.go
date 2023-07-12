package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (server *Server) getArticle(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		logs.Logger.Error("string to int", zap.Error(err), zap.String("id", ctx.Param("id")))
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
		logs.Logger.Error("Article Not Found err: ", zap.Error(err))
		result.Error(ctx, http.StatusNotFound, errors.NotExistsAtricleErr.Error())
	default:
		logs.Logger.Error("Get Article err: ", zap.Error(err))
		result.ServerError(ctx, errors.ServerErr.Error())
	}
}
