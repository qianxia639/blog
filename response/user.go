package response

import "github.com/qianxia/blog/model"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func ToUser(user model.User) User {
	return User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
}
