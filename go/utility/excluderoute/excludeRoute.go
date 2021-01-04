package excluderoute

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/model"
)

func Authorization(excludedRoutes []*mux.Route) func(http.Handler) http.Handler {
	// Cache the regex object of each route (obviously for performance purposes)

	var excludedRoutesRegexp []*regexp.Regexp
	rl := len(excludedRoutes)
	for i := 0; i < rl; i++ {
		r := excludedRoutes[i]
		// log.Println("Routes -> ", r)
		pathRegexp, _ := r.GetPathRegexp()
		regx, _ := regexp.Compile(pathRegexp)

		excludedRoutesRegexp = append(excludedRoutesRegexp, regx)
	}
	// log.Println("ExculdedRoutes -> ", excludedRoutesRegexp)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// log.Println("Inside validation")

			exclude := false
			requestMethod := r.Method

			// log.Println("Request Method -> ", requestMethod)

			for i := 0; i < rl; i++ {
				excludedRoute := excludedRoutes[i]
				methods, _ := excludedRoute.GetMethods()
				ml := len(methods)
				// log.Println("Route Method ->", methods, "lenght -> ", ml)

				methodMatched := false
				if ml < 1 {
					// log.Println("Making method matched true")
					methodMatched = true
				} else {
					for j := 0; j < ml; j++ {
						// log.Println("Methods[j] -> ", methods[j], "Request Method -> ", requestMethod)
						if methods[j] == requestMethod {
							methodMatched = true
							break
						}
					}
				}
				// log.Println("Matched ->", methodMatched)
				if methodMatched {
					uri := r.RequestURI
					// log.Println("Excluded Routes ->", excludedRoutesRegexp[i], "URI -> ", uri)
					if excludedRoutesRegexp[i].MatchString(uri) {
						exclude = true
						break
					}
				}
			}
			if !exclude {
				// validationUserToken(next)
				// log.Println("Token -> ", r.Header["Token"])
				var jwtKey = []byte("some_secret_key")

				if r.Header["Token"] != nil {

					claims := &model.Claim{}

					token, err := jwt.ParseWithClaims(r.Header["Token"][0], claims, func(token *jwt.Token) (interface{}, error) {
						if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
							return nil, fmt.Errorf("There was an error")
						}
						return jwtKey, nil
					})
					if err != nil {
						if err == jwt.ErrSignatureInvalid {
							http.Error(w, "User Not Authorized", http.StatusUnauthorized)
							// w.WriteHeader(http.StatusUnauthorized)
							return
						}
						http.Error(w, "User Not Authorized", http.StatusBadRequest)
						// w.WriteHeader(http.StatusBadRequest)
						return
					}

					// log.Println("Token->", *token)
					// log.Println("Claims->", time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))

					// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 60*time.Second {
					// 	refresh
					// }

					if token.Valid {
						next.ServeHTTP(w, r)
					}
				} else {
					http.Error(w, "User Not Authorized", http.StatusUnauthorized)
					// fmt.Fprintf(w, "Not Authorized")
				}
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
