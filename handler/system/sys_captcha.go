package system

import (
	"image/color"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
)

type CaptchaHandler struct{}

var store = base64Captcha.DefaultMemStore

// @Summary 生成验证码
// @Tags Captcha
// @accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/captcha [post]
func (ch *CaptchaHandler) Captcha(ctx *gin.Context) {
	// 配置生成字符验证码
	driver := base64Captcha.NewDriverString(
		global.QX_CONFIG.Captcha.Height,
		global.QX_CONFIG.Captcha.Width,
		global.QX_CONFIG.Captcha.NoiseCount,
		global.QX_CONFIG.Captcha.ShowLineOptions,
		global.QX_CONFIG.Captcha.Length,
		global.QX_CONFIG.Captcha.Source,
		(*color.RGBA)(&global.QX_CONFIG.Captcha.Color), nil, nil)
	captcha := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := captcha.Generate(); err != nil {
		global.QX_LOG.Error("验证码获取失败!", err)
		command.RFailed(ctx, 500, "验证码获取失败")
	} else {
		command.Success(ctx, "验证码获取成功", gin.H{
			"captchaId": id,
			"base64":    b64s,
		})
	}
}
