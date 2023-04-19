package api

import (
	db "Blog/db/sqlc"
	"Blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string  `json:"token"`
	User  db.User `json:"user"`
}

func (server *Server) login(ctx *gin.Context) {

	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == ErrNoRows {
			ctx.SecureJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = utils.Decrypt(req.Password, user.Password)
	if err != nil {
		ctx.SecureJSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := server.maker.CreateToken(user.Username, server.conf.Token.AccessTokenDuration)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp := loginResponse{
		Token: token,
		User: db.User{
			ID:           user.ID,
			Username:     user.Username,
			Nickname:     user.Nickname,
			Email:        utils.DesnsitizeEmail(user.Email),
			Avatar:       user.Avatar,
			RegisterTime: user.RegisterTime,
		},
	}

	ctx.JSON(http.StatusOK, resp)
}
