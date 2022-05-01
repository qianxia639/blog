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
// @Param        blog body request.Post  true  "Create Blog"
// @Success 	 200  {object}  string
// @Security	 X-Token
// @Router       /blog/save [post]
func (bh BlogHandler) CreateBlog(ctx *gin.Context) {
	var post request.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.RFailed(ctx, http.StatusBadRequest, "缺少必要的参数")
	}
	// 获取登录的用户信息
	post.UserId = utils.GetUserId(ctx)
	post.Username = utils.GetUsername(ctx)

	err := bh.blogService.Save(post)
	if err != nil {
		global.QX_LOG.Error(err)
		command.RFailed(ctx, http.StatusInternalServerError, "发布博客失败")
	}
	command.Success(ctx, "发布博客成功", nil)
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
func (bh BlogHandler) BlogList(ctx *gin.Context) {
	// 获取登录的用户信息
	userId := utils.GetUserId(ctx)

	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "1"))

	blogs, err := bh.blogService.List(userId, pageNo, pageSize)
	if err != nil {
		global.QX_LOG.Error(err)
		command.RFailed(ctx, http.StatusInternalServerError, "查询失败")
	}
	command.Success(ctx, "查询成功", blogs)
}

// @Summary      个人博客删除
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param        id		path	int		true	"删除的博客id"
// @Success 	 200  {object}  string
// @Security	 X-Token
// @Router       /blog/{id} [delete]
func (bh BlogHandler) DeleteBlog(ctx *gin.Context) {

	id, _ := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)

	err := bh.blogService.Delete(id)
	if err != nil {
		global.QX_LOG.Error(err)
		command.RFailed(ctx, http.StatusInternalServerError, "删除失败")
	}
	command.Success(ctx, "删除成功", nil)
}

// @Summary      修改博客
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param        blog	body	request.Post	true	"Update Blog"
// @Success 	 200  {object}  string
// @Security	 X-Token
// @Router       /blog/update [put]
func (bh *BlogHandler) UpdateBlog(ctx *gin.Context) {

	var post request.Post

	ctx.ShouldBindJSON(&post)

	if err := bh.blogService.Update(post); err != nil {
		global.QX_LOG.Error(err)
		command.RFailed(ctx, http.StatusInternalServerError, "修改失败")
	}
	command.Success(ctx, "修改成功", nil)
}

// @Summary      博客展示
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param        pageSize	query	int		false	"每页显示的"
// @Param        pageNo		query	int		false	"页码"
// @Success 	 200  {object}  response.PageList	{data=response.PageList}
// @Router       /blog/pageList [get]
func (bh BlogHandler) BlogPageList(ctx *gin.Context) {

	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNo, _ := strconv.Atoi(ctx.Query("pageNo"))

	pageList, err := bh.blogService.PageList(pageSize, pageNo)

	if err != nil {
		global.QX_LOG.Error(err)
		command.RFailed(ctx, http.StatusInternalServerError, "查询失败")
	}

	command.Success(ctx, "查询成功", pageList)
}

// @Summary      最新推荐
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  []model.Blog	{data=[]model.Blog}
// @Router       /blog/latestList [get]
func (bh BlogHandler) LatestList(ctx *gin.Context) {
	list, err := bh.blogService.LatestList()
	if err != nil {
		global.QX_LOG.Error(err)
		command.RFailed(ctx, http.StatusInternalServerError, "查询失败")
	}
	command.Success(ctx, "查询成功", gin.H{"latestList": list})
}

// @Summary      获取博客信息
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param		 id	  path  int   true    "根据id获取博客信息"
// @Success 	 200  {object}  map[string]interface{}
// @Router       /blog/{id} [get]
func (bh BlogHandler) GetBlog(ctx *gin.Context) {
	blogId, _ := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	avatar := utils.GetAvatar(ctx)
	if blogs, err := bh.blogService.GetBlog(blogId, avatar); err != nil {
		global.QX_LOG.Error(err)
		command.RFailed(ctx, http.StatusInternalServerError, "查询失败")
	} else {
		command.Success(ctx, "查询成功", gin.H{"blogs": blogs})
	}
}

// @Summary      获取编辑的博客信息
// @Tags         Example/Blog
// @Accept       json
// @Produce      json
// @Param		 id	  path  int   true    "根据id获取博客信息"
// @Success 	 200  {object}  map[string]interface{}
// @security	 X-Token
// @Router       /blog/update/{id} [get]
func (bh *BlogHandler) GetUpdateBlog(ctx *gin.Context) {
	blogId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if blogs, err := bh.blogService.GetUpdateBlog(blogId); err != nil {
		global.QX_LOG.Error(err)
		command.RFailed(ctx, http.StatusInternalServerError, "查询失败")
	} else {
		command.Success(ctx, "查询成功", gin.H{"blogs": blogs})
	}
}
