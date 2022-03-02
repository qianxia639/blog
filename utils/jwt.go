package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/qianxia/blog/global"
)

type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// 生成token
func CreateToken(email string) string {
	claims := &CustomClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                         // 生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 过期时间
			Issuer:    "qianxia",                                              // 签发人
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
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return global.RY_JWT_Key, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, jwt.ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, jwt.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, jwt.ErrTokenNotValidYet
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("无法解析的token")
	} else {
		return nil, errors.New("无法解析的token")
	}
}
