package api

import (
	"Blog/core/result"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (server *Server) logout(ctx *gin.Context) {
	payload, err := server.readToken(ctx.Request)
	if err != nil {
		result.UnauthorizedError(ctx, err.Error())
		return
	}

	_ = server.rdb.Del(ctx, fmt.Sprintf("t_%s", payload.Username)).Err()

	result.OK(ctx, "Delete Successfully")
}
