package structToMap

import (
	"log"

	"github.com/fatih/structs"
	"github.com/shaileshhb/restapi/model"
)

func ConvertStructToMap(student *model.Student, id string) map[string]interface{} {

	studentMap := structs.Map(student)
	log.Println("Students -> ", studentMap)

	if student.RollNo == nil {
		studentMap["RollNo"] = nil
	} else {
		studentMap["RollNo"] = *student.RollNo
	}

	if student.Age == nil {
		studentMap["Age"] = nil
	} else {
		studentMap["Age"] = *student.Age
	}

	if student.PhoneNumber == nil {
		studentMap["PhoneNumber"] = nil
	} else {
		studentMap["PhoneNumber"] = *student.PhoneNumber
	}

	if student.Date == nil {
		studentMap["Date"] = nil
	} else {
		studentMap["Date"] = *student.Date
	}

	if student.DateTime == nil {
		studentMap["DateTime"] = nil
	} else {
		studentMap["DateTime"] = *student.DateTime
	}

	log.Println("In struct to map, student:", student)

	// if student.IsMale == nil {
	// 	studentMap["IsMale"] = nil
	// } else {
	// 	studentMap["IsMale"] = *student.IsMale
	// }

	log.Println("Student in struct to map func:", studentMap)

	return studentMap

}

func EmptyToNull(student *model.Student) {

	if student.PhoneNumber != nil {
		if *student.PhoneNumber == "" {
			student.PhoneNumber = nil
		}
	}

	if student.Date != nil {
		if *student.Date == "" {
			student.Date = nil
		}
	}

	if student.DateTime != nil {
		if *student.DateTime == "" {
			student.DateTime = nil
		}
	}
}
