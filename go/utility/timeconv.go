package utility

import (
	"github.com/shaileshhb/restapi/model"
)

func TrimDate(students *[]model.Student) {

	tempStudentsDate := *students

	for i := 0; i < len(*students); i++ {
		if tempStudentsDate[i].Date == nil {
			continue
		}
		trimDate := *tempStudentsDate[i].Date
		trimDate = trimDate[:10]
		tempStudentsDate[i].Date = &trimDate
	}

}

func TrimDateTime(students *[]model.Student) {

	tempStudentsDate := *students

	for i := 0; i < len(*students); i++ {
		if tempStudentsDate[i].Date == nil {
			continue
		}
		trimDate := *tempStudentsDate[i].DateTime
		trimTime := *tempStudentsDate[i].DateTime
		trimDate = trimDate[:10]
		trimTime = trimTime[11:19]
		tempDateTime := trimDate + "T" + trimTime
		tempStudentsDate[i].DateTime = &tempDateTime
	}
}
