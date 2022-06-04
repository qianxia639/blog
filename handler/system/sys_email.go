package system

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/utils"
)

type EmailHandler struct{}

// @Summary      发送邮箱验证码
// @Tags         System/Email
// @Accept       json
// @Produce      json
// @Param        email query string  true  "send mail code"
// @Success 	 200  {object}  string
// @Router       /system/email 	[get]
func (eh *EmailHandler) SendMail(ctx *gin.Context) {

	email := ctx.Query("email")

	if err := utils.SendMail(email); err != nil {
		global.QX_LOG.Errorf("send mail code err: %v", err)
		command.Failed(ctx, http.StatusInternalServerError, "验证码发送失败")
		return
	}
	command.Success(ctx, "验证码发送成功", nil)
}

// @Summary      校验邮箱验证码
// @Tags         System/Email
// @Accept       json
// @Produce      json
// @Param        Email body   request.Email true "verify mail code"
// @Success 	 200  {object}  string
// @Router       /system/verifyMail 	[post]
func (eh *EmailHandler) VerifyMail(ctx *gin.Context) {

	var e request.Email

	_ = ctx.ShouldBindJSON(&e)

	code, err := global.QX_REDIS.Get(context.Background(), e.Email).Result()
	if err != nil || code != e.Code {
		global.QX_LOG.Errorf("verify mail code err: %v", err)
		command.Failed(ctx, http.StatusInternalServerError, "验证码不正确或已失效")
		return
	}
	command.Success(ctx, "校验成功", nil)
}
