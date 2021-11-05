package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaileshhb/restapi/model/general"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Token -> ", r.Header)
		// var jwtKey = os.Getenv("ACCESS_SECRET")
		var jwtKey = []byte("98hbun98h")
		fmt.Println("-------- jwtKey ->", jwtKey)

		if r.Header.Get("Token") != "" {

			claims := &general.Claim{}

			token, err := jwt.ParseWithClaims(r.Header.Get("Token"), claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return jwtKey, nil
			})

			fmt.Println("-------- token.Valid ->", token.Valid)

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
