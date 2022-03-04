package system

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/service/system"
)

type SearchHandler struct {
	searchService system.SearchService
}

/**
* 搜索博客
 */
func (sh *SearchHandler) SearchBlog(ctx *gin.Context) {
	title := ctx.Param("title")
	sh.searchService.SearchBlog(title)
}
