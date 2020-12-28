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
	_ "github.com/shaileshhb/restapi/docs"
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

	db.AutoMigrate(&model.User{}, &model.Student{})

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

	// db.AutoMigrate()

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
