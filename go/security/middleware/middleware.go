package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaileshhb/restapi/model/general"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Token -> ", r.Header.Get("Token"))
		var jwtKey = os.Getenv("ACCESS_SECRET")

		if r.Header.Get("Token") != "" {

			claims := &general.Claim{}

			token, err := jwt.ParseWithClaims(r.Header.Get("Token"), claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return jwtKey, nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					http.Error(w, "User Not Authorized", http.StatusUnauthorized)
					return
				}
				http.Error(w, "User Not Authorized", http.StatusBadRequest)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			}
		} else {
			http.Error(w, "User Not Authorized", http.StatusUnauthorized)
		}
	})
}
