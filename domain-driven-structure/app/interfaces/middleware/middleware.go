package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaileshhb/restapi/app/infrastructure/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "User Not Authorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "User Not Authorized", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// CORSMiddleware not sure abt functionality.
func CORSMiddleware() {
}
