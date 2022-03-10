package command

import "errors"

var (
	FindFailed  = errors.New("查询失败").Error()
	FindSuccess = errors.New("查询成功").Error()

	OperationFailed  = errors.New("操作失败").Error()
	OperationSuccess = errors.New("操作成功").Error()
)
