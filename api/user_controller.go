package api

import (
	"Blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createUser(ctx *gin.Context) {

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogionResponse struct {
	Token string `json:"token"`
	// User User	`json:"user"`
}

func (server *Server) login(ctx *gin.Context) {

	// 参数绑定
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}
	// 判断用户是否存在
	// user,err := GetUser(req.Usernamee)
	// 校验密码
	err := utils.Decrypt(req.Password, "")
	if err != nil {
		ctx.SecureJSON(http.StatusUnauthorized, err.Error())
		return
	}

	// 颁发token
	// server.maker.CreateToken(user.Username,time.mintue)

	// 返回
}
