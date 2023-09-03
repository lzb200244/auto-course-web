package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-template/global"
	"strings"
	"time"
)

var secret = []byte(global.Config.Jwt.SECRET)

type Claims struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// GenerateToken 签发用户Token
func GenerateToken(id uint, username, email string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * global.Config.Jwt.Expire)
	claims := Claims{
		Id:        id,
		Username:  username,
		Email:     email,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

// ParseToken 验证用户token
func ParseToken(tokenHead string) (*Claims, error) {
	if !strings.HasPrefix(tokenHead, "Bearer ") {
		return nil, errors.New("无效的token")
	}
	var token = strings.SplitN(tokenHead, " ", 2)[1]
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
