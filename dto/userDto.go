package dto

import "github.com/qianxia/blog/model"

type UserDto struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id:       user.Id,
		Username: user.Username,
	}
}
