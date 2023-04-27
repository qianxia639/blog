package errors

import "errors"

var (
	ParamErr          = errors.New("参数错误")
	ServerErr         = errors.New("服务异常")
	NotExistsUserErr  = errors.New("用户不存在")
	NicknameExistsErr = errors.New("用户名重复")
	PasswordErr       = errors.New("密码错误")
	AccountLockedErr  = errors.New("账户已锁定,请稍后在试")
	UsernameErr       = errors.New("用户名错误")
	NotExistsBlogErr  = errors.New("文章不存在")
	TitleExistsErr    = errors.New("标题已存在,请更换标题")
	UnauthorizedError = errors.New("身份认证失败")
)
