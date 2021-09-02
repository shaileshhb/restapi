package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/uuid"
	"github.com/shaileshhb/restapi/model/general"
)

func CreateToken(userID uuid.UUID, tokenDetails *general.TokenDetails) error {

	tokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenDetails.TokenUUID = uuid.NewV4()

	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUUID = userID

	err := createAccessToken(userID, tokenDetails)
	if err != nil {
		return err
	}

	err = createRefreshToken(userID, tokenDetails)
	if err != nil {
		return err
	}

	return nil
}

//Creating Access Token
func createAccessToken(userID uuid.UUID, tokenDetails *general.TokenDetails) error {

	var err error
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["access_uuid"] = tokenDetails.TokenUUID
	accessTokenClaims["user_id"] = userID
	accessTokenClaims["exp"] = tokenDetails.AtExpires

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	tokenDetails.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return err
	}
	return nil
}

//Creating Refresh Token
func createRefreshToken(userID uuid.UUID, tokenDetails *general.TokenDetails) error {

	var err error
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["refresh_uuid"] = tokenDetails.RefreshUUID
	refreshTokenClaims["user_id"] = userID
	refreshTokenClaims["exp"] = tokenDetails.RtExpires

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenDetails.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return err
	}

	return nil
}
