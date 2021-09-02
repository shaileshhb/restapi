package stdcontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/model/student"
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

func (c *Controller) RegisterRoutes(getRouter, middlewareRouter *mux.Router) {

	getRouter.HandleFunc("/students", c.GetAllStudents).Methods("GET")
	getRouter.HandleFunc("/students/sum", c.GetSum).Methods("GET")
	getRouter.HandleFunc("/students/diff", c.GetDiff).Methods("GET")
	getRouter.HandleFunc("/students/recordsDiff", c.GetDiffOfAgeAndRecord).Methods("GET")
	getRouter.HandleFunc("/students/search", c.Search).Methods(http.MethodGet)
	getRouter.HandleFunc("/students/{id}", c.GetStudent).Methods("GET")

	middlewareRouter.HandleFunc("/students", c.AddNewStudent).Methods("POST")
	middlewareRouter.HandleFunc("/students/{id}", c.UpdateStudent).Methods("PUT")
	middlewareRouter.HandleFunc("/students/{id}", c.DeleteStudent).Methods("DELETE")

}

// swagger:route GET /students student GetAllStudents
// Returns all students.
// responses:
// 200: StudentResponse

// GetAllStudents will return all the students.
func (c *Controller) GetAllStudents(w http.ResponseWriter, r *http.Request) {

	var err error

	log.Printf("\nINSIDE GET ALL STUDENT\n")

	var students = []student.Student{}
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

	var students = student.Student{}
	params := mux.Vars(r)

	log.Println("GET STUDENT -> ", params["id"])

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

	var student = &student.Student{}
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

// swagger:route PUT /student/{id} student updateStudent
// Updates the specifie student
// responses:
// 		200: Student successfully updated

// UpdateStudent will update the specified student.
func (c *Controller) UpdateStudent(w http.ResponseWriter, r *http.Request) {

	var err error

	var student = &student.Student{}

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
	log.Println("PhoneNumber -> ", student.PhoneNumber)
	log.Println("ID -> ", params["id"])
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

	var student = &student.Student{}
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

	var students = &student.Student{}
	var sum = &student.Sum{}

	err := c.Service.GetSum(students, sum)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Sum of age and roll no is -> " + strconv.FormatInt(sum.N, 10)))
	log.Println("Sum of age and roll no is -> ", sum.N)

}

func (c *Controller) GetDiff(w http.ResponseWriter, r *http.Request) {
	var students = &student.Student{}
	var sum = &student.Sum{}

	err := c.Service.GetDiff(students, sum)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Diff of age and roll no is -> " + strconv.FormatInt(sum.N, 10)))
	log.Println("Diff of age and roll no is -> ", sum.N)

}

func (c *Controller) GetDiffOfAgeAndRecord(w http.ResponseWriter, r *http.Request) {
	var students = &student.Student{}
	var sum = &student.Sum{}

	err := c.Service.GetDiffOfAgeAndRecord(students, sum)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Age and record diff is -> " + strconv.FormatInt(sum.N, 10)))
	log.Println("Age and record diff is -> ", sum.N)

}

func (c *Controller) Search(w http.ResponseWriter, r *http.Request) {
	var students = []student.Student{}
	// var searchStudent = student.SearchStudent{}

	params := r.URL.Query()

	log.Println("Params Eamil len -> ", len(params["email"]))
	if len(params) == 0 {
		c.GetAllStudents(w, r)
		return
	}

	err := c.Service.Search(&students, params)
	if err != nil {
		log.Println("Error in Search -> ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	studentJSON, err := json.Marshal(students)
	if err != nil {
		log.Println("Error in json -> ", studentJSON)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Student Searched -> ", students)
	w.Write(studentJSON)
}
