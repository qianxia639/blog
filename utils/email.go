package utils

import (
	"context"
	"net/smtp"
	"net/textproto"
	"time"

	"github.com/jordan-wright/email"
	"github.com/qianxia/blog/global"
)

var code = CreateRandom()

func SendMail(to ...string) error {
	e := &email.Email{
		To:      to,
		From:    "",
		Subject: "验证码",
		HTML:    []byte(`<p>验证码为: ` + code + `</p>`),
		Headers: textproto.MIMEHeader{},
	}

	if err := setCache(to[0], code); err != nil {
		return err
	} else {
		return e.Send(global.QX_CONFIG.Email.Addr, smtp.PlainAuth("", "", "sxkfftyexmhpdcaa", global.QX_CONFIG.Email.Host))
	}
}

func VerifyMail(to string) (bool, error) {
	if res, err := getCache(to); err != nil {
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

func getCache(to string) (string, error) {
	return global.QX_REDIS.Get(Redis().Context(), to).Result()
}
