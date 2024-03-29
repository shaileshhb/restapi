package utility

import "github.com/shaileshhb/restapi/model/student"

func TrimDate(students *student.Student) {

	if students.Date != nil {
		trimDate := *students.Date
		trimDate = trimDate[:10]
		students.Date = &trimDate
	}

}

func TrimDates(students *[]student.Student) {

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
