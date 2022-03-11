package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/service/system"
)

type SearchHandler struct {
	searchService system.SearchService
}

/**
* 搜索所有博客
 */
func (sh *SearchHandler) SearchBlog(ctx *gin.Context) {
	query := ctx.Query("query")
	if blogs, err := sh.searchService.SearchBlog(query); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "搜索失败")
		return
	} else {
		command.Success(ctx, "搜索成功", blogs)
	}
}

/**
* 搜索个人博客
 */
func (sh *SearchHandler) SearchPriBlog(ctx *gin.Context) {
	title := ctx.Query("title")
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("paginate", "10"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	if m, err := sh.searchService.SearchPriBlog(title, startDate, endDate, pageSize, pageNo); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "搜索失败")
		return
	} else {
		command.Success(ctx, "搜索成功", m)
	}
}
