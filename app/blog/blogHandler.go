package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/request"
	"github.com/qianxia/blog/utils"
)

type IBlogHandler interface {
	Save(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
	pageList(ctx *gin.Context)
}

type BlogHandler struct {
	Service BlogService
}

func NewBlogHandler() IBlogHandler {
	var blogService BlogService

	return BlogHandler{Service: blogService}
}

// 新增博客
func (b BlogHandler) Save(ctx *gin.Context) {
	var post request.Post
	ctx.ShouldBindJSON(&post)
	// 数据校验
	if post.Title == "" || len(post.Title) < 6 || len(post.Title) > 20 {
		command.Failed(ctx, http.StatusInternalServerError, "标题为空或标题长度少于6位")
		return
	}

	if post.Content == "" || len(post.Content) < 6 {
		command.Failed(ctx, http.StatusInternalServerError, "博客内容不能小于6位")
		return
	}

	if post.Flag == "" {
		command.Failed(ctx, http.StatusInternalServerError, "博客来源还未选择")
		return
	}

	if len(post.Tags) < 1 {
		command.Failed(ctx, http.StatusInternalServerError, "博客标签未选择")
		return
	}

	// 获取登录的用户信息
	userInfo, _ := utils.GetSession(ctx, "userInfo")
	post.UserId = userInfo.(model.User).Id

	fmt.Println("post ===> ", post)

	err := b.Service.Save(post)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", nil)
}

// 显示所有博客
func (b BlogHandler) List(ctx *gin.Context) {
	// 获取登录的用户信息
	userInfo, _ := utils.GetSession(ctx, "userInfo")
	blogs, err := b.Service.List(userInfo.(model.User).Id)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"blog": blogs})
}

// 个人博客删除
func (b BlogHandler) Delete(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)

	err := b.Service.Delete(id)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", nil)
}

// 查询博客显示在首页并分页
func (b BlogHandler) pageList(ctx *gin.Context) {

	pageMap := make(map[string]int)
	ctx.ShouldBindJSON(&pageMap)

	skipCount := (pageMap["pageNo"] - 1) * pageMap["pageSize"]
	pageMap["skipCount"] = skipCount

	pageListVO, err := b.Service.PageList(pageMap)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "查询成功", gin.H{"dataList": pageListVO})
}
