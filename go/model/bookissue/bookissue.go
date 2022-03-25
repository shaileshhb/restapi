package bookissue

import (
	"time"

	uuid "github.com/satori/uuid"
	"github.com/shaileshhb/restapi/customtime"
	"github.com/shaileshhb/restapi/model/general"
)

type BookIssue struct {
	general.Base
	Date         customtime.CustomTime `gorm:"type:date" json:"date" time:"2006-01-02"`
	BookID       *uuid.UUID            `gorm:"type:varchar(36); foreign_key" json:"bookID"`
	Time         time.Time             `gorm:"type:datetime" json:"time"`
	ReturnedFlag bool                  `gorm:"type:tinyint" json:"returnedFlag"`
	Penalty      float64               `gorm:"type:double" json:"penalty"`
	StudentID    *uuid.UUID            `gorm:"type:varchar(36); foreign_key" json:"studentID"`
	IssueDate    string                `gorm:"type:datetime" json:"issueDate"`
	// IssueDate    customtime.CustomTime `gorm:"type:date" json:"issueDate"`
}

// type BookIssueWithPenalty struct {
// 	Base
// 	BookID       uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"bookID"`
// 	IssueDate    string    `gorm:"type:datetime" json:"issueDate"`
// 	ReturnedFlag *bool     `gorm:"type:tinyint" json:"returnedFlag"`
// 	Penalty      float64   `gorm:"type:double" json:"penalty"`
// 	StudentID    uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"studentID"`
// }
