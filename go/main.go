// Student API
//
// Documentation for Student API.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.1.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/book/bookcontroller"
	"github.com/shaileshhb/restapi/book/bookservice"
	"github.com/shaileshhb/restapi/bookissue/bookissuecontroller"
	"github.com/shaileshhb/restapi/bookissue/bookissueservice"
	"github.com/shaileshhb/restapi/config"
	"github.com/shaileshhb/restapi/model/book"
	"github.com/shaileshhb/restapi/model/bookissue"
	"github.com/shaileshhb/restapi/model/student"
	"github.com/shaileshhb/restapi/model/user"
	"github.com/shaileshhb/restapi/repository"
	md "github.com/shaileshhb/restapi/security/middleware"
	stdcontroller "github.com/shaileshhb/restapi/student/std-controller"
	stdservice "github.com/shaileshhb/restapi/student/std-service"
	usercontroller "github.com/shaileshhb/restapi/user/usercontroller"
	userservice "github.com/shaileshhb/restapi/user/userservice"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBDatabase)

	fmt.Println(DBURL)

	// "root:root@tcp(localhost:3306)/student_app?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	fmt.Println("DB connected Successfully")

	router := mux.NewRouter().StrictSlash(true)
	if router == nil {
		log.Fatal("No Route Created")
	}

	middlewareRouter := router.PathPrefix("/").Subrouter()
	getRouter := router.PathPrefix("/").Subrouter()
	middlewareRouter.Use(md.Middleware)

	repos := repository.NewGormRepository()

	db = db.AutoMigrate(&user.User{}, &student.Student{}, &book.Book{}, &bookissue.BookIssue{})

	// Setting Foreign keys
	db = db.Model(&bookissue.BookIssue{}).AddForeignKey("student_id", "students(id)", "RESTRICT", "RESTRICT")
	db = db.Model(&bookissue.BookIssue{}).AddForeignKey("book_id", "books(id)", "RESTRICT", "RESTRICT")

	RegisterControllerAndService(middlewareRouter, getRouter, repos, db)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Token"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origin := handlers.AllowedOrigins([]string{"*"})

	srv := &http.Server{
		Addr:         ":8081",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
		Handler:      handlers.CORS(headers, methods, origin)(router),
	}

	var wait time.Duration

	go func() {
		fmt.Println(" ======= Listening at port", srv.Addr)
		log.Fatal(srv.ListenAndServe())
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	func() {
		fmt.Println("Closing DB")
		db.Close()
	}()
	fmt.Println("Server ShutDown.......")
	os.Exit(0)
	fmt.Println("========== waiting ==========")
}

func RegisterControllerAndService(middlewareRouter, getRouter *mux.Router,
	repos *repository.GormRepository, db *gorm.DB) {

	//login
	userService := userservice.NewUserService(repos, db)
	userController := usercontroller.NewController(userService)
	userController.RegisterUserRoutes(getRouter, middlewareRouter)

	//student
	serv := stdservice.NewService(repos, db)
	controller := stdcontroller.NewController(serv)
	controller.RegisterRoutes(getRouter, middlewareRouter)

	// book
	bookserv := bookservice.NewBookService(repos, db)
	bookController := bookcontroller.NewBookController(bookserv)
	bookController.RegisterBookRoutes(getRouter, middlewareRouter)

	// issues
	issuesServ := bookissueservice.NewIssueService(repos, db)
	issueController := bookissuecontroller.NewBookIssueController(issuesServ)
	issueController.RegisterBookIssueRoutes(getRouter, middlewareRouter)
}

// availability query
// select id, stock,
// 	if(sum(returned_flag=0)>0, abs(stock - sum(returned_flag=0)), stock) total,
// 	returned_flag
// from books
// left join book_issues
// on id = book_id
// group by id

// When if statement removed then it returns null

// penalty update
// SELECT book_id, student_id, issue_date, returned_flag, abs(DATEDIFF(issue_date, CURDATE()))-10.0,
// 	if(abs(DATEDIFF(issue_date, CURDATE())) > 10 and returned_flag = false, (abs(DATEDIFF(issue_date, CURDATE()))-10)*2.5, 0) penalty
// FROM book_issues

// Query 1
// select (count(books.id) - count(book_issues.book_id)) - sum(penalty) as diff
// from books
// left join book_issues
// on books.id = book_id
