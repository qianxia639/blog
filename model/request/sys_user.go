package request

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	CheckPwd string `json:"checkPwd"`
	// Captcha   string `json:"captcha"`
	// CaptchaId string `json:"captchaId"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// Captcha   string `json:"captcha"`
	// CaptchaId string `json:"captchaId"`
}

type Leave struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
