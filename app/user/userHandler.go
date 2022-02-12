package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/response"
	"github.com/qianxia/blog/utils"
)

type IUserHandler interface {
	register(ctx *gin.Context)
	login(ctx *gin.Context)
	info(ctx *gin.Context)
	logout(ctx *gin.Context)
}

type UserHandler struct {
	Service UserService
}

func NewUserHandler() IUserHandler {
	var userService UserService

	return UserHandler{Service: userService}
}

// 注册
func (u UserHandler) register(ctx *gin.Context) {
	var user model.User
	// 绑定表单数据
	if err := ctx.ShouldBindJSON(&user); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	uh, err := u.Service.Register(user)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "注册成功", &uh)
}

// 登录
func (u UserHandler) login(ctx *gin.Context) {
	// 绑定表单参数
	var form model.User
	ctx.ShouldBindJSON(&form)
	user, err := u.Service.Login(form)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// 生成token
	token := utils.CreateToken(user.Id)
	command.Success(ctx, "登录成功", gin.H{"token": token})
}

// 获取用户信息
func (u UserHandler) info(ctx *gin.Context) {
	userInfo := ctx.MustGet("user")
	command.Success(ctx, "信息获取成功", gin.H{"user": response.ToUser(userInfo.(model.User))})
}

// 登出
func (u UserHandler) logout(ctx *gin.Context) {

	command.Success(ctx, "登出成功", nil)
}
