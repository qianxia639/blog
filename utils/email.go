package utils

import (
	"net/smtp"
	"net/textproto"
	"time"

	"github.com/jordan-wright/email"
	"github.com/qianxia/blog/global"
)

func SendMail(to ...string) error {
	var code = CreateRandom()
	e := &email.Email{
		To:      to,
		From:    global.QX_CONFIG.Email.Username,
		Subject: "验证码",
		HTML:    []byte(`<p>验证码为: ` + code + `,该验证码将在5分钟后失效</p>`),
		Headers: textproto.MIMEHeader{},
	}

	err := SetTtlCache(to[0], code, 5*time.Minute)
	if err != nil {
		return err
	}
	return e.Send(global.QX_CONFIG.Email.Addr, smtp.PlainAuth("", global.QX_CONFIG.Email.Username, global.QX_CONFIG.Email.Password, global.QX_CONFIG.Email.Host))
}
