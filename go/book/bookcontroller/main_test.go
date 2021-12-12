package bookcontroller

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/book/bookservice"
	"github.com/shaileshhb/restapi/model/book"
	"github.com/shaileshhb/restapi/model/bookissue"
	"github.com/shaileshhb/restapi/repository"
)

const (
	DBURL = "swabhav:swabhav@tcp(localhost:3306)/swabhav_test?charset=utf8&parseTime=True&loc=Local"
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMCIsImV4cCI6MTYzODAzMDQ4Nn0.V1Sk-p_nvdp3e5IQOqYW0CufN_z3KBUvCiSarSBAQjc"
)

var repos *repository.GormRepository

var bookService *bookservice.BookService
var controller *BookController

func TestMain(m *testing.M) {

	db, err := gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	fmt.Println("Test DB connected Successfully")

	db.AutoMigrate(&book.Book{}, &bookissue.BookIssue{})

	// Setting Foreign keys
	db.Model(&bookissue.BookIssue{}).AddForeignKey("student_id", "students(id)", "RESTRICT", "RESTRICT")
	db.Model(&bookissue.BookIssue{}).AddForeignKey("book_id", "books(id)", "RESTRICT", "RESTRICT")

	repos = repository.NewGormRepository()

	bookService = bookservice.NewBookService(repos, db)
	controller = NewBookController(bookService)

	os.Exit(m.Run())
}
