package middleware

import (
	"net/http"
	"strings"
	"tublessin/api_gateway/utils/token"

	log "github.com/sirupsen/logrus"
)

// Validate Token from cookies
func TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) <= 1 {
			http.Error(w, "Token Expired", http.StatusUnauthorized)
			return
		}
		reqToken = splitToken[1]

		// getUser, _ := r.Cookie("user")
		// getToken, err := r.Cookie("token")
		// if err != nil {
		// 	http.Error(w, "Token Expired", http.StatusUnauthorized)
		// 	return
		// }

		validity, userName, id, _ := token.VerifyToken(reqToken)
		log.Println(id, userName+" accessing "+r.RequestURI)
		if validity == true {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid Sessions", http.StatusUnauthorized)
		}
	})
}
