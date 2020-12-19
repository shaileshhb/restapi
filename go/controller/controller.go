package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/model"
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

	var students = []model.Student{}
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

	var students = []model.Student{}
	var err error

	params := mux.Vars(r)

	err = c.Service.Get(&students, params["id"])
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

	fmt.Println(students[0].Date[:10])

}

func (c *Controller) AddNewStudent(w http.ResponseWriter, r *http.Request) {

	var student = &model.Student{}
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

	var student = &model.Student{}

	params := mux.Vars(r)

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

	if err := c.Service.Update(student, params["id"]); err != nil {
		fmt.Println(err)
		w.Write([]byte("Error while updating student"))
	} else {
		w.Write([]byte(student.ID.String()))
	}

}

func (c *Controller) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	var student = &model.Student{}

	params := mux.Vars(r)

	if err := c.Service.Delete(student, params["id"]); err != nil {
		fmt.Println(err)
		w.Write([]byte("Error while deleting student"))
	} else {
		w.Write([]byte(student.ID.String()))

	}
}
