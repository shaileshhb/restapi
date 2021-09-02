package main

import (
	"log"
	"os"

	"github.com/shaileshhb/restapi/app/infrastructure/db"
)

func main() {

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	repository, err := db.DBConnenction(dbdriver, host, port, user, dbname, password)
	if err != nil {
		log.Fatal(err)
	}
	defer repository.Close()

	repository.Automigrate()

	// token := auth.NewToken()
}
