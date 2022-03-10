package response

import "github.com/qianxia/blog/model"

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func ToUser(user model.User) *User {
	return &User{
		Id:       user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
}
