package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/qianxia/blog/global"
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
		},
	}
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenStr.SignedString(global.RY_JWT_Key)
	if err != nil {
		panic(err)
	}
	return token
}

// 解析token
func ParseJwt(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return global.RY_JWT_Key, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token格式不正确")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("无效的token")
			}
		}
	}
	if token != nil {
		if token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("无效的token")
}
