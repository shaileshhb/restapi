package bookcontroller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/book/bookservice"
	"github.com/shaileshhb/restapi/model/book"
	"github.com/shaileshhb/restapi/web"
)

type BookController struct {
	service *bookservice.BookService
}

func NewBookController(service *bookservice.BookService) *BookController {
	return &BookController{
		service: service,
	}
}

func (c *BookController) RegisterBookRoutes(getRouter, middlewareRouter *mux.Router) {

	getRouter.HandleFunc("/books", c.GetAllBooks).Methods(http.MethodGet)
	getRouter.HandleFunc("/book/{id}", c.GetBook).Methods(http.MethodGet)

	middlewareRouter.HandleFunc("/books", c.AddBook).Methods(http.MethodPost)
	middlewareRouter.HandleFunc("/books/{id}", c.UpdateBook).Methods(http.MethodPut)
	middlewareRouter.HandleFunc("/books/{id}", c.DeleteBook).Methods(http.MethodDelete)
}

func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("In GetAllBooks")

	var err error

	var books = []book.BookAvailability{}

	if err = c.service.GetAllBooks(&books); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	web.RespondJSON(w, http.StatusOK, books)
}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {

	var book = book.Book{}
	params := mux.Vars(r)

	err := c.service.GetBook(&book, params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	web.RespondJSON(w, http.StatusOK, book)
}

func (c *BookController) AddBook(w http.ResponseWriter, r *http.Request) {
	log.Println("In AddBook")

	var book = &book.Book{}

	err := web.UnmarshalJSON(r, book)
	if err != nil {
		log.Println("error from add", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Book in controller %#v", *book)

	err = c.service.AddNewBook(book)
	if err != nil {
		log.Println("error from add", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	web.RespondJSON(w, http.StatusOK, "Book successfully added")
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {

	var book = &book.Book{}

	params := mux.Vars(r)

	err := web.UnmarshalJSON(r, book)
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
	web.RespondJSON(w, http.StatusOK, "Book successfully updated")

	// w.Write([]byte("Book successfully updated"))
	// log.Println("Book successfully updated")
}

func (i *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {

	var err error

	var book = &book.Book{}
	params := mux.Vars(r)

	err = i.service.Delete(book, params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// w.Write([]byte(book.ID.String()))
	// log.Println("Book successfully deleted")
	web.RespondJSON(w, http.StatusOK, "Book successfully deleted")
}
