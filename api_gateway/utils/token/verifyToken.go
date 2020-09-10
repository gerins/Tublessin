package token

import (
	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(tokenString string) (bool, string, string, error) {
	mySigningKey := []byte("admin")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	var id string
	var userName string
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		id = claims.Id
		userName = claims.Audience
	} else {
		log.Println(`error => `, err)
	}

	return token.Valid, userName, id, err
}
