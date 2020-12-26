package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	stdmodel "github.com/shaileshhb/restapi/student/std-model"
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

	// apiRoutes := router.PathPrefix("/students/api").Subrouter()

	// apiRoutes.Use(validationUserToken)

	// apiRoutes.HandleFunc("/students", c.GetAllStudents).Methods("GET")
	// apiRoutes.HandleFunc("/students/{id}", c.GetStudent).Methods("GET")
	// apiRoutes.HandleFunc("/students", c.AddNewStudent).Methods("POST")
	// apiRoutes.HandleFunc("/students/{id}", c.UpdateStudent).Methods("PUT")
	// apiRoutes.HandleFunc("/students/{id}", c.DeleteStudent).Methods("DELETE")

	router.Handle("/students", validationUserToken(c.GetAllStudents)).Methods("GET")
	router.Handle("/students/{id}", validationUserToken(c.GetStudent)).Methods("GET")
	router.Handle("/students", validationUserToken(c.AddNewStudent)).Methods("POST")
	router.Handle("/students/{id}", validationUserToken(c.UpdateStudent)).Methods("PUT")
	router.Handle("/students/{id}", validationUserToken(c.DeleteStudent)).Methods("DELETE")

}

// func validationUserToken(endpoint http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		log.Println("Inside Validation token")

// 		var jwtKey = []byte("some_secret_key")

// 		log.Println(r.Header["Token"])

// 		if r.Header["Token"][0] != "" {

// 			log.Println("Inside validation")

// 			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}
// 				return jwtKey, nil
// 			})

// 			if err != nil {
// 				fmt.Fprintf(w, err.Error())
// 			}

// 			if token.Valid {
// 				endpoint.ServeHTTP(w, r)
// 			}
// 		} else {
// 			http.Error(w, "User Not Authorized", http.StatusUnauthorized)
// 			// fmt.Fprintf(w, "Not Authorized")
// 		}
// 	})
// }

func validationUserToken(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			http.Error(w, "User Not Authorized", http.StatusUnauthorized)
			// fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func (c *Controller) GetAllStudents(w http.ResponseWriter, r *http.Request) {

	var err error

	log.Printf("\n\nINSIDE GET ALL STUDENT\n\n")

	// userToken, err := c.validationUserToken(r)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// if !userToken.Valid {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	var students = []stdmodel.Student{}
	err = c.Service.GetAll(&students)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Student not found"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if studentJSON, err := json.Marshal(students); err != nil {
		log.Println(err)
		w.Write([]byte("Could not convert to json"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.Write(studentJSON)
		log.Println("Student Successfully returned")
	}
}

func (c *Controller) GetStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	var students = []stdmodel.Student{}
	params := mux.Vars(r)

	err = c.Service.Get(&students, params["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte("Student Not Found"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if studentJSON, err := json.Marshal(students); err != nil {
		log.Println(err)
		w.Write([]byte("Could not convert to json"))
		http.Error(w, err.Error(), http.StatusBadRequest)

	} else {
		w.Write(studentJSON)
		log.Println("Student successfully returned")
	}

}

func (c *Controller) AddNewStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	log.Printf("\n\nINSIDE ADD STUDENT\n\n")

	var student = &stdmodel.Student{}
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

	if err := c.Service.AddNewStudent(student); err != nil {
		log.Println(err)
		w.Write([]byte("Error while adding student, " + err.Error()))
	} else {
		w.Write([]byte(student.ID.String()))
		log.Println("Student successfully added", student.ID)
	}

}

func (c *Controller) UpdateStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	var student = &stdmodel.Student{}

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

	if err := c.Service.Update(student, params["id"]); err != nil {
		log.Println(err)
		w.Write([]byte("Error while updating student"))
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Write([]byte(params["id"]))
		log.Println("Student successfully updated", *student)
	}

}

func (c *Controller) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	var student = &stdmodel.Student{}
	params := mux.Vars(r)

	if err = c.Service.Delete(student, params["id"]); err != nil {
		log.Println(err)
		w.Write([]byte("Error while deleting student"))
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Write([]byte(student.ID.String()))
		log.Println("Student successfully deleted", student.ID)
	}
}
