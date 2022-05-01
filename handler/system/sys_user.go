package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/model/response"
	"github.com/qianxia/blog/service/system"
	"github.com/qianxia/blog/utils"
)

type UserHandler struct {
	userService system.UserService
}

// @Summary      注册
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        user body request.Register  true  "Create User"
// @Success 	 200  {object}  string
// @Router       /user/register [post]
func (uh *UserHandler) Register(ctx *gin.Context) {
	var r request.Register

	_ = ctx.ShouldBindJSON(&r)
	if err := utils.Verify(r); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	_, err := uh.userService.Register(r)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "注册成功", nil)

	// if store.Verify(r.CaptchaId, r.Captcha, true) {
	// 	_, err := uh.userService.Register(r)

	// 	if err != nil {
	// 		command.Failed(ctx, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// 	command.Success(ctx, "注册成功", nil)
	// } else {
	// 	command.Failed(ctx, http.StatusUnauthorized, "验证码错误")
	// 	return
	// }

}

// @Summary      登录
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        user body request.Login  true  "Login"
// @Success 	 200  {object}  string {data=token}
// @Router       /user/login [post]
func (uh *UserHandler) Login(ctx *gin.Context) {
	// 绑定表单参数
	var l request.Login

	_ = ctx.ShouldBindJSON(&l)

	if err := utils.Verify(&l); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := uh.userService.Login(l)
	if err != nil {
		command.Failed(ctx, http.StatusUnauthorized, err.Error())
		return
	} else {
		uh.createToken(ctx, *user)
	}

	// if store.Verify(l.CaptchaId, l.Captcha, true) {
	// 	user, err := uh.userService.Login(l)
	// 	if err != nil {
	// 		command.Failed(ctx, http.StatusUnauthorized, err.Error())
	// 		return
	// 	} else {
	// 		uh.createToken(ctx, *user)
	// 	}
	// } else {
	// 	command.Failed(ctx, http.StatusUnauthorized, "验证码错误")
	// 	return
	// }

}

// 登录后签发token
func (uh *UserHandler) createToken(ctx *gin.Context, user model.User) {
	bc := utils.BaseClaims{
		Id:       user.Id,
		UUID:     user.UUID,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
	if token, err := utils.CreateToken(bc); err != nil {
		global.QX_LOG.Error("token生成失败!", err)
		command.Failed(ctx, http.StatusInternalServerError, "获取token失败")
		return
	} else {
		command.Success(ctx, "登录成功", gin.H{"token": token})
	}

}

// @Summary      获取用户信息
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  response.User {data=response.User}
// @Security 	 X-Token
// @Router       /user/info [get]
func (uh *UserHandler) Info(ctx *gin.Context) {
	uuid := utils.GetUserUUID(ctx)
	if user, err := uh.userService.GetUserInfo(uuid); err != nil {
		global.QX_LOG.Errorf("用户信息获取失败! - {%s}", err)
		command.Failed(ctx, http.StatusInternalServerError, "获取失败")
	} else {
		command.Success(ctx, "获取成功", gin.H{"user": response.ToUser(*user)})
	}
}

// @Summary      修改用户名
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        user body model.User  true  "update user"
// @Success 	 200  {object}  string
// @Security 	 X-Token
// @Router       /user/updateName [put]
func (uh *UserHandler) UpdateUsername(ctx *gin.Context) {

	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		command.Failed(ctx, http.StatusBadRequest, "缺少必要的参数")
		global.QX_LOG.Errorf("parame bind err:", err)
		return
	}
	user.Id = utils.GetUserId(ctx)
	if err := uh.userService.UpdateUsername(user); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "修改成功", nil)
}
