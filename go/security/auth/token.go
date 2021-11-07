package auth

import (
	"fmt"
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

func GenerateToken(userID uuid.UUID) (string, error) {

	// secret key
	// var jwtKey = os.Getenv("ACCESS_SECRET")
	var jwtKey = []byte("98hbun98h")
	fmt.Println("-------- jwtKey ->", jwtKey)

	expirationTime := time.Now().Add(60 * time.Minute)

	// Creating JWT Claim which includes username and claims
	claims := &general.Claim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Access Token
	// token having algo form signing method and the claim
	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenString, err := userToken.SignedString(jwtKey)
	if err != nil {
		// w.Write([]byte("Failed"))
		// http.Error(w, err.Error(), http.StatusBadRequest)
		// log.Println("Username or password is invalid")
		return "", err
	}

	return accessTokenString, nil
}
