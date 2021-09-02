package entity

// Student will consist of student details.
type Student struct {
	Base        `structs:"id,omitnested"`
	Name        string      `gorm:"type:varchar(30)" json:"name,omitempty"`
	Age         *int        `gorm:"type:int" json:"age,omitempty"`
	RollNo      *int        `gorm:"type:int" json:"rollNo,omitempty"`
	PhoneNumber *string     `gorm:"type:varchar(10)" json:"phone,omitempty"`
	Email       string      `gorm:"type:varchar(50)" json:"email,omitempty"`
	IsMale      *bool       `gorm:"type:tinyint" json:"isMale,omitempty"`
	Date        *string     `gorm:"type:date" json:"date,omitempty"`
	BookIssues  []BookIssue `json:"bookIssues"`
}
