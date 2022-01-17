package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/dto"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
	"golang.org/x/crypto/bcrypt"
)

type IUserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Info(ctx *gin.Context)
}

type UserHandler struct {
	Service UserService
}

func NewUserHandler() IUserHandler {
	var userService UserService

	return UserHandler{Service: userService}
}

// 注册
func (u UserHandler) Register(ctx *gin.Context) {
	var user model.User
	// 绑定表单数据
	ctx.ShouldBindJSON(&user)

	_, err := u.Service.Register(user)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "注册成功", nil)
}

// 登录
func (u UserHandler) Login(ctx *gin.Context) {
	// 绑定表单参数
	var form model.User
	ctx.ShouldBindJSON(&form)

	user, err := u.Service.Login(form)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		command.Failed(ctx, http.StatusUnauthorized, "密码错误")
		return
	}

	// 生成token
	token := utils.CreateToken(user.Id)
	command.Success(ctx, "登录成功", gin.H{"token": token})
}

// 获取用户信息
func (u UserHandler) Info(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	command.Success(ctx, "信息获取成功", gin.H{"user": dto.ToUserDto(userInfo.(model.User))})
}
