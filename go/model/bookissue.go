package model

import uuid "github.com/satori/go.uuid"

type BookIssue struct {
	BookID       uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"bookID"`
	IssueDate    *string   `gorm:"type:date" json:"issueDate"`
	ReturnedFlag *bool     `gorm:"type:tinyint" json:"returnedFlag"`
	StudentID    uuid.UUID `gorm:"type:varchar(36); foreign_key" json:"studentID"`
}
