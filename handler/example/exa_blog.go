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

/**
* 新增博客
 */
func (bh BlogHandler) CreateBlog(ctx *gin.Context) {
	var post request.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		command.Failed(ctx, http.StatusBadRequest, "缺少必要的参数")
		global.QX_LOG.Errorf("parame bind err:", err)
		return
	}
	// 获取登录的用户信息
	post.UserId = utils.GetUserId(ctx)
	post.Username = utils.GetUsername(ctx)

	err := bh.blogService.Save(post)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "发布博客失败")
		return
	}
	command.Success(ctx, "发布博客成功", nil)
}

/**
* 个人博客展示
 */
func (bh BlogHandler) BlogList(ctx *gin.Context) {
	// 获取登录的用户信息
	userId := utils.GetUserId(ctx)

	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))

	blogs, err := bh.blogService.List(userId, pageNum, pageSize)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", blogs)
}

/**
* 个人博客删除
 */
func (bh BlogHandler) DeleteBlog(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)

	err := bh.blogService.Delete(id)
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "删除失败")
		return
	}
	command.Success(ctx, "删除成功", nil)
}

/**
* 修改博客
 */
func (bh *BlogHandler) UpdateBlog(ctx *gin.Context) {

	var post request.Post

	ctx.ShouldBindJSON(&post)

	if err := bh.blogService.Update(post); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "修改失败")
		return
	}
	command.Success(ctx, "修改成功", nil)
}

/**
* 查询博客显示在首页并分页
 */
func (bh BlogHandler) BlogPageList(ctx *gin.Context) {

	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

	pageList, err := bh.blogService.PageList(pageSize, pageNum)

	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}

	command.Success(ctx, "查询成功", pageList)
}

/**
* 最新推荐
 */
func (bh BlogHandler) LatestList(ctx *gin.Context) {
	list, err := bh.blogService.LatestList()
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	}
	command.Success(ctx, "查询成功", gin.H{"latestList": list})
}

/**
* 根据id获取博客信息
 */
func (bh BlogHandler) GetBlog(ctx *gin.Context) {
	blogId, _ := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	avatar := utils.GetAvatar(ctx)
	if blogs, err := bh.blogService.GetBlog(blogId, avatar); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	} else {
		command.Success(ctx, "查询成功", gin.H{"blogs": blogs})
	}
}

/**
* 获取要编辑的博客的信息
 */
func (bh *BlogHandler) GetUpdateBlog(ctx *gin.Context) {
	blogId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if blogs, err := bh.blogService.GetUpdateBlog(blogId); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "查询失败")
		return
	} else {
		command.Success(ctx, "查询成功", gin.H{"blogs": blogs})
	}
}
