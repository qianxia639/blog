package example

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/request"
	"github.com/qianxia/blog/service/example"
)

type BlogHandler struct {
	blogService example.BlogService
}

// 新增博客
func (bh BlogHandler) CreateBlog(ctx *gin.Context) {
	var post request.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, "数据绑定失败")
		global.RY_LOG.Errorf("%s-{%v}", "数据绑定失败", err)
		return
	}
	// 获取登录的用户信息
	userInfo := ctx.MustGet("user")
	post.UserId = userInfo.(model.User).Id

	err := bh.blogService.Save(post)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", nil)
}

// 显示所有博客
func (bh BlogHandler) BlogList(ctx *gin.Context) {
	// 获取登录的用户信息
	userInfo := ctx.MustGet("user")

	pageSize, _ := strconv.Atoi(ctx.Query("paginate"))
	pageNo, _ := strconv.Atoi(ctx.Query("page"))
	offset := (pageNo - 1) * pageSize

	pageMap := make(map[string]int, 3)
	pageMap["pageSize"] = pageSize
	pageMap["pageNo"] = pageNo
	pageMap["offset"] = offset

	switch {
	case pageMap["pageSize"] == 0:
		pageMap["pageSize"] = 5
	case pageMap["pageNo"] == 0:
		pageMap["pageNo"] = 1
	}

	blogs, err := bh.blogService.List(userInfo.(model.User).Id, pageMap)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"pageList": blogs})
}

// 个人博客删除
func (bh BlogHandler) DeleteBlog(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)

	err := bh.blogService.Delete(id)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", nil)
}

// 查询博客显示在首页并分页
func (bh BlogHandler) BlogPageList(ctx *gin.Context) {

	pageSize, _ := strconv.Atoi(ctx.Query("paginate"))
	pageNo, _ := strconv.Atoi(ctx.Query("page"))
	offset := (pageNo - 1) * pageSize

	pageMap := make(map[string]int, 3)

	pageMap["pageSize"] = pageSize
	pageMap["pageNo"] = pageNo

	switch {
	case pageMap["pageSize"] == 0:
		pageMap["pageSize"] = 6
	case pageMap["pageNo"] == 0:
		pageMap["pageNo"] = 1
	}
	pageMap["offset"] = offset

	pageList, err := bh.blogService.PageList(pageMap)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "查询成功", gin.H{"pageList": pageList})
}

// 最新推荐
func (bh BlogHandler) LatestList(ctx *gin.Context) {
	list, err := bh.blogService.LatestList()
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"latestList": list})
}

// 根据id获取博客信息
func (bh BlogHandler) GetBlog(ctx *gin.Context) {
	blogId, _ := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if blogs, err := bh.blogService.GetBlog(uint64(blogId)); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	} else {
		command.Success(ctx, "查询成功", gin.H{"blogs": blogs})
	}
}
