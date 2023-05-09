package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) deleteArticle(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		logs.Logs.Errorf("err: %s, param: %s", err.Error(), ctx.Param("id"))
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	if id < 1 {
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	_, err = server.store.GetArticle(ctx, id)
	if err != nil {
		if err == errors.NoRowsErr {
			logs.Logs.Error("Get Article err: ", err)
			result.Error(ctx, http.StatusNotFound, errors.NotExistsUserErr.Error())
			return
		}
		logs.Logs.Error("Get Article err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	err = server.store.DeleteArticle(ctx, id)
	if err != nil {
		logs.Logs.Error("Delete Article err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.OK(ctx, "Delete Article Successfully")
}
