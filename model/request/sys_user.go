package request

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

type Email struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type UpdateUsername struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdatePwd struct {
	Email        string `json:"email"`
	OldPassword  string `json:"old_password"`
	LastPassword string `json:"last_password"`
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
