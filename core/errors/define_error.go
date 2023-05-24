package errors

import (
	"database/sql"
)

var (
	NoRowsErr             = sql.ErrNoRows
	UniqueViolationErr    = "unique_violation"
	ForeignKyViolationErr = "foreign_key_violation"
)

var (
	ParamErr                  = NewWrapError("参数错误")
	ServerErr                 = NewWrapError("服务异常")
	InvalidSyntaxErr          = NewWrapError("无效语法")
	NotExistsUserErr          = NewWrapError("用户不存在")
	NicknameExistsErr         = NewWrapError("用户名重复")
	UsernameOrEmailEexistsErr = NewWrapError("用户名或邮箱已注册")
	PasswordErr               = NewWrapError("密码错误")
	AccountLockedErr          = NewWrapError("账户已锁定,请稍后在试")
	UsernameErr               = NewWrapError("用户名错误")
	NotExistsAtricleErr       = NewWrapError("文章不存在")
	TitleExistsErr            = NewWrapError("标题已存在,请更换标题")
	UnauthorizedErr           = NewWrapError("身份认证失败")
	FileContentTypeErr        = NewWrapError("文件类型错误")
	ExceedingLenggthErr       = NewWrapError("超出大小限制")
)
