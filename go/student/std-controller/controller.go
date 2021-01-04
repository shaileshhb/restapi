package stdcontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/model"
	stdservice "github.com/shaileshhb/restapi/student/std-service"
	"github.com/shaileshhb/restapi/utility/excluderoute"
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
	getDiff := apiRoutes.HandleFunc("/students/diff", c.GetDiff).Methods("GET")
	getAgeAndRecordDiff := apiRoutes.HandleFunc("/students/recordsDiff", c.GetDiffOfAgeAndRecord).Methods("GET")
	checkAge := apiRoutes.HandleFunc("/students/age", c.GetAge).Methods("GET")

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
	getHandlerWithID := apiRoutes.HandleFunc("/students/{id}", c.GetStudent).Methods("GET")
	// apiRoutes.HandleFunc("/students/{id}", c.GetStudent).Methods("GET")

	excludeRoutes := []*mux.Route{getHandler, getHandlerWithID, getSum, getDiff, getAgeAndRecordDiff, checkAge}
	apiRoutes.Use(excluderoute.Authorization(excludeRoutes))

	// swagger:operation POST /students add-student studentModel
	// ---
	// summary: Get students data and add it to the db.
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
	//   example: domain@abc.com
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

// func validationUserToken(endpoint http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		var jwtKey = []byte("some_secret_key")

// 		if r.Header["Token"][0] != "" {

// 			claims := &model.Claim{}

// 			token, err := jwt.ParseWithClaims(r.Header["Token"][0], claims, func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}
// 				return jwtKey, nil
// 			})
// 			if err != nil {
// 				if err == jwt.ErrSignatureInvalid {
// 					w.WriteHeader(http.StatusUnauthorized)
// 					return
// 				}
// 				w.WriteHeader(http.StatusBadRequest)
// 				return
// 			}

// 			log.Println("Token->", *token)
// 			log.Println("Claims->", time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))

// 			if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 60*time.Second {
// 				// refresh
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

func (c *Controller) GetAllStudents(w http.ResponseWriter, r *http.Request) {

	var err error

	log.Printf("\nINSIDE GET ALL STUDENT\n")

	var students = []model.Student{}
	err = c.Service.GetAll(&students)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	studentJSON, err := json.Marshal(students)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(studentJSON)
	log.Println("Student Successfully returned -> ", students)

}

func (c *Controller) GetStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	var students = model.Student{}
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
	w.Write([]byte(student.ID.String()))
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
	w.Write([]byte(student.ID.String()))
	log.Println("Student successfully deleted")

}

func (c *Controller) GetSum(w http.ResponseWriter, r *http.Request) {

	var err error

	var students = &model.Student{}
	var sum = &model.Sum{}

	err = c.Service.GetSum(students, sum)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Sum of age and roll no is -> " + strconv.FormatInt(sum.N, 10)))
	log.Println("Sum of age and roll no is -> ", sum.N)

}

func (c *Controller) GetDiff(w http.ResponseWriter, r *http.Request) {

	var err error

	var students = &model.Student{}
	var sum = &model.Sum{}

	err = c.Service.GetDiff(students, sum)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Diff of age and roll no is -> " + strconv.FormatInt(sum.N, 10)))
	log.Println("Diff of age and roll no is -> ", sum.N)

}

func (c *Controller) GetDiffOfAgeAndRecord(w http.ResponseWriter, r *http.Request) {

	var err error

	var students = &model.Student{}
	var sum = &model.Sum{}

	err = c.Service.GetDiffOfAgeAndRecord(students, sum)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Age and record diff is -> " + strconv.FormatInt(sum.N, 10)))
	log.Println("Age and record diff is -> ", sum.N)

}

func (c *Controller) GetAge(w http.ResponseWriter, r *http.Request) {

	var err error

	var students = &[]model.Student{}

	err = c.Service.GetAge(students)
	if err != nil {
		log.Println(err)
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

}
