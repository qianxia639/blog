package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("lyyBlog")

type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	Id       uint64
	Username string
	Avatar   string
}

// 生成token
func CreateToken(baseClaims BaseClaims) (string, error) {
	claims := &CustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                         // 生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 过期时间
			Issuer:    "qianxia",                                              // 签发人
		},
	}
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenStr.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

// 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
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
		return nil, jwt.ErrTokenUnverifiable
	} else {
		return nil, jwt.ErrTokenUnverifiable
	}
}
