package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/request"
)

type IBlogHandler interface {
	createBlog(ctx *gin.Context)
	blogList(ctx *gin.Context)
	deleteBlog(ctx *gin.Context)
	blogPageList(ctx *gin.Context)
}

type BlogHandler struct {
	Service BlogService
}

func NewBlogHandler() IBlogHandler {
	var blogService BlogService

	return BlogHandler{Service: blogService}
}

// 新增博客
func (b BlogHandler) createBlog(ctx *gin.Context) {
	var post request.Post
	ctx.ShouldBindJSON(&post)

	// 获取登录的用户信息
	userInfo := ctx.MustGet("user")
	post.UserId = userInfo.(model.User).Id

	err := b.Service.Save(post)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", nil)
}

// 显示所有博客
func (b BlogHandler) blogList(ctx *gin.Context) {
	// 获取登录的用户信息
	userInfo := ctx.MustGet("user")
	blogs, err := b.Service.List(userInfo.(model.User).Id)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"blog": blogs})
}

// 个人博客删除
func (b BlogHandler) deleteBlog(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)

	err := b.Service.Delete(id)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", nil)
}

// 查询博客显示在首页并分页
func (b BlogHandler) blogPageList(ctx *gin.Context) {

	pageMap := make(map[string]int)
	ctx.ShouldBindJSON(&pageMap)

	switch {
	case pageMap["pageSize"] == 0:
		pageMap["pageSize"] = -1
	case pageMap["pageNo"] == 0:
		pageMap["pageNo"] = 1
	}

	skipCount := (pageMap["pageNo"] - 1) * pageMap["pageSize"]
	pageMap["skipCount"] = skipCount

	pageListVO, err := b.Service.PageList(pageMap)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "查询成功", gin.H{"dataList": pageListVO})
}
