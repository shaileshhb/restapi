package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/shaileshhb/restapi/model"
	service "github.com/shaileshhb/restapi/user/user-service"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	Service *service.UserService
}

func NewController(service *service.UserService) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) RegisterUserRoutes(router *mux.Router) {

	router.HandleFunc("/students/login", c.UserLogin).Methods("POST")
	router.HandleFunc("/students/register", c.Register).Methods("POST")

}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {

	var user = &model.User{}
	var err error

	userDetails, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Response could not be read"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(userDetails, user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password = string(hashPassword)

	err = c.Service.Add(user)
	if err != nil {
		log.Println(err)
		// w.Write([]byte("Error while adding user, " + err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("User ID -> ", user.ID)
	// tokenString, err := c.generateJWT(user.ID, w)
	// if err != nil {
	// 	w.Write([]byte("Token string failed"))
	// 	log.Println("Token string failed")
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// w.Write([]byte(tokenString))
	// log.Println("User successfully added", tokenString)

}

func (c *Controller) UserLogin(w http.ResponseWriter, r *http.Request) {

	var user = &model.User{}
	var validateUser = &model.User{}
	var err error

	log.Println("Inside userlogin")

	userDetails, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(userDetails, user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Service.Get(validateUser, user.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		// w.WriteHeader(http.StatusBadRequest)
		return
	}
	if validateUser.Username != user.Username {
		// w.Write([]byte("Username or password is invalid"))
		http.Error(w, "Login failed for username!", http.StatusUnauthorized)
		log.Println("Username or password is invalid")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(validateUser.Password), []byte(user.Password)); err != nil {
		// w.Write([]byte("Failed"))
		http.Error(w, "Username or password is invalid", http.StatusUnauthorized)
		log.Println("Username or password is invalid password", err)
		return
	}

	// var tokenResponse = model.TokenResponse{}

	tokenResponse, err := c.generateJWT(validateUser.ID, w)
	if err != nil {
		// w.Write([]byte("Token string failed"))
		log.Println("Token string failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(tokenResponse)

	// response, err := json.Marshal(tokenResponse)
	// if err != nil {
	// 	log.Println("Token string failed")
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// log.Println(string(response))
	w.Write([]byte(tokenResponse))

}

func (c *Controller) generateJWT(userID uuid.UUID, w http.ResponseWriter) (string, error) {

	// secret key
	var jwtKey = []byte("some_secret_key")

	expirationTime := time.Now().Add(2 * time.Minute)

	// Creating JWT Claim which includes username and claims
	claims := &model.Claim{
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

// func (c *Controller) generateJWT(userID uuid.UUID, tokens *model.TokenResponse, w http.ResponseWriter) error {

// 	// secret key
// 	var jwtKey = []byte("some_secret_key")

// 	expirationTime := time.Now().Add(2 * time.Minute)

// 	// Creating JWT Claim which includes username and claims
// 	claims := &model.Claim{
// 		ID: userID,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	// Access Token
// 	// token having algo form signing method and the claim
// 	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	accessTokenString, err := userToken.SignedString(jwtKey)
// 	if err != nil {
// 		// w.Write([]byte("Failed"))
// 		// http.Error(w, err.Error(), http.StatusBadRequest)
// 		// log.Println("Username or password is invalid")
// 		return err
// 	}

// 	// Refresh Token
// 	refreshToken := jwt.New(jwt.SigningMethodHS256)
// 	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
// 	refreshTokenClaims["sub"] = 1
// 	refreshTokenClaims["exp"] = time.Now().Add(2 * time.Minute)

// 	refreshTokenString, err := refreshToken.SignedString(jwtKey)
// 	if err != nil {
// 		return err
// 	}

// 	tokens.AccessToken = accessTokenString
// 	tokens.RefreshToken = refreshTokenString

// 	return nil

// }
