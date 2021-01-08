package bookcontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/book/bookservice"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/utility"
)

type BookController struct {
	service *bookservice.BookService
}

func NewBookController(service *bookservice.BookService) *BookController {
	return &BookController{
		service: service,
	}
}

func (c *BookController) RegisterBookRoutes(router *mux.Router) {

	apiRoutes := router.PathPrefix("/").Subrouter()

	apiRoutes.Use()

	getBooks := apiRoutes.HandleFunc("/books", c.GetAllBooks).Methods("GET")
	getBook := apiRoutes.HandleFunc("/books/{id}", c.GetBook).Methods("GET")

	excludedRoutes := []*mux.Route{getBooks, getBook}
	apiRoutes.Use(utility.Authorization(excludedRoutes))

	apiRoutes.HandleFunc("/books", c.AddBook).Methods("POST")
	apiRoutes.HandleFunc("/books/{id}", c.UpdateBook).Methods("PUT")
	apiRoutes.HandleFunc("/books/{id}", c.DeleteBook).Methods("DELETE")
}

func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {

	var err error

	var books = []model.BookAvailability{}

	if err = c.service.GetAllBooks(&books); err != nil {
		log.Println(err)
		return
	}
	bookJoinJSON, err := json.Marshal(books)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(bookJoinJSON)
	// log.Println("Join Books -> ", books)
}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {

	var err error

	var book = model.Book{}
	params := mux.Vars(r)

	err = c.service.GetBook(&book, params["id"])
	if err != nil {
		log.Println(err)
		// w.Write([]byte("Book Not Found"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookJSON, err := json.Marshal(book)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(bookJSON)
	log.Println("Book successfully returned")
}

func (c *BookController) AddBook(w http.ResponseWriter, r *http.Request) {

	var err error

	var book = &model.Book{}
	bookResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		// w.Write([]byte("Response could not be read"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bookResponse, book)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Book in controller", book)

	err = c.service.AddNewBook(book)
	if err != nil {
		log.Println("error from add", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		// w.Write([]byte("Error while adding book, " + err.Error()))
		return
	}

	w.Write([]byte(book.ID.String()))
	log.Println("Book successfully added", book.ID)
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {

	var err error
	var book = &model.Book{}

	params := mux.Vars(r)

	bookResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Response could not be read"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bookResponse, book)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.Update(book, params["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while updating book"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Book successfully updated"))
	log.Println("Book successfully updated")
}

func (i *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {

	var err error

	var book = &model.Book{}
	params := mux.Vars(r)

	err = i.service.Delete(book, params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(book.ID.String()))
	log.Println("Book successfully deleted")

}
