package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/uuid"
)

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {

	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Token")
	log.Println("Token -> ", token)
	return token
}

func (t *Token) CreateToken(userID uuid.UUID) (*TokenDetails, error) {

	tokenDetails := &TokenDetails{}
	tokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenDetails.TokenUUID = uuid.NewV4()

	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUUID = userID

	err := t.createAccessToken(userID, tokenDetails)
	if err != nil {
		return nil, err
	}

	err = t.createRefreshToken(userID, tokenDetails)
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}

//Creating Access Token
func (t *Token) createAccessToken(userID uuid.UUID, tokenDetails *TokenDetails) error {

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
func (t *Token) createRefreshToken(userID uuid.UUID, tokenDetails *TokenDetails) error {

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

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (t *Token) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	fmt.Println(" ------------------------ WE ENTERED METADATA ------------------------ ")
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			TokenUuid: accessUUID,
			UserId:    userId,
		}, nil
	}
	return nil, err
}
