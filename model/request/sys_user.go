package request

type Register struct {
	Username string
	Password string
	CheckPwd string
}

type Login struct {
	Username  string
	Password  string
	Captcha   string
	CaptchaId string
}

type UpdateUsername struct {
	Email    string
	Username string
}

type UpdatePwd struct {
	Signer       string
	OldPassword  string
	LastPassword string
}

type ForgetPwd struct {
	Signer   string
	Password string
}
