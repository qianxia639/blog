package system

import (
	"net/http"
	"sync"

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
	wg          sync.WaitGroup
}

// @Summary      注册
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        Register body request.Register  true  "Create User"
// @Success 	 200  {object}  string
// @Router       /user/register [post]
func (uh *UserHandler) Register(ctx *gin.Context) {
	var r request.Register

	_ = ctx.ShouldBindJSON(&r)
	if err := utils.Verify(r); err != nil {
		global.QX_LOG.Errorf("parame bind err: %v", err)
		command.RFailed(ctx, http.StatusBadRequest, err.Error())
	}

	_, err := uh.userService.Register(r)

	if err != nil {
		command.RFailed(ctx, http.StatusInternalServerError, err.Error())
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
// @Param        Login body request.Login  true  "Login"
// @Success 	 200  {object}  string {data=token}
// @Router       /user/login [post]
func (uh *UserHandler) Login(ctx *gin.Context) {
	// 绑定表单参数
	var l request.Login

	_ = ctx.ShouldBindJSON(&l)

	if err := utils.Verify(&l); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.RFailed(ctx, http.StatusBadRequest, err.Error())
	}

	user, err := uh.userService.Login(l)
	if err != nil {
		command.RFailed(ctx, http.StatusUnauthorized, err.Error())
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
	c := ctx.Copy()
	bc := utils.BaseClaims{
		Id:       user.Id,
		UUID:     user.UUID,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
	if token, err := utils.CreateToken(bc); err != nil {
		global.QX_LOG.Error("token生成失败!", err)
		command.RFailed(c, http.StatusInternalServerError, "获取token失败")
	} else {
		command.Success(c, "登录成功", gin.H{"token": token})
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
		command.RFailed(ctx, http.StatusInternalServerError, "获取失败")
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
// @Router       /user/updateName [put]
func (uh *UserHandler) UpdateUsername(ctx *gin.Context) {

	var u request.UpdateUsername
	_ = ctx.ShouldBindJSON(&u)

	if err := utils.Verify(&u); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.RFailed(ctx, http.StatusBadRequest, err.Error())
	}

	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	if err := uh.userService.UpdateUsername(u, id, uuid); err != nil {
		command.RFailed(ctx, http.StatusInternalServerError, err.Error())
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
// @Router       /user/updatePwd [put]
func (uh *UserHandler) UpdatePwd(ctx *gin.Context) {
	var u request.UpdatePwd

	_ = ctx.ShouldBindJSON(&u)
	if err := utils.Verify(&u); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.RFailed(ctx, http.StatusBadRequest, err.Error())
	}

	// todo 发送邮件并进行校验
	// uh.wg.Add(1)
	// go uh.sendMail(ctx, u.Email)
	// if exists, err := utils.VerifyMail(u.Email); err != nil && !exists {
	// 	global.QX_LOG.Errorf("verify email code err:", err)
	// 	fmt.Printf("err: %v\n", err)
	// 	fmt.Printf("exists: %v\n", exists)
	// 	command.RFailed(ctx, http.StatusInternalServerError, "验证码错误")
	// }
	// uh.wg.Wait()
	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	if err := uh.userService.UpdatePwd(u, id, uuid); err != nil {
		command.RFailed(ctx, http.StatusInternalServerError, err.Error())
	}
	command.Success(ctx, "修改成功", nil)
}

func (uh *UserHandler) sendMail(ctx *gin.Context, email string) {
	// defer uh.wg.Done()
	cc := ctx.Copy()
	if err := utils.SendMail(email); err != nil {
		global.QX_LOG.Errorf("send mail err: %v", err)
		command.RFailed(cc, 500, "邮件发送失败")
	}
}

// @Summary      修改头像
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        UpdateAvatar body request.UpdateAvatar  true  "update user"
// @Success 	 200  {object}  string
// @Security 	 X-Token
// @Router       /user/updateAvatar [put]
func (uh *UserHandler) UpdateAvatar(ctx *gin.Context) {
	var u request.UpdateAvatar

	_ = ctx.ShouldBindJSON(&u)
	if err := utils.Verify(&u); err != nil {
		global.QX_LOG.Errorf("parame bind err:", err)
		command.RFailed(ctx, http.StatusBadRequest, err.Error())
	}

	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	if err := uh.userService.UpdateAvatar(u, id, uuid); err != nil {
		command.RFailed(ctx, http.StatusInternalServerError, err.Error())
	}
	command.Success(ctx, "修改成功", nil)
}
