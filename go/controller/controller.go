package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/repository"
	"github.com/shaileshhb/restapi/service"
)

type Controller struct {
	Service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) RegisterEndPoints(router *mux.Router) {

	router.HandleFunc("/students", c.GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", c.GetStudent).Methods("GET")
	router.HandleFunc("/students", c.AddNewStudent).Methods("POST")
	router.HandleFunc("/students/{id}", c.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", c.DeleteStudent).Methods("DELETE")

}

func (c *Controller) GetAllStudents(w http.ResponseWriter, r *http.Request) {

	students := []model.Student{}
	err := c.Service.GetAll(&students)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Student not found"))
		return
	}

	if studentJSON, err := json.Marshal(students); err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not convert to json"))
	} else {
		w.Write(studentJSON)
	}
}

func (c *Controller) GetStudent(w http.ResponseWriter, r *http.Request) {

	students := []model.Student{}
	var err error
	var queryProcessors []repository.QueryProcessor

	params := mux.Vars(r)
	queryProcessors = append(queryProcessors, repository.GetStudentByID(params["id"]))

	err = c.Service.Get(&students, queryProcessors)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Student Not Found"))
		return
	}

	if studentJSON, err := json.Marshal(students); err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not convert to json"))
	} else {
		w.Write(studentJSON)
	}

}

func (c *Controller) AddNewStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student
	var err error

	studentResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Response could not be read"))
		return
	}
	err = json.Unmarshal(studentResponse, &student)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.AddNewStudent(student); err != nil {
		fmt.Println(err)
		w.Write([]byte("Error while adding student"))
	} else {
		w.Write([]byte(student.ID.String()))
	}

}

func (c *Controller) UpdateStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student
	var queryProcessors []repository.QueryProcessor

	params := mux.Vars(r)

	studentID, err := uuid.FromString(params["studentID"])
	if err != nil {
		fmt.Println(err)
	}
	student.ID = studentID

	studentResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Response could not be read"))
		return
	}

	queryProcessors = append(queryProcessors, repository.GetStudentByID(params["id"]))

	err = json.Unmarshal(studentResponse, &student)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.Update(student, queryProcessors); err != nil {
		fmt.Println(err)
		w.Write([]byte("Error while updating student"))
	} else {
		w.Write([]byte(student.ID.String()))
	}

}

func (c *Controller) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student
	var queryProcessors []repository.QueryProcessor
	params := mux.Vars(r)

	queryProcessors = append(queryProcessors, repository.GetStudentByID(params["id"]))
	if err := c.Service.Delete(&student, queryProcessors); err != nil {
		fmt.Println(err)
		w.Write([]byte("Error while deleting student"))
	} else {
		w.Write([]byte(student.ID.String()))

	}
}
