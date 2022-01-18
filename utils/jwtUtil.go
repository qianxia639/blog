package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jwt_key = "l_qian_xia_y_y"
)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

// 生成token
func CreateToken(id int64) string {
	exp := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "qianxia",
			Subject:   "user token",
		},
	}
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenStr.SignedString([]byte(jwt_key))
	if err != nil {
		panic(err)
	}
	return token
}

// 解析token
func ParseJwt(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwt_key), nil
	})
	return token, claims, err
}
