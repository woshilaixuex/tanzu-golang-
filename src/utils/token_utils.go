package utils

import (
	"github.com/dgrijalva/jwt-go"
	"goweb/models"
	"log"
	"time"
)

const key = "this is jwt`s key"

type GoClaims struct {
	ID       string
	Username string
	jwt.StandardClaims
}

func CreatToken(user models.User) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, GoClaims{
		ID:       user.ID,
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 24*60*60,
			Issuer:    "lyric",
		},
	})
	token, err := tokenClaims.SignedString(key)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return token, nil
}

// GetTokenUserID token获取id(返回错误未写)
func GetTokenUserID(token string) string {
	claims, err := jwt.ParseWithClaims(token, &GoClaims{}, func(jwt *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return ""
	}
	tokenClaims := claims.Claims.(*GoClaims)
	return tokenClaims.ID
}
