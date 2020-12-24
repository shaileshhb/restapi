package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	usermodel "github.com/shaileshhb/restapi/user/usermodel"
	userservice "github.com/shaileshhb/restapi/user/userservice"
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

	if err = c.Service.Add(user); err != nil {
		log.Println(err)
		w.Write([]byte("Error while adding user, " + err.Error()))
	} else {
		w.Write([]byte(user.ID.String()))
		log.Println("User successfully added", user.ID)
	}
}

func (c *Controller) UserLogin(w http.ResponseWriter, r *http.Request) {

	var user = &usermodel.User{}
	var validateUser = &usermodel.User{}
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

	if err = c.Service.Get(validateUser, user.Username); err != nil {
		log.Println(err)
		w.Write([]byte("Error while fetching user, " + err.Error()))
	} else {
		log.Println("User found", validateUser)

		if validateUser.Username == user.Username && validateUser.Password == user.Password {
			w.Write([]byte("Success"))
			log.Println("Student successfully logged in")

		} else {
			w.Write([]byte("Failed"))
			log.Println("Username or password is invalid")
			// http.Error(w,, http.StatusBadRequest)

		}

	}

}
