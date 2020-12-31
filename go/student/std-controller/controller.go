package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

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

	// router.Path("/").HandlerFunc()

	apiRoutes := router.PathPrefix("/").Subrouter()
	// apiRoutes.Use(validationUserToken)

	// swagger:operation GET /students get-all-students getStudents
	// ---
	// summary: Get all students
	// description: Returns Student name, email, rollno, age, date, gender, phone-number of all the students
	// responses:
	//     '200':
	//         description: Authenticated
	//     '404':
	//         description: Bad request
	getHandler := apiRoutes.HandleFunc("/students", c.GetAllStudents).Methods("GET")

	getSum := apiRoutes.HandleFunc("/students/sum", c.GetSum).Methods("GET")

	// apiRoutes.HandleFunc("/students", c.GetAllStudents).Methods("GET")

	// apiRoutes.Use(validationUserToken)
	// swagger:operation GET /students/{id} get-student getStudent
	// ---
	// summary: List the repositories owned by the given author.
	// description: Returns Student name, email, rollno, age, date, gender, phone-number of the specified students
	// parameters:
	// - name: Student ID
	//   in: path
	//   required: true
	//   type: string
	// responses:
	//     '200':
	//         description: Authenticated
	//     '404':
	//         description: Bad request
	apiRoutes.HandleFunc("/students/{id}", c.GetStudent).Methods("GET")

	// getHandlerWithID := apiRoutes.HandleFunc("/students/{id}", c.GetStudent).Methods("GET")

	excludeRoutes := []*mux.Route{getHandler, getSum}
	apiRoutes.Use(c.Authorization(excludeRoutes))

	// swagger:operation POST /students add-student addStudent
	// ---
	// summary: Student gets all students data.
	// description: Get all students info
	// parameters:
	// - name: Student Name
	//   in: body
	//   required: true
	//   type: string
	// - name: Student Age
	//   in: body
	//   required: false
	//   type: integer
	// - name: Student Roll No
	//   in: body
	//   required: false
	//   type: integer
	// - name: Student Email
	//   in: body
	//   required: true
	//   type: string
	// - name: Student Phone Number
	//   in: body
	//   required: false
	//   type: string
	// - name: DOB
	//   in: body
	//   required: false
	//   type: string
	// - name: Student Gender
	//   in: body
	//   required: false
	//   type: boolean
	// responses:
	//     '200':
	//         description: Authenticated
	//     '404':
	//         description: Bad request
	apiRoutes.HandleFunc("/students", c.AddNewStudent).Methods("POST")

	// swagger:operation PUT /students/{id} update-student updateStudent
	// ---
	// summary: Update student details
	// description: Update student
	// parameters:
	// - name: Student ID
	//   required: true
	//   in: path
	//   type: string
	// responses:
	//     '200':
	//         description: Authenticated
	//     '404':
	//         description: Bad request
	apiRoutes.HandleFunc("/students/{id}", c.UpdateStudent).Methods("PUT")

	// swagger:operation DELETE /students/{id} delete-student deleteStudent
	// ---
	// summary: Delete Student details
	// description: Delete Student
	// parameters:
	// - name: Student ID
	//   required: true
	//   in: path
	//   type: string
	// responses:
	//     '200':
	//         description: Authenticated
	//     '404':
	//         description: Bad request
	apiRoutes.HandleFunc("/students/{id}", c.DeleteStudent).Methods("DELETE")

}

func (c *Controller) Authorization(excludedRoutes []*mux.Route) func(http.Handler) http.Handler {
	// Cache the regex object of each route (obviously for performance purposes)

	var excludedRoutesRegexp []*regexp.Regexp
	rl := len(excludedRoutes)
	for i := 0; i < rl; i++ {
		r := excludedRoutes[i]
		// log.Println("Routes -> ", r)
		pathRegexp, _ := r.GetPathRegexp()
		log.Println("Path Regexp -> ", pathRegexp)

		regx, _ := regexp.Compile(pathRegexp)
		log.Println("Regx for comparing -> ", regx)

		excludedRoutesRegexp = append(excludedRoutesRegexp, regx)
	}
	log.Println("ExculdedRoutes -> ", excludedRoutesRegexp)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			log.Println("Inside validation")

			exclude := false
			requestMethod := r.Method

			log.Println("Request Method -> ", requestMethod)

			for i := 0; i < rl; i++ {
				excludedRoute := excludedRoutes[i]
				methods, _ := excludedRoute.GetMethods()
				ml := len(methods)
				log.Println("Route Method ->", methods, "lenght -> ", ml)

				methodMatched := false
				if ml < 1 {
					log.Println("Making method matched true")
					methodMatched = true
				} else {
					for j := 0; j < ml; j++ {
						log.Println("Methods[j] -> ", methods[j], "Request Method -> ", requestMethod)
						if methods[j] == requestMethod {
							methodMatched = true
							break
						}
					}
				}
				log.Println("Matched ->", methodMatched)
				if methodMatched {
					uri := r.RequestURI
					log.Println("Excluded Routes ->", excludedRoutesRegexp[i], "URI -> ", uri)
					if excludedRoutesRegexp[i].MatchString(uri) {
						exclude = true
						break
					}
				}
			}
			if !exclude {
				// validationUserToken(next)
				log.Println("Token -> ", r.Header["Token"])
				var jwtKey = []byte("some_secret_key")

				if r.Header["Token"] != nil {

					claims := &model.Claim{}

					token, err := jwt.ParseWithClaims(r.Header["Token"][0], claims, func(token *jwt.Token) (interface{}, error) {
						if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
							return nil, fmt.Errorf("There was an error")
						}
						return jwtKey, nil
					})
					if err != nil {
						if err == jwt.ErrSignatureInvalid {
							w.WriteHeader(http.StatusUnauthorized)
							return
						}
						w.WriteHeader(http.StatusBadRequest)
						return
					}

					log.Println("Token->", *token)
					log.Println("Claims->", time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))

					// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 60*time.Second {
					// 	refresh
					// }

					if token.Valid {
						next.ServeHTTP(w, r)
					}
				} else {
					http.Error(w, "User Not Authorized", http.StatusUnauthorized)
					// fmt.Fprintf(w, "Not Authorized")
					log.Println("Hello world")
				}
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}

func validationUserToken(endpoint http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var jwtKey = []byte("some_secret_key")

		if r.Header["Token"][0] != "" {

			claims := &model.Claim{}

			token, err := jwt.ParseWithClaims(r.Header["Token"][0], claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return jwtKey, nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Println("Token->", *token)
			log.Println("Claims->", time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))

			if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 60*time.Second {
				// refresh
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// w.Write([]byte(student.ID.String()))
	log.Println("Student successfully deleted")

}

func (c *Controller) GetSum(w http.ResponseWriter, r *http.Request) {

	var err error

	var students = &model.Student{}

	result, err := c.Service.GetSum(students)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Sum of age and roll no is -> " + strconv.FormatInt(result, 10)))
	log.Println("Sum of age and roll no is -> ", result)

}
