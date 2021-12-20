package utils

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
	"gopkg.in/yaml.v2"
)

func SendMail(captcha string, to ...string) {

	dc, y := DeCode()
	yaml.Unmarshal(dc, &y)

	args := fmt.Sprintf("%s:%d", y.QQMail.Host, y.QQMail.Port)

	// 创建电子邮件
	e := email.NewEmail()
	// 设置发件人
	e.From = y.QQMail.Username
	// 设置收件人
	e.To = to
	// 设置主题
	e.Subject = "欢迎注册!!"
	// 设置邮件主体
	e.HTML = []byte(`您的本次验证码为【` + captcha + `】,请及时进行激活操作，非本人操作请无视`)
	// 发送邮件
	if err := e.Send(args, smtp.PlainAuth("", y.QQMail.Username, y.QQMail.Password, y.QQMail.Host)); err != nil {
		log.Printf("邮箱【：%s 】不存在", to[0])
		return
	}
}
