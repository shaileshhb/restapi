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
	"github.com/shaileshhb/restapi/claim"
	usermodel "github.com/shaileshhb/restapi/user/user-model"
	userservice "github.com/shaileshhb/restapi/user/user-service"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	Service *userservice.UserService
}

func NewController(service *userservice.UserService) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) RegisterUserRoutes(router *mux.Router) {

	router.HandleFunc("/students/login", c.UserLogin).Methods("POST")
	router.HandleFunc("/students/register", c.Register).Methods("POST")

}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {

	var user = &usermodel.User{}
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
	user.Password = string(hashPassword)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = c.Service.Add(user); err != nil {
		log.Println(err)
		// w.Write([]byte("Error while adding user, " + err.Error()))
	} else {
		// w.Write([]byte(user.ID.String()))
		tokenString, err := c.generateJWT(user.ID, w)
		if err != nil {
			w.Write([]byte("Token string failed"))
			log.Println("Token string failed")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte(tokenString))
		// log.Println("User successfully added", user.ID)
	}
}

func (c *Controller) UserLogin(w http.ResponseWriter, r *http.Request) {

	var user = &usermodel.User{}
	var validateUser = &usermodel.User{}
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

	if err = c.Service.Get(validateUser, user.Username); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		// w.WriteHeader(http.StatusBadRequest)
	} else {

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

		tokenString, err := c.generateJWT(validateUser.ID, w)
		if err != nil {
			w.Write([]byte("Token string failed"))
			log.Println("Token string failed")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(tokenString)
		w.Write([]byte(tokenString))

	}

}

func (c *Controller) generateJWT(userID uuid.UUID, w http.ResponseWriter) (string, error) {

	// secret key
	var jwtKey = []byte("some_secret_key")

	expirationTime := time.Now().Add(24 * time.Hour)

	// Creating JWT Claim which includes username and claims
	claims := &claim.Claim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// token having algo form signing method and the claim
	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := userToken.SignedString(jwtKey)
	if err != nil {
		// w.Write([]byte("Failed"))
		// http.Error(w, err.Error(), http.StatusBadRequest)
		// log.Println("Username or password is invalid")
		return tokenString, err
	}

	return tokenString, nil

	// Setting up the cookie for userToken
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "UserToken",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })

}
