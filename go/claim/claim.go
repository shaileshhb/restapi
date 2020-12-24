package claim

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type Claim struct {
	ID uuid.UUID `json:"username"`
	jwt.StandardClaims
}
