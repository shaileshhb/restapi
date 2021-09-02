package general

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/uuid"
)

type Claim struct {
	ID uuid.UUID
	jwt.StandardClaims
}
