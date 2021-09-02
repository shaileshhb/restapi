package utility

import (
	"errors"

	"github.com/shaileshhb/restapi/model/bookissue"
)

func ValidateBookIssue(bookIssue *bookissue.BookIssue, issue []bookissue.BookIssue) error {

	for i := 0; i < len(issue); i++ {
		if issue[i].StudentID == bookIssue.StudentID && issue[i].ReturnedFlag == false {
			if issue[i].BookID == bookIssue.BookID {
				return errors.New("Book Already issued")
			}
		}
	}
	return nil
}
