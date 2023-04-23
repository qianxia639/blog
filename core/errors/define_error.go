package errors

import "errors"

var (
	ParamErr         = errors.New("参数错误")
	ServerErr        = errors.New("服务异常")
	NotExistsUserErr = errors.New("用户不存在")
	PasswordErr      = errors.New("密码错误")
)
