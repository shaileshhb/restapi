package entity

import "github.com/satori/uuid"

// BookIssue will consist of data regarding which book is issued to which student.
type BookIssue struct {
	Base
	BookID       *uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"bookID"`
	IssueDate    string     `gorm:"type:datetime" json:"issueDate"`
	ReturnedFlag bool       `gorm:"type:tinyint" json:"returnedFlag"`
	Penalty      float64    `gorm:"type:double" json:"penalty"`
	StudentID    *uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"studentID"`
}
