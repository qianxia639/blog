package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/service/system"
	"github.com/qianxia/blog/utils"
)

type UserHandler struct {
	userService system.UserService
}

// 注册
func (uh UserHandler) Register(ctx *gin.Context) {
	var user model.User
	// 绑定表单数据
	if err := ctx.ShouldBindJSON(&user); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	_, err := uh.userService.Register(user)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "注册成功", nil)
}

// 登录
func (uh UserHandler) Login(ctx *gin.Context) {
	// 绑定表单参数
	var form model.User
	ctx.ShouldBindJSON(&form)
	user, err := uh.userService.Login(form)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		global.RY_LOG.Warn(err)
		return
	}
	// 生成token
	token := utils.CreateToken(user.Id)
	command.Success(ctx, "登录成功", gin.H{"token": token})
}

// 获取用户信息
func (uh UserHandler) Info(ctx *gin.Context) {
	userInfo := ctx.MustGet("user")
	userMap := make(map[string]interface{}, 1)
	userMap["id"] = userInfo.(model.User).Id
	userMap["username"] = userInfo.(model.User).Username
	userMap["email"] = userInfo.(model.User).Email
	userMap["avatar"] = userInfo.(model.User).Avatar
	command.Success(ctx, "信息获取成功", gin.H{"user": userMap})
}
