package util

import (
	"dev-framework-go/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Uuid     string `json:"uuid"`
	jwt.StandardClaims
}

func GenerateToken(username, uuid string) string {
	claims := Claims{
		EncodeMD5(username),
		EncodeMD5(uuid),
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix()), //到期时间
			IssuedAt:  int64(time.Now().Unix()), //发布时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(conf.JWTTOKEN)
	if err != nil {
		return ""
	}
	return ss
}

func VerifyToken(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return conf.JWTTOKEN, nil
	})
	if err != nil {
		return false
	}
	return true
}
