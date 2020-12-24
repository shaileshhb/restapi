package claim

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
