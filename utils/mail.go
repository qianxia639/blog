package utils

import (
	"github.com/qianxia/blog/global"
	"gopkg.in/gomail.v2"
)

func SendMail(content string, to ...string) {

	// 创建一条消息
	msg := gomail.NewMessage()
	msg.SetHeader("From", global.RY_YAML_CONFIG.Mail.Username) // 发件人
	msg.SetHeader("To", to...)                                 // 收件人
	msg.SetHeader("Subject", "欢迎注册")                           // 主题
	msg.SetBody("text/html", content)                          // 邮件主体
	d := gomail.NewDialer(global.RY_YAML_CONFIG.Mail.Host,
		global.RY_YAML_CONFIG.Mail.Port,
		global.RY_YAML_CONFIG.Mail.Username,
		global.RY_YAML_CONFIG.Mail.Password)
	if err := d.DialAndSend(msg); err != nil {
		global.RY_LOG.Errorf("%s", err.Error())
	}
	msg.Reset()
}
