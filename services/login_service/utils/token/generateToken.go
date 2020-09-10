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

	// rdb := config.NewRedisConnection()
	// err = rdb.Set(context.Background(), userId+"-"+username, tokenKey, 24*time.Hour).Err()
	// if err != nil {
	// 	log.Println("Failed inserting data to Redis", err)
	// }

	return tokenKey
}
