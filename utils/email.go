package utils

import (
	"context"
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

	if err := setCache(to[0], code); err != nil {
		return err
	} else {
		return e.Send(global.QX_CONFIG.Email.Addr, smtp.PlainAuth("", global.QX_CONFIG.Email.Username, global.QX_CONFIG.Email.Password, global.QX_CONFIG.Email.Host))
	}
}

func VerifyMail(to, code string) (bool, error) {
	if res, err := GetCache(to); err != nil {
		return false, err
	} else if res != code {
		return false, err
	} else {
		return true, nil
	}
}

func setCache(to, code string) error {
	return global.QX_REDIS.Set(context.Background(), to, code, time.Minute*5).Err()
}

func GetCache(to string) (string, error) {
	return global.QX_REDIS.Get(context.Background(), to).Result()
}

func delCache(to string) error {
	return global.QX_REDIS.Del(context.Background(), to).Err()
}
