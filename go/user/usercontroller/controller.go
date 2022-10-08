package usercontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/model/user"
	"github.com/shaileshhb/restapi/security/auth"
	service "github.com/shaileshhb/restapi/user/userservice"
	"github.com/shaileshhb/restapi/web"
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

func (c *Controller) RegisterUserRoutes(getRouter, middlewareRouter *mux.Router) {

	getRouter.HandleFunc("/students/login", c.UserLogin).Methods("POST")
	getRouter.HandleFunc("/students/register", c.Register).Methods("POST")

}

// Register will register the user.
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" =========================== REGISTER =========================== ")
	var user = &user.User{}
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
	web.RespondJSON(w, http.StatusCreated, user.ID)
}

func (c *Controller) UserLogin(w http.ResponseWriter, r *http.Request) {

	var loginUser = &user.User{}
	var validateUser = &user.User{}
	var err error

	log.Println(" ---------------------- Inside userlogin ---------------------- ")

	userDetails, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(userDetails, loginUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Service.Get(validateUser, loginUser.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Println(" loginusername -> ", loginUser.Username)
	fmt.Println(" validateusername -> ", validateUser.Username)

	if validateUser.Username != loginUser.Username {
		http.Error(w, "Login failed for username!", http.StatusUnauthorized)
		log.Println("Username or password is invalid")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(validateUser.Password), []byte(loginUser.Password)); err != nil {
		http.Error(w, "Username or password is invalid", http.StatusUnauthorized)
		log.Println("Username or password is invalid password", err)
		return
	}

	// var tokenDetails = general.TokenDetails{}

	// err = auth.CreateToken(validateUser.Base.ID, &tokenDetails)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// log.Println(tokenDetails)

	// log.Println(string(response))
	// response, error := json.Marshal(tokenDetails)
	// if error != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(error.Error()))
	// 	return
	// }

	tokenDetails, err := auth.GenerateToken(loginUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(tokenDetails))
}
