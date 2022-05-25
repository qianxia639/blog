package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
)

type CaptchaHandler struct{}

var store = base64Captcha.DefaultMemStore

// @Summary 生成验证码
// @Tags System/Captcha
// @accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /system/captcha [post]
func (ch *CaptchaHandler) Captcha(ctx *gin.Context) {
	// 配置生成字符验证码
	driver := base64Captcha.NewDriverDigit(
		global.QX_CONFIG.Captcha.Height,
		global.QX_CONFIG.Captcha.Width,
		global.QX_CONFIG.Captcha.Length,
		global.QX_CONFIG.Captcha.MaxSkew,
		global.QX_CONFIG.Captcha.DotCount)
	captcha := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := captcha.Generate(); err != nil {
		global.QX_LOG.Error("验证码获取失败!", err)
		command.Failed(ctx, 500, "验证码获取失败")
		return
	} else {
		command.Success(ctx, "验证码获取成功", gin.H{
			"captchaId":   id,
			"captchaData": b64s,
		})
	}
}
