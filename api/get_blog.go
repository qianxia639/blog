package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) getBlog(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		logs.Logs.Errorf("err: %s, param: %s", err.Error(), ctx.Param("id"))
		result.BadRequestError(ctx, errors.ParamErr.Error())
		return
	}

	blog, err := server.store.GetBlog(ctx, id)
	switch err {
	case nil:
		result.Obj(ctx, blog)
	case ErrNoRows:
		logs.Logs.Error("User Not Found err: ", err)
		result.Error(ctx, http.StatusNotFound, errors.NotExistsUserErr.Error())
	default:
		logs.Logs.Error("Get Blog err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
	}
}
