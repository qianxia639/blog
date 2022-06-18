package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/utils"
)

type SearchHandler struct{}

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

	blogs, err := searchService.SearchBlog(query, pageNo, pageSize)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "搜索成功", blogs)
}

// @Summary      查询个人博客
// @Tags         System/Search
// @Accept       json
// @Produce      json
// @Param        title	 	query	string    false  "标题"
// @Param        startDate	query	string    false  "起始时间"
// @Param        endDate	query	string    false  "结束时间"
// @Param        pageSize	query	int  	  false  "每页显示"
// @Param        pageNo		query	int  	  false  "页数"
// @Success 	 200  	{object}  	response.PageList 	{data=response.PageList}
// @Security 	 X-Token
// @Router       /search/priblog [get]
func (sh *SearchHandler) SearchPriBlog(ctx *gin.Context) {
	title := ctx.Query("title")
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "1"))
	userId := utils.GetUserId(ctx)

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize > 10 || pageSize < 10 {
		pageSize = 10
	}

	m, err := searchService.SearchPriBlog(title, startDate, endDate, pageSize, pageNo, userId)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "搜索失败")
		return
	}
	command.Success(ctx, "搜索成功", m)

}
