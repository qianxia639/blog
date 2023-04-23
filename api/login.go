package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (server *Server) login(ctx *gin.Context) {

	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error("Bind Param err: ", err)
		ctx.SecureJSON(http.StatusBadRequest, errors.ParamErr.Error())
		return
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == ErrNoRows {
			logs.Logs.Error("Not User err: ", err)
			ctx.SecureJSON(http.StatusNotFound, errors.NotExistsUserErr.Error())
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, errors.ServerErr.Error())
		return
	}

	err = utils.Decrypt(req.Password, user.Password)
	if err != nil {
		logs.Logs.Error("Decrypy Password err: ", err)
		ctx.SecureJSON(http.StatusUnauthorized, errors.PasswordErr.Error())
		return
	}

	token, err := server.maker.CreateToken(user.Username, server.conf.Token.AccessTokenDuration)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, errors.ServerErr.Error())
		return
	}

	// key := fmt.Sprintf("userInfo:%s", user.Username)
	// err = server.rdb.Set(ctx, key, &user, 24*time.Hour).Err()
	// if err != nil {
	// 	logs.Logs.Error("redis err: ", err)
	// 	ctx.JSON(http.StatusInsufficientStorage, errors.ServerErr.Error())
	// 	return
	// }

	ctx.JSON(http.StatusOK, loginResponse{Token: token})
}
