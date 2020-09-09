package middleware

import (
	"log"
	"net/http"
	"tublessin/api_gateway/utils/token"
)

// Validate Token from cookies
func TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		return // Token Checking tidak di aktifkan dulu, masih tahap development

		getUser, _ := r.Cookie("user")
		getToken, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Token Expired", http.StatusUnauthorized)
			return
		}

		validity, userName, id, _ := token.VerifyToken(getToken.Value)
		log.Println(id, userName+" accessing "+r.RequestURI)
		if validity == true && userName == getUser.Value {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid Sessions", http.StatusUnauthorized)
		}
	})
}
