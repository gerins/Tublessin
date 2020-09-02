package token

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username, userId string, duration int64) string {
	mySigningKey := []byte("admin")

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + duration,
		Audience:  username,
		Id:        userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenKey, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err)
	}

	return tokenKey
}
