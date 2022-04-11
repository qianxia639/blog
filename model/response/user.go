package response

import (
	"github.com/qianxia/blog/model"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id       uint64    `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
}

func ToUser(user model.User) *User {
	return &User{
		Id:       user.Id,
		UUID:     user.UUID,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
}
