package model

type Student struct {
	Base
	Name   string `json:"name"`
	Age    int    `json:"age"`
	RollNo int    `json:"rollNo"`
	Date   string `gorm:"type:datetime" json:"date"`
	Email  string `json:"email"`
	IsMale *bool  `json:"isMale"`
}
