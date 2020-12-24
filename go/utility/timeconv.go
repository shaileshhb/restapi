package utility

import (
	stdmodel "github.com/shaileshhb/restapi/student/stdmodel"
)

func ConvertDateTime(students *[]stdmodel.Student) {

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

func TrimDateTime(students *[]stdmodel.Student) {

	tempStudentsDate := *students

	for i := 0; i < len(*students); i++ {
		if tempStudentsDate[i].Date == nil {
			continue
		}
		trimDate := *tempStudentsDate[i].DateTime
		trimTime := *tempStudentsDate[i].DateTime
		trimDate = trimDate[:10]
		trimTime = trimTime[11:19]
		tempDateTime := trimDate + " " + trimTime
		tempStudentsDate[i].DateTime = &tempDateTime
	}
}
