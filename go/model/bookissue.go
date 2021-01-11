package model

import uuid "github.com/satori/go.uuid"

type BookIssue struct {
	Base
	BookID       uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"bookID"`
	IssueDate    string    `gorm:"type:datetime" json:"issueDate"`
	ReturnedFlag bool      `gorm:"type:tinyint" json:"returnedFlag"`
	Penalty      float64   `gorm:"type:double" json:"penalty"`
	StudentID    uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"studentID"`
}

// type BookIssueWithPenalty struct {
// 	Base
// 	BookID       uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"bookID"`
// 	IssueDate    string    `gorm:"type:datetime" json:"issueDate"`
// 	ReturnedFlag *bool     `gorm:"type:tinyint" json:"returnedFlag"`
// 	Penalty      float64   `gorm:"type:double" json:"penalty"`
// 	StudentID    uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"studentID"`
// }
