package calculatepenalty

import (
	"log"
	"time"

	"github.com/shaileshhb/restapi/model"
)

func Penalty(bookIssues *[]model.BookIssue) error {

	for i := 0; i < len(*bookIssues); i++ {

		t, err := time.Parse(time.RFC3339, *(*bookIssues)[i].IssueDate)
		if err != nil {
			log.Println("ERROR in date -> ", err)
			return err
		}
		diff := time.Now().Sub(t)
		days := int(diff.Hours() / 24)
		isBookReturned := *(*bookIssues)[i].ReturnedFlag

		log.Println("Diff -> ", diff)
		log.Println("Current Date -> ", *(*bookIssues)[i].IssueDate, "days -> ", days)

		if days > 10 && !isBookReturned {
			(*bookIssues)[i].Penalty = (days - 10) * 2
			log.Println("Updated Penalty -> ", ((*bookIssues)[i].Penalty))

		}

		if isBookReturned {
			(*bookIssues)[i].Penalty = 0
		}
	}
	return nil
}
