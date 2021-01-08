package utility

import (
	"errors"
	"log"

	"github.com/shaileshhb/restapi/model"
)

func ValidateBookIssue(bookIssue *model.BookIssue, issue []model.BookIssue) error {

	log.Println("Book Issue -> ", bookIssue)
	log.Println("Issue -> ", issue)

	for i := 0; i < len(issue); i++ {
		log.Println("student.BookIssues Flag -> ", *issue[i].ReturnedFlag)
		if issue[i].StudentID == bookIssue.StudentID && *issue[i].ReturnedFlag == false {
			log.Println("student.BookIssues -> ", issue[i])
			if issue[i].BookID == bookIssue.BookID {
				return errors.New("Book Already issued")
			}
		}
	}
	return nil
}
