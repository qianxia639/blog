package blog

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/model"
)

type IBlogHandler interface {
	Save(ctx *gin.Context)
}

type BlogHandler struct {
	Service BlogService
}

func NewBlogHandler() IBlogHandler {
	blogService := NewBlogService()

	return BlogHandler{Service: blogService}
}

func (b BlogHandler) Save(ctx *gin.Context) {
	var blog model.Blog
	ctx.ShouldBindJSON(&blog)

	userId, _ := strconv.Atoi(ctx.Params.ByName("userId"))
	typeId, _ := strconv.Atoi(ctx.Params.ByName("typeId"))

	fmt.Printf("userId: %d,typeId: %d\n", userId, typeId)

	fmt.Println("blog ===> ", blog)

	b.Service.Save(blog)
}
