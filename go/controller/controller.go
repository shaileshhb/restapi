package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Println(err)
		w.Write([]byte("Student not found"))
		return
	}

	if studentJSON, err := json.Marshal(students); err != nil {
		log.Println(err)
		w.Write([]byte("Could not convert to json"))
	} else {
		w.Write(studentJSON)
		log.Println("Student Successfully returned")
	}
}

func (c *Controller) GetStudent(w http.ResponseWriter, r *http.Request) {

	var students = []model.Student{}
	var err error

	params := mux.Vars(r)

	err = c.Service.Get(&students, params["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte("Student Not Found"))
		return
	}

	if studentJSON, err := json.Marshal(students); err != nil {
		log.Println(err)
		w.Write([]byte("Could not convert to json"))
	} else {
		w.Write(studentJSON)
		log.Println("Student successfully returned")
	}

}

func (c *Controller) AddNewStudent(w http.ResponseWriter, r *http.Request) {

	var student = &model.Student{}
	var err error

	studentResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Response could not be read"))
		return
	}

	err = json.Unmarshal(studentResponse, &student)
	// err = json.Unmarshal([]byte(student.DateTime.String()), &student)
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

	var student = &model.Student{}

	params := mux.Vars(r)

	studentResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Response could not be read"))
		return
	}

	err = json.Unmarshal(studentResponse, &student)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.Update(student, params["id"]); err != nil {
		log.Println(err)
		w.Write([]byte("Error while updating student"))
	} else {
		w.Write([]byte(student.ID.String()))
		log.Println("Student successfully updated", student.ID)
	}

}

func (c *Controller) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	var student = &model.Student{}

	params := mux.Vars(r)

	if err := c.Service.Delete(student, params["id"]); err != nil {
		log.Println(err)
		w.Write([]byte("Error while deleting student"))
	} else {
		w.Write([]byte(student.ID.String()))
		log.Println("Student successfully deleted", student.ID)

	}
}
