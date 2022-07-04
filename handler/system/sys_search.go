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

// @Summary      查询所有博客
// @Tags         System/Search
// @Accept       json
// @Produce      json
// @Param        query		query	string    true   "标题"
// @Param        pageSize	query	int  	  false  "每页显示"
// @Param        pageNo		query	int  	  false  "页数"
// @Success 	 200  	{object}  	response.PageList 	{data=response.PageList}
// @Router       /search/blog [get]
func (sh *SearchHandler) SearchBlog(ctx *gin.Context) {
	query := ctx.Query("query")
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "1"))

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 10 || pageSize > 10 {
		pageNo = 10
	}

	blogs, err := sh.searchService.SearchBlog(query, pageNo, pageSize)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "搜索成功", gin.H{"pageList": blogs})
}
