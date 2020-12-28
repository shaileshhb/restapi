package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/model"
	stdservice "github.com/shaileshhb/restapi/student/std-service"
)

type Controller struct {
	Service *stdservice.Service
}

func NewController(service *stdservice.Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) RegisterRoutes(router *mux.Router) {

	apiRoutes := router.PathPrefix("/").Subrouter()

	apiRoutes.Use(validationUserToken)

	apiRoutes.HandleFunc("/students", c.GetAllStudents).Methods("GET")
	apiRoutes.HandleFunc("/students/{id}", c.GetStudent).Methods("GET")
	apiRoutes.HandleFunc("/students", c.AddNewStudent).Methods("POST")
	apiRoutes.HandleFunc("/students/{id}", c.UpdateStudent).Methods("PUT")
	apiRoutes.HandleFunc("/students/{id}", c.DeleteStudent).Methods("DELETE")

	// router.Handle("/students", validationUserToken(c.GetAllStudents)).Methods("GET")
	// router.Handle("/students/{id}", validationUserToken(c.GetStudent)).Methods("GET")
	// router.Handle("/students", validationUserToken(c.AddNewStudent)).Methods("POST")
	// router.Handle("/students/{id}", validationUserToken(c.UpdateStudent)).Methods("PUT")
	// router.Handle("/students/{id}", validationUserToken(c.DeleteStudent)).Methods("DELETE")

}

func validationUserToken(endpoint http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Inside Validation token")

		var jwtKey = []byte("some_secret_key")

		log.Println(r.Header["Token"])

		if r.Header["Token"][0] != "" {

			log.Println("Inside validation")

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return jwtKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
				http.Error(w, err.Error(), http.StatusUnauthorized)

			}

			if token.Valid {
				endpoint.ServeHTTP(w, r)
			}
		} else {
			http.Error(w, "User Not Authorized", http.StatusUnauthorized)
			// fmt.Fprintf(w, "Not Authorized")
		}
	})
}

// func validationUserToken(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		var jwtKey = []byte("some_secret_key")
// 		// log.Println("Header ->", r.Header)

// 		if r.Header["Token"][0] != "" {

// 			log.Println("Inside validation")

// 			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}
// 				return jwtKey, nil
// 			})

// 			if err != nil {
// 				log.Println("invalid signature")
// 				// fmt.Fprintf(w, "Session Expired")

// 				http.Error(w, err.Error(), http.StatusUnauthorized)
// 				// return
// 			}

// 			if token.Valid {
// 				endpoint(w, r)
// 			}
// 		} else {
// 			log.Println("cookie header not found")
// 			http.Error(w, "User Not Authorized", http.StatusUnauthorized)
// 			// fmt.Fprintf(w, "Not Authorized")
// 		}
// 	})
// }

func (c *Controller) GetAllStudents(w http.ResponseWriter, r *http.Request) {

	var err error

	log.Printf("\nINSIDE GET ALL STUDENT\n")

	var students = []model.Student{}
	err = c.Service.GetAll(&students)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Student not found"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	studentJSON, err := json.Marshal(students)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Could not convert to json"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(studentJSON)
	log.Println("Student Successfully returned")

}

func (c *Controller) GetStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	var students = []model.Student{}
	params := mux.Vars(r)

	err = c.Service.Get(&students, params["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte("Student Not Found"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	studentJSON, err := json.Marshal(students)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Could not convert to json"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(studentJSON)
	log.Println("Student successfully returned")

}

func (c *Controller) AddNewStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	log.Printf("\nINSIDE ADD STUDENT\n")

	var student = &model.Student{}
	studentResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		// w.Write([]byte("Response could not be read"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(studentResponse, student)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Student in controller", student)

	err = c.Service.AddNewStudent(student)
	if err != nil {
		log.Println("error from add", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		// w.Write([]byte("Error while adding student, " + err.Error()))
		return
	}

	w.Write([]byte(student.ID.String()))
	log.Println("Student successfully added", student.ID)

}

func (c *Controller) UpdateStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	var student = &model.Student{}

	params := mux.Vars(r)

	studentResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Response could not be read"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(studentResponse, student)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Service.Update(student, params["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while updating student"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// w.Write([]byte("Student successfully updated"))
	log.Println("Student successfully updated")

}

func (c *Controller) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	var student = &model.Student{}
	params := mux.Vars(r)

	err = c.Service.Delete(student, params["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while deleting student"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// w.Write([]byte(student.ID.String()))
	log.Println("Student successfully deleted")

}
