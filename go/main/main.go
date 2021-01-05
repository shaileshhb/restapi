// Student API
//
// Example Swagger spec.
//
//	Schemes: [http]
//	BasePath: http://localhost:8080
//	Version: 0.0.1
//	Contact: ABC<admin@studentAPI.in>
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//      SecurityDefinitions:
//          basic:
//              type: basic
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
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/repository"
	stdcontroller "github.com/shaileshhb/restapi/student/std-controller"
	stdservice "github.com/shaileshhb/restapi/student/std-service"
	usercontroller "github.com/shaileshhb/restapi/user/user-controller"
	userservice "github.com/shaileshhb/restapi/user/user-service"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/student_app?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("DB connected Successfully")

	db.AutoMigrate(&model.User{}, &model.Student{}, &model.Book{}, &model.BookIssue{})

	// Setting Foreign keys
	db.Model(&model.BookIssue{}).AddForeignKey("student_id", "students(id)", "RESTRICT", "RESTRICT")
	db.Model(&model.BookIssue{}).AddForeignKey("book_id", "books(id)", "RESTRICT", "RESTRICT")

	router := mux.NewRouter()
	if router == nil {
		log.Fatal("No Route Created")
	}

	repos := repository.NewGormRepository()

	//login
	userService := userservice.NewUserService(repos, db)
	userController := usercontroller.NewController(userService)
	userController.RegisterUserRoutes(router)

	//student
	serv := stdservice.NewService(repos, db)
	controller := stdcontroller.NewController(serv)
	controller.RegisterRoutes(router)

	// book
	bookserv := bookservice.NewBookService(repos, db)
	bookController := bookcontroller.NewBookController(bookserv)
	bookController.RegisterBookRoutes(router)

	// issues
	issuesServ := bookissueservice.NewIssueService(repos, db)
	issueController := bookissuecontroller.NewBookIssueController(issuesServ)
	issueController.RegisterBookIssueRoutes(router)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Token"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})

	server := &http.Server{
		Handler:      handlers.CORS(headers, methods, origin)(router),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":8080",
	}

	var wait time.Duration

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	func() {
		fmt.Println("Closing DB")
		db.Close()
	}()
	fmt.Println("Server ShutDown.......")
	os.Exit(0)

}

// availability query
// select id, stock,
// 	if(sum(returned_flag=0)>0, abs(stock - sum(returned_flag=0)), stock) total,
// 	returned_flag
// from books
// left join book_issues
// on id = book_id
// group by id


// penalty update
// SELECT book_id, student_id, issue_date, returned_flag, abs(DATEDIFF(issue_date, CURDATE()))-10.0,
// 	if(abs(DATEDIFF(issue_date, CURDATE())) > 10 and returned_flag = false, (abs(DATEDIFF(issue_date, CURDATE()))-10)*2.5, 0) penalty
// FROM book_issues