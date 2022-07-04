package example

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/service/example"
	"github.com/qianxia/blog/utils"
)

type BlogHandler struct {
	blogService example.BlogService
}

// @Summary      新增博客
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param        blog body request.SaveBlog  true  "Create Blog"
// @Success 	 200  {object}  string
// @Security	 X-Token
// @Router       /blog/save [post]
func (bh *BlogHandler) CreateBlog(ctx *gin.Context) {
	var saveBlog request.SaveBlog
	_ = ctx.ShouldBindJSON(&saveBlog)

	err := utils.Verify(saveBlog)
	if err != nil {
		global.LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, "缺少必要的参数")
		return
	}
	// 获取登录的用户信息
	userId := utils.GetUserId(ctx)

	blog, err := bh.blogService.SaveBlog(saveBlog, userId)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "博客发布失败")
		return
	}
	command.Success(ctx, "博客发布成功", gin.H{"blo": blog})
}

// @Summary      个人博客展示
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param        pageSize	query	int    false	"每页显示的"
// @Param        pageNo		query	int    false	"页码"
// @Success 	 200  {object}  response.PageList	{data=response.PageList}
// @Security	 X-Token
// @Router       /blog/list [get]
func (bh *BlogHandler) BlogList(ctx *gin.Context) {
	// 获取登录的用户信息
	userId := utils.GetUserId(ctx)
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "1"))

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 10 || pageSize > 10 {
		pageSize = 10
	}

	blogs, err := bh.blogService.List(userId, pageNo, pageSize)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"blogs": blogs})
}

// @Summary      个人博客删除
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param        id		path	int		true	"删除的博客id"
// @Success 	 200  {object}  string
// @Security	 X-Token
// @Router       /blog/{id} [delete]
func (bh *BlogHandler) DeleteBlog(ctx *gin.Context) {

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	userId := utils.GetUserId(ctx)
	err := bh.blogService.DeleteBlog(id, userId)
	if err != nil {
		global.LOG.Errorf("要删除的博客不存在或已删除: ", err)
		command.Failed(ctx, http.StatusInternalServerError, "删除失败")
		return
	}
	command.Success(ctx, "删除成功", nil)
}

// @Summary      修改博客
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param        blog	body	request.UpdateBlog	true	"Update Blog"
// @Success 	 200  {object}  string
// @Security	 X-Token
// @Router       /blog/update [put]
func (bh *BlogHandler) UpdateBlog(ctx *gin.Context) {

	var ub request.UpdateBlog

	_ = ctx.ShouldBindJSON(&ub)

	if err := utils.Verify(ub); err != nil {
		global.LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, "缺少必要的参数")
		return
	}

	if err := bh.blogService.UpdateBlog(ub); err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "修改失败")
		return
	}
	command.Success(ctx, "修改成功", nil)
}

// @Summary      首页博客展示
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param        pageSize	query	int		false	"每页显示的"
// @Param        pageNo		query	int		false	"页码"
// @Success 	 200  {object}  response.PageList	{data=response.PageList}
// @Router       /blog/pageList [get]
func (bh *BlogHandler) BlogPageList(ctx *gin.Context) {

	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "6"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "1"))

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 6 || pageSize > 6 {
		pageSize = 6
	}

	pageList, err := bh.blogService.PageList(pageSize, pageNo)

	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}

	command.Success(ctx, "查询成功", gin.H{"pageList": pageList})
}

// @Summary      最新推荐
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  []model.Blog	{data=[]model.Blog}
// @Router       /blog/latestList [get]
func (bh *BlogHandler) LatestList(ctx *gin.Context) {
	list, err := bh.blogService.LatestList()
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"latestList": list})
}

// @Summary      获取博客信息
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param		 id	  path  int   true    "根据id获取博客信息"
// @Success 	 200  {object} model.Blog
// @Router       /blog/{id} [get]
func (bh *BlogHandler) GetBlogInfo(ctx *gin.Context) {
	blogId, _ := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	if blogs, err := bh.blogService.GetBlogInfo(blogId); err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	} else {
		command.Success(ctx, "查询成功", gin.H{"blogs": blogs})
	}
}

// @Summary      增加浏览数
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param		 id	  query  int   true    "根据id获取博客信息"
// @Success 	 200  {object}	string
// @Router       /blog/incrViews [get]
func (bh *BlogHandler) IncrViews(ctx *gin.Context) {
	blogId, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	if err := bh.blogService.IncrViews(blogId); err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "操作成功", nil)
}

// @Summary      博客列表(所有博客)
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Success 	 200  {object}	response.PageList
// @Security	 X-Token
// @Router       /blog/all [get]
func (bh *BlogHandler) QueryAll(ctx *gin.Context) {
	blogList := bh.blogService.QueryAll()
	command.Success(ctx, "查询成功", gin.H{"blogList": blogList})
}

// @Summary      按flag分组显示
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  map[string]interface{}
// @Security	 X-Token
// @Router       /blog/flag/list [get]
func (bh *BlogHandler) GetBlogGroupByFlag(ctx *gin.Context) {
	list, err := bh.blogService.GetBlogGroupByFlag()
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"list": list})
}
