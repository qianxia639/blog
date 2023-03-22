package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) searchBlog(ctx *gin.Context) {
	title := ctx.Query("title")
	title = fmt.Sprintf("%%%s%%", title)
	blogs, err := server.store.SearchBlog(ctx, title)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}
