package app

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
)

type IUserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Info(ctx *gin.Context)
	Logout(ctx *gin.Context)
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

	uh, err := u.Service.Register(user)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	command.Success(ctx, "注册成功", &uh)
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
	gob.Register(model.User{})
	if err := utils.SaveSession(ctx, "user", user); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// 生成token
	token := utils.CreateToken(user.Id)
	command.Success(ctx, "登录成功", gin.H{"token": token})
}

// 获取用户信息
func (u UserHandler) Info(ctx *gin.Context) {
	// userInfo, _ := ctx.Get("user")
	userInfo, _ := utils.GetSession(ctx, "user")
	// userMap := make(map[interface{}]interface{})
	// userMap["id"] = userInfo.(model.User).Id
	// userMap["username"] = userInfo.(model.User).Username
	// fmt.Println(userInfo.(model.User))
	command.Success(ctx, "信息获取成功", gin.H{
		"id":       userInfo.(model.User).Id,
		"username": userInfo.(model.User).Username,
	})
}

func (u UserHandler) Logout(ctx *gin.Context) {
	if err := utils.RemoveSession(ctx); err != nil {
		command.Failed(ctx, 500, err.Error())
	}
	command.Success(ctx, "成功", nil)
}
