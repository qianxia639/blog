package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/service/system"
)

type SearchHandler struct {
	searchService system.SearchService
}

/**
* 搜索博客
 */
func (sh *SearchHandler) SearchBlog(ctx *gin.Context) {
	query := ctx.Query("query")
	if blogs, err := sh.searchService.SearchBlog(query); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	} else {
		command.Success(ctx, "查询成功", blogs)
	}
}
