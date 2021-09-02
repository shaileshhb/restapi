package student

import (
	"github.com/satori/uuid"
	"github.com/shaileshhb/restapi/model/bookissue"
	"github.com/shaileshhb/restapi/model/general"
)

// Student will contain all details of student.
// swagger:model
type Student struct {
	// the id for this student
	//
	// required: true
	general.Base `structs:"id,omitnested"`

	// the name for this student
	// required: true
	// max length: 30
	Name string `gorm:"type:varchar(30)" json:"name,omitempty"`

	// age of the student
	// required: false
	// min: 1
	Age *int `gorm:"type:int" json:"age,omitempty"`

	// age of the student
	// required: false
	// min: 1
	RollNo *int `gorm:"type:int" json:"rollNo,omitempty"`

	// age of the student
	// required: false
	PhoneNumber *string `gorm:"type:varchar(10)" json:"phone,omitempty"`

	// age of the student
	// required: true
	// example: user@provider.net
	Email string `gorm:"type:varchar(50)" json:"email,omitempty"`

	// age of the student
	// required: true
	IsMale *bool `gorm:"type:tinyint" json:"isMale,omitempty"`

	// age of the student
	// required: false
	Date *string `gorm:"type:date" json:"date,omitempty"`

	// the books issued to the student
	//
	BookIssues []bookissue.BookIssue `json:"bookIssues"`
}

// StudentResponse will contain all students.
// swagger:response StudentResponse
type StudentResponse struct {
	// All students from the database.
	// in: body
	Body []Student
}

// swagger:parameters updateStudent
type studentIDWrapper struct {
	// ID to update the student
	// in: path
	// required: true
	ID uuid.UUID `json:"id"`
}
