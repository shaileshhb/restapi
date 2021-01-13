package bookissuecontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/bookissue/bookissueservice"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/utility"
)

type BookIssueController struct {
	service *bookissueservice.BookIssueService
}

func NewBookIssueController(service *bookissueservice.BookIssueService) *BookIssueController {

	return &BookIssueController{
		service: service,
	}
}

func (i *BookIssueController) RegisterBookIssueRoutes(router *mux.Router) {

	apiRoutes := router.PathPrefix("/").Subrouter()

	apiRoutes.Use()

	getBookIssues := apiRoutes.HandleFunc("/bookIssues", i.GetAllBookIssues).Methods("GET")
	getBookIssue := apiRoutes.HandleFunc("/bookIssues/{studentID}", i.GetBookIssues).Methods("GET")
	// penalty := apiRoutes.HandleFunc("/peanlty", i.GetPenalty).Methods("GET")

	excludedRoutes := []*mux.Route{getBookIssues, getBookIssue}
	apiRoutes.Use(utility.Authorization(excludedRoutes))

	apiRoutes.HandleFunc("/bookIssues", i.AddNewBookIssue).Methods("POST")
	apiRoutes.HandleFunc("/bookIssues/{id}", i.UpdateBookIssue).Methods("PUT")
	apiRoutes.HandleFunc("/bookIssues/{id}", i.DeleteBookIssue).Methods("DELETE")
}

func (i *BookIssueController) GetAllBookIssues(w http.ResponseWriter, r *http.Request) {

	var bookIssues = []model.BookIssue{}

	err := i.service.GetAll(&bookIssues)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookIssueJSON, err := json.Marshal(bookIssues)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(bookIssueJSON)
	log.Println("Book Issue Successfully returned", string(bookIssueJSON))
}

func (i *BookIssueController) GetBookIssues(w http.ResponseWriter, r *http.Request) {

	var bookIssue = []model.BookIssue{}
	params := mux.Vars(r)

	log.Println("student id ->", params["studentID"])

	err := i.service.Get(&bookIssue, params["studentID"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookIssueJSON, err := json.Marshal(bookIssue)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(bookIssueJSON)
	log.Println("Book Issue Successfully returned", bookIssue)
}

func (i *BookIssueController) AddNewBookIssue(w http.ResponseWriter, r *http.Request) {

	var err error

	log.Printf("\nINSIDE ADD Book Issue\n")

	var bookIssue = model.BookIssue{}
	issueResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(issueResponse, &bookIssue)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("===============================================================")
	log.Println("StudentID ->", bookIssue.StudentID)
	log.Println("BookID -> ", bookIssue.BookID)
	log.Println("===============================================================")

	currentTime := time.Now()
	// currentTime.Format("2006-01-02T15:04:05")
	currentTimeInString := currentTime.String()
	currentTimeInString = currentTimeInString[:19]
	bookIssue.IssueDate = currentTimeInString

	log.Println("Book ID -> ", bookIssue.BookID)

	err = i.service.AddNewBookIssue(&bookIssue)
	if err != nil {
		log.Println("error from add", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(bookIssue.ID.String()))
	log.Println("Book successfully added", bookIssue.BookID)
}

func (i *BookIssueController) UpdateBookIssue(w http.ResponseWriter, r *http.Request) {

	var err error
	var bookIssue = model.BookIssue{}

	params := mux.Vars(r)

	bookIssueResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Response could not be read"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Book Issue response", string(bookIssueResponse))
	log.Println("ID ->", params["id"])

	err = json.Unmarshal(bookIssueResponse, &bookIssue)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = i.service.UpdateBook(&bookIssue, params["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while updating book issue"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(bookIssue.ID.String()))
	log.Println("Book Issue successfully updated", bookIssue)
}

func (i *BookIssueController) DeleteBookIssue(w http.ResponseWriter, r *http.Request) {

	var err error

	var bookIssue = &model.BookIssue{}
	params := mux.Vars(r)

	err = i.service.Delete(bookIssue, params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(bookIssue.BookID.String()))
	log.Println("Book Issue successfully deleted")

}
