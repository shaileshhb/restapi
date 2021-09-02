package auth

import (
	"github.com/satori/uuid"
)

type AuthInterface interface {
	CreateAuth(uint64, *TokenDetails) error
	FetchAuth(string) (uint64, error)
	DeleteRefresh(string) error
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
	TokenUUID    uuid.UUID
	RefreshUUID  uuid.UUID
}

type AccessDetails struct {
	TokenUuid string
	UserId    uint64
}

// func CreateAuth(userID uuid.UUID, tokenDetails *TokenDetails) error {

// 	accessTokenExpiry := time.Unix(tokenDetails.AtExpires, 0) //converting Unix to UTC(to Time object)
// 	refreshTokenExpiry := time.Unix(tokenDetails.RtExpires, 0)
// 	now := time.Now()

// 	tokenDetails.TokenUUID = userID
// 	tokenDetails.AtExpires = now.Add(500 * time.Hour).Unix()
// 	tokenDetails.RtExpires = now.Add(500 * time.Hour).Unix()

// 	return nil

// }
