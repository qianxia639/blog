package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	OK   bool        `json:"ok"`             // 请求是否成功
	Des  string      `json:"des,omitempty"`  // 特殊说明
	Data interface{} `json:"data,omitempty"` // 数据信息
	Err  *ErrorInfo  `json:"err,omitempty"`  //错误信息
}

type ResultList struct {
	PageNo   int32       `json:"page_no"`   // 页号码
	PageSize int32       `json:"page_size"` // 页大小
	Total    int64       `json:"total"`     // 数据总条数
	List     interface{} `json:"list"`      // 数组数据
}

type ErrorInfo struct {
	Msg string `json:"msg"` //错误信息
}

func OK(ctx *gin.Context, des string) {
	res := Result{OK: true, Des: des}
	ctx.SecureJSON(http.StatusOK, res)
}

func Obj(ctx *gin.Context, data interface{}) {
	res := Result{OK: true, Data: data}
	ctx.SecureJSON(http.StatusOK, res)
}

func ArrarPage(ctx *gin.Context, data []interface{}, total int64, pageNo, pageSize int32) {
	res := Result{
		OK: true,
		Data: ResultList{
			List:     data,
			PageNo:   pageNo,
			PageSize: pageSize,
			Total:    total,
		},
	}
	ctx.JSON(http.StatusOK, res)
}

func Error(ctx *gin.Context, statusCode int, err string) {
	res := Result{
		OK: false,
		Err: &ErrorInfo{
			Msg: err,
		},
	}
	ctx.SecureJSON(statusCode, res)
}

func ServerError(ctx *gin.Context, err string) {
	res := Result{
		OK: false,
		Err: &ErrorInfo{
			Msg: err,
		},
	}
	ctx.SecureJSON(http.StatusInternalServerError, res)
}

func UnauthorizedError(ctx *gin.Context, err string) {
	res := Result{
		OK: false,
		Err: &ErrorInfo{
			Msg: err,
		},
	}
	ctx.SecureJSON(http.StatusUnauthorized, res)
}

func BadRequestError(ctx *gin.Context, err string) {
	res := Result{
		OK: false,
		Err: &ErrorInfo{
			Msg: err,
		},
	}
	ctx.SecureJSON(http.StatusBadRequest, res)
}
