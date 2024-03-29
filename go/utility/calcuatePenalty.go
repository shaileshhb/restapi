package utility

import (
	"log"
	"time"

	"github.com/shaileshhb/restapi/model/bookissue"
)

func Penalty(bookIssues *[]bookissue.BookIssue) error {

	for i := 0; i < len(*bookIssues); i++ {

		t, err := time.Parse(time.RFC3339, (*bookIssues)[i].IssueDate)
		if err != nil {
			log.Println("ERROR in date -> ", err)
			return err
		}

		diff := time.Now().Sub(t)
		days := int(diff.Hours() / 24)
		isBookReturned := (*bookIssues)[i].ReturnedFlag

		log.Println("Diff -> ", diff)
		log.Println("Current Date -> ", (*bookIssues)[i].IssueDate, "days -> ", days)
		log.Println("Penalty Before-> ", (*bookIssues)[i].Penalty)

		if days > 10 && !isBookReturned {
			(*bookIssues)[i].Penalty = float64(days-10) * 2.5
			log.Println("Updated Penalty -> ", ((*bookIssues)[i].Penalty))
		}

		if isBookReturned {
			(*bookIssues)[i].Penalty = 0.0
		}

		(*bookIssues)[i].IssueDate = (*bookIssues)[i].IssueDate
	}
	return nil
}
