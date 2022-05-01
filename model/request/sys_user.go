package request

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// Captcha   string `json:"captcha"`
	// CaptchaId string `json:"captchaId"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// Captcha   string `json:"captcha"`
	// CaptchaId string `json:"captchaId"`
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

type UpdateAvatar struct {
	Avatar string `json:"avater"`
}

type Leave struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
