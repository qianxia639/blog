package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) incrViews(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		logs.Logs.Errorf("err: %s, param: %s", err.Error(), ctx.Param("id"))
		result.BadRequestError(ctx, errors.ParamErr.Error())
		return
	}

	if id < 1 {
		result.BadRequestError(ctx, errors.ParamErr.Error())
		return
	}

	_, err = server.store.GetBlog(ctx, id)
	if err != nil {
		if err == ErrNoRows {
			logs.Logs.Error("User Not Found err: ", err)
			result.Error(ctx, http.StatusNotFound, errors.NotExistsUserErr.Error())
			return
		}
		logs.Logs.Error("Get Blog err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	err = server.store.IncrViews(ctx, id)
	if err != nil {
		logs.Logs.Error("Incr Views err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.OK(ctx, "Increment Views Successfully")
}
