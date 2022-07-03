package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/utils"
)

type UserHandler struct{}

// @Summary      注册
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        Register body request.Register true "Create User"
// @Success 	 200  {object}  string
// @Router       /user/register [post]
func (uh *UserHandler) Register(ctx *gin.Context) {
	var r request.Register

	_ = ctx.ShouldBindJSON(&r)
	// 参数校验
	if err := utils.Verify(r); err != nil {
		global.LOG.Errorf("parame bind err: %v", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if r.Password != r.CheckPwd {
		command.Failed(ctx, http.StatusBadRequest, "两次密码不相符")
		return
	}

	user, err := userService.Register(r)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "注册成功", user)
}

// @Summary      登录
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        Login body request.Login true "Login"
// @Success 	 200  {object}  string {data=token}
// @Router       /user/login [post]
func (uh *UserHandler) Login(ctx *gin.Context) {
	// 绑定表单参数
	var l request.Login

	_ = ctx.ShouldBindJSON(&l)
	// 参数校验
	if err := utils.Verify(l); err != nil {
		global.LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	// 验证码校验
	// if !store.Verify(l.CaptchaId, l.Captcha, true) {
	// 	command.Failed(ctx, http.StatusUnauthorized, "验证码不正确或已失效")
	// 	return
	// }

	user, err := userService.Login(l)
	if err != nil {
		command.Failed(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	uh.createToken(ctx, *user)

}

// 登录后签发token
func (uh *UserHandler) createToken(ctx *gin.Context, user model.User) {
	bc := utils.BaseClaims{
		Id:       user.Id,
		RoleId:   user.RoleId,
		UUID:     user.UUID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
	token, err := utils.CreateToken(bc)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "获取身份认证失败")
		return
	}
	command.Success(ctx, "登录成功", gin.H{"token": token})
}

// @Summary      获取用户信息
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  model.User {data=model.User}
// @Security 	 X-Token
// @Router       /user/info [get]
func (uh *UserHandler) UserInfo(ctx *gin.Context) {
	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	user, err := userService.GetUserInfo(id, uuid)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}
	command.Success(ctx, "获取成功", gin.H{"user": user})
}

// @Summary      登出
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  string
// @Security 	 X-Token
// @Router       /user/logout [get]
func (uh *UserHandler) Logout(ctx *gin.Context) {
	command.Success(ctx, "登出成功", nil)
}

// @Summary      修改用户名
// @Tags         System/User
// @Accept       mpfd
// @Produce      json
// @Param        nickname formData string true  "update nickname"
// @Success 	 200  {object}  string
// @Security 	 X-Token
// @Router       /user/name [put]
func (uh *UserHandler) UpdateNickname(ctx *gin.Context) {
	nickname := ctx.PostForm("nickname")
	if nickname == "" {
		global.LOG.Error("nickname cannot be empty")
		command.Failed(ctx, http.StatusBadRequest, "nickname cannot be empty")
		return
	}

	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	if err := userService.UpdateNickname(nickname, id, uuid); err != nil {
		global.LOG.Error(err)
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
		global.LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if u.OldPassword == u.LastPassword {
		command.Failed(ctx, http.StatusBadRequest, "新旧密码不能相同")
		return
	}

	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	if err := userService.UpdatePwd(u, id, uuid); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "修改成功", nil)
}

// @Summary      找回密码
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        ForgetPwd body request.ForgetPwd  true  "update user"
// @Success 	 200  {object}  string
// @Router       /user/forgetPwd [post]
func (uh *UserHandler) ForgetPwd(ctx *gin.Context) {
	var f request.ForgetPwd
	_ = ctx.ShouldBindJSON(&f)

	// 参数校验
	if err := utils.Verify(f); err != nil {
		global.LOG.Errorf("parame bind err:", err)
		command.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := userService.ForgetPwd(f); err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "修改成功", nil)
}

// @Summary      修改头像
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Param        file formData file true  "update user"
// @Success 	 200  {object}  string
// @Security 	 X-Token
// @Router       /user/avatar [put]
func (uh *UserHandler) UpdateAvatar(ctx *gin.Context) {

	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}

	url, err := utils.UploadFile(file, fileHeader.Size)
	if err != nil {
		global.LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "服务异常")
		return
	}

	uuid := utils.GetUserUUID(ctx)
	id := utils.GetUserId(ctx)
	if err := userService.UpdateAvatar(url, uuid, id); err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	command.Success(ctx, "修改成功", nil)
}

// @Summary      获取用户用户信息列表
// @Tags         System/User
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  response.PageList
// @Security 	 X-Token
// @Router       /user/list [get]
func (uh *UserHandler) QueryAll(ctx *gin.Context) {
	userList := userService.QueryAll()
	command.Success(ctx, "查询成功", gin.H{"userList": userList})
}
