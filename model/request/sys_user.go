package request

type Register struct {
	Username string
	Password string
	CheckPwd string
}

type Login struct {
	Username string
	Password string
	// Captcha   string
	// CaptchaId string
}

type UpdatePwd struct {
	OldPassword  string
	LastPassword string
}

type ForgetPwd struct {
	Password string
}
