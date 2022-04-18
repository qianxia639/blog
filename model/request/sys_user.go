package request

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	CheckPwd string `json:"checkPwd"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Leave struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
