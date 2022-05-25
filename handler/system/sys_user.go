package system

import (
	"fmt"
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
// @Tags         System
// @Accept       json
// @Produce      json
// @Param        Register body request.Register  true  "Create User"
// @Success 	 200  {object}  string
// @Router       /system/register [post]
func (uh *UserHandler) Register(ctx *gin.Context) {
	var r request.Register

	_ = ctx.ShouldBindJSON(&r)
	if err := utils.Verify(r); err != nil {
		global.QX_LOG.Errorf("parame bind err: %v", err)
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
// @Tags         System
// @Accept       json
// @Produce      json
// @Param        Login body request.Login  true  "Login"
// @Success 	 200  {object}  string {data=token}
// @Router       /system/login [post]
func (uh *UserHandler) Login(ctx *gin.Context) {
	// 绑定表单参数
	var l request.Login

	_ = ctx.ShouldBindJSON(&l)

	if err := utils.Verify(l); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if store.Verify(l.CaptchaId, l.Captcha, true) {
		if user, err := uh.userService.Login(l); err != nil {
			command.Failed(ctx, http.StatusUnauthorized, err.Error())
			return
		} else {
			uh.createToken(ctx, *user)
		}
	} else {
		command.Failed(ctx, http.StatusUnauthorized, "验证码错误")
		return
	}

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

// @Summary      邮箱登录
// @Tags         System
// @Accept       json
// @Produce      json
// @Param        EmailLogin body request.EmailLogin  true  "EmailLogin"
// @Success 	 200  {object}  string {data=token}
// @Router       /system/emailLogin [post]
func (uh *UserHandler) EmailLogin(ctx *gin.Context) {
	var el request.EmailLogin

	_ = ctx.ShouldBindJSON(&el)

	if ok, err := utils.VerifyMail(el.Email, el.Code); err == nil && ok {
		if user, err := uh.userService.GetUser(el.Email); err != nil {
			global.QX_LOG.Error("get user err: ", err)
			command.Failed(ctx, http.StatusUnauthorized, "服务错误")
			return
		} else {
			uh.createToken(ctx, *user)
		}
	} else {
		global.QX_LOG.Error("verify mail err: ", err)
		command.Failed(ctx, http.StatusUnauthorized, "验证码错误")
		return
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
	id := utils.GetUserId(ctx)
	if user, err := uh.userService.GetUserInfo(id, uuid); err != nil {
		global.QX_LOG.Errorf("用户信息获取失败! - {%s}", err)
		command.Failed(ctx, http.StatusInternalServerError, "服务错误")
		return
	} else {
		command.Success(ctx, "获取成功", gin.H{"user": response.ToUser(*user)})
	}
}

// @Summary      修改用户名
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        UpdateUsername body request.UpdateUsername  true  "update user"
// @Success 	 200  {object}  string
// @Security 	 X-Token
// @Router       /user/name [put]
func (uh *UserHandler) UpdateUsername(ctx *gin.Context) {

	// var u request.UpdateUsername
	// _ = ctx.ShouldBindJSON(&u)

	username := ctx.PostForm("username")
	fmt.Printf("username: %v\n", username)

	if err := utils.Verify(username); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	if err := uh.userService.UpdateUsername(username, id, uuid); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "修改成功", nil)
}

// @Summary      修改密码
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        UpdatePwd body request.UpdatePwd  true  "update user"
// @Success 	 200  {object}  string
// @Security 	 X-Token
// @Router       /user/pwd [put]
func (uh *UserHandler) UpdatePwd(ctx *gin.Context) {
	var u request.UpdatePwd

	_ = ctx.ShouldBindJSON(&u)
	if err := utils.Verify(&u); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// todo 发送邮件并进行校验
	// go uh.sendMail(ctx, u.Email)
	// if exists, err := utils.VerifyMail(u.Email); err != nil && !exists {
	// 	global.QX_LOG.Errorf("verify email code err:", err)
	// 	fmt.Printf("err: %v\n", err)
	// 	fmt.Printf("exists: %v\n", exists)
	// 	command.RFailed(ctx, http.StatusInternalServerError, "验证码错误")
	// }
	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	if err := uh.userService.UpdatePwd(u, id, uuid); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "修改成功", nil)
}

func (uh *UserHandler) sendMail(ctx *gin.Context, email string) {
	// defer uh.wg.Done()
	cc := ctx.Copy()
	if err := utils.SendMail(email); err != nil {
		global.QX_LOG.Errorf("send mail err: %v", err)
		command.Failed(cc, 500, "邮件发送失败")
		return
	}
}

func (uh *UserHandler) ForgetPwd(ctx *gin.Context) {
	var f request.ForgetPwd
	_ = ctx.ShouldBindJSON(&f)

	if err := uh.userService.ForgetPwd(f); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, "")
		return
	}
	command.Success(ctx, "修改成功", nil)
}

// @Summary      修改头像
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        UpdateAvatar body request.UpdateAvatar  true  "update user"
// @Success 	 200  {object}  string
// @Security 	 X-Token
// @Router       /user/avatar [put]
func (uh *UserHandler) UpdateAvatar(ctx *gin.Context) {
	var u request.UpdateAvatar

	_ = ctx.ShouldBindJSON(&u)
	if err := utils.Verify(&u); err != nil {
		global.QX_LOG.Errorf("parame bind err: %v", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	if err := uh.userService.UpdateAvatar(u, id, uuid); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "修改成功", nil)
}

// @Summary      修改邮箱
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        UpdateEmail body request.UpdateEmail  true  "update user"
// @Success 	 200  {object}  string
// @Security 	 X-Token
// @Router       /user/email [put]
func (uh *UserHandler) UpdateEmail(ctx *gin.Context) {
	var u request.UpdateEmail

	if err := ctx.ShouldBindJSON(&u); err != nil {
		global.QX_LOG.Errorf("parame bind err: %v", err)
		command.Failed(ctx, http.StatusBadRequest, "参数错误")
		return
	}

	// uuid := utils.GetUserUUID(ctx)
	// id := utils.GetUserId(ctx)

	// if err := uh.userService.UpdateEmail(u, id, uuid); err != nil {
	// 	command.RFailed(ctx, http.StatusInternalServerError, err.Error())
	// }
	command.Success(ctx, "修改成功", nil)
}
