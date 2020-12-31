package issuecontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/bookissue/issueservice"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/utility/excluderoute"
)

type IssueController struct {
	service *issueservice.IssueService
}

func NewIssueController(service *issueservice.IssueService) *IssueController {

	return &IssueController{
		service: service,
	}
}

func (i *IssueController) RegisterIssueRoutes(router *mux.Router) {

	apiRoutes := router.PathPrefix("/").Subrouter()

	apiRoutes.Use()

	getBookIssues := apiRoutes.HandleFunc("/bookIssues", i.GetAllIssues).Methods("GET")
	getBookIssue := apiRoutes.HandleFunc("/bookIssues/{id}", i.GetIssue).Methods("GET")

	excludedRoutes := []*mux.Route{getBookIssues, getBookIssue}
	apiRoutes.Use(excluderoute.Authorization(excludedRoutes))

	apiRoutes.HandleFunc("/bookIssues", i.AddNewBookIssue).Methods("POST")
	apiRoutes.HandleFunc("/bookIssues", i.DeleteBookIssue).Methods("DELETE")
}

func (i *IssueController) GetAllIssues(w http.ResponseWriter, r *http.Request) {

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
	log.Println("Book Issue Successfully returned")
}

func (i *IssueController) GetIssue(w http.ResponseWriter, r *http.Request) {

	var bookIssue = model.BookIssue{}
	params := mux.Vars(r)

	err := i.service.Get(&bookIssue, params["id"])
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
	log.Println("Book Issue Successfully returned")
}

func (i *IssueController) AddNewBookIssue(w http.ResponseWriter, r *http.Request) {

	var err error

	log.Printf("\nINSIDE ADD STUDENT\n")

	var bookIssue = &model.BookIssue{}
	issueRespomse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		// w.Write([]byte("Response could not be read"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(issueRespomse, bookIssue)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Book in controller", bookIssue)

	err = i.service.AddNewBook(bookIssue)
	if err != nil {
		log.Println("error from add", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		// w.Write([]byte("Error while adding student, " + err.Error()))
		return
	}

	w.Write([]byte(bookIssue.BookID.String()))
	log.Println("Book successfully added", bookIssue.BookID)
}

func (i *IssueController) DeleteBookIssue(w http.ResponseWriter, r *http.Request) {

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
	log.Println("Book successfully deleted")

}
