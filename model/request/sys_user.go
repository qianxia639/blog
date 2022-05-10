package request

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	emailRegexp    = `\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}`
	passwordRegexp = `[0-9a-zA-Z]{6,15}`
)

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// Captcha   string `json:"captcha"`
	// CaptchaId string `json:"captchaId"`
}

type Login struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required, validation.Match(regexp.MustCompile(passwordRegexp))),
		validation.Field(&l.Captcha, validation.Required),
		validation.Field(&l.CaptchaId, validation.Required))
}

type UpdateUsername struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (u UpdateUsername) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Username, validation.Required))
}

type UpdatePwd struct {
	Email        string `json:"email"`
	OldPassword  string `json:"old_password"`
	LastPassword string `json:"last_password"`
}

func (u UpdatePwd) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.OldPassword, validation.Required, validation.Match(regexp.MustCompile(passwordRegexp))),
		validation.Field(&u.LastPassword, validation.Required, validation.Match(regexp.MustCompile(passwordRegexp))))
}

type ForgetPwd struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateAvatar struct {
	Avatar string `json:"avater"`
}

type UpdateEmail struct {
	OldEmail  string `json:"old_email"`
	LastEmail string `json:"last_email"`
	Code      string `json:"code"`
}

type VerifyMail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type Leave struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
