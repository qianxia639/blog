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
		command.Failed(ctx, http.StatusInternalServerError, "数据绑定失败")
		global.QX_LOG.Errorf("%s-{%v}", "数据绑定失败", err)
		return
	}
	// 获取登录的用户信息
	userId := utils.GetUserId(ctx)
	post.UserId = userId

	err := bh.blogService.Save(post)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", nil)
}

/**
* 保存博客
 */
func (bh BlogHandler) SaveBlog(ctx *gin.Context) {
	var post request.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, "数据绑定失败")
		global.QX_LOG.Errorf("%s-{%v}", "数据绑定失败", err)
		return
	}
	// 获取登录的用户信息
	userId := utils.GetUserId(ctx)
	post.UserId = userId

	err := bh.blogService.SaveBlog(post)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", nil)
}

/**
* 个人博客展示
 */
func (bh BlogHandler) BlogList(ctx *gin.Context) {
	// 获取登录的用户信息
	userId := utils.GetUserId(ctx)

	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("paginate", "10"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	offset := (pageNo - 1) * pageSize

	pageMap := make(map[string]int, 3)
	pageMap["pageSize"] = pageSize
	pageMap["pageNo"] = pageNo
	pageMap["offset"] = offset

	blogs, err := bh.blogService.List(userId, pageMap)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
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
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "操作成功", nil)
}

/**
* 修改博客
 */
func (bh *BlogHandler) UpdateBlog(ctx *gin.Context) {

	var m map[string]interface{}

	ctx.ShouldBindJSON(&m)

	mb := m["params"]
	b := mb.(map[string]interface{})["blog"]
	id := mb.(map[string]interface{})["id"]

	if err := bh.blogService.Update(uint64(id.(float64)), b); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, command.OperationSuccess, nil)
}

/**
* 查询博客显示在首页并分页
 */
func (bh BlogHandler) BlogPageList(ctx *gin.Context) {

	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("paginate", "6"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	offset := (pageNo - 1) * pageSize

	pageMap := make(map[string]int, 3)

	pageMap["pageSize"] = pageSize
	pageMap["pageNo"] = pageNo

	pageMap["offset"] = offset

	pageList, err := bh.blogService.PageList(pageMap)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
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
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "查询成功", gin.H{"latestList": list})
}

/**
* 根据id获取博客信息
 */
func (bh BlogHandler) GetBlog(ctx *gin.Context) {
	blogId, _ := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	if blogs, err := bh.blogService.GetBlog(blogId); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
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
		command.Failed(ctx, http.StatusInternalServerError, command.FindFailed)
		return
	} else {
		command.Success(ctx, command.FindSuccess, gin.H{"blogs": blogs})
	}
}
