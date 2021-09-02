package bookissue

import (
	uuid "github.com/satori/uuid"
	"github.com/shaileshhb/restapi/model/general"
)

type BookIssue struct {
	general.Base
	BookID       *uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"bookID"`
	IssueDate    string     `gorm:"type:datetime" json:"issueDate"`
	ReturnedFlag bool       `gorm:"type:tinyint" json:"returnedFlag"`
	Penalty      float64    `gorm:"type:double" json:"penalty"`
	StudentID    *uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"studentID"`
}

// type BookIssueWithPenalty struct {
// 	Base
// 	BookID       uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"bookID"`
// 	IssueDate    string    `gorm:"type:datetime" json:"issueDate"`
// 	ReturnedFlag *bool     `gorm:"type:tinyint" json:"returnedFlag"`
// 	Penalty      float64   `gorm:"type:double" json:"penalty"`
// 	StudentID    uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"studentID"`
// }
