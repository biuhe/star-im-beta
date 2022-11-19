package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// TokenExpireDuration 过期时间 - 2天
const TokenExpireDuration = time.Hour * 24 * 2

var Secret = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")

// MySecret (盐) 用来解密
func MySecret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenToken 生成 jwt token
func GenToken(username string) (string, error) {
	var claims = MyClaims{
		username,
		jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			// 签发人
			Issuer: "heyongbiao",
		},
	}

	// 使用HS256算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(Secret)
	if err != nil {
		return "", fmt.Errorf("生成token失败:%v", err)
	}
	return signedToken, nil
}

// ParseToken 验证jwt token
func ParseToken(tokenStr string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, MySecret())

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
