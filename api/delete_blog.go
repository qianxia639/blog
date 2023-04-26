package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) deleteBlog(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		logs.Logs.Errorf("err: %s, param: %s", err.Error(), ctx.Param("id"))
		result.BadRequestError(ctx, errors.ParamErr.Error())
		return
	}

	_, err = server.store.GetBlog(ctx, id)
	if err != nil {
		if err == ErrNoRows {
			logs.Logs.Error("Get Blog err: ", err)
			result.Error(ctx, http.StatusNotFound, errors.NotExistsUserErr.Error())
			return
		}
		logs.Logs.Error("Get Blog err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	err = server.store.DeleteBlog(ctx, id)
	if err != nil {
		logs.Logs.Error("Delete Blog err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.OK(ctx, "Delete Blog Successfully")
}
