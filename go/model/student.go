package model

type Student struct {
	Base
	Name     string `json:"name"`
	Age      int    `json:"age"`
	RollNo   int    `json:"rollNo"`
	Date     string `gorm:"type:date" json:"date"`
	DateTime string `gorm:"type:datetime" json:"dateTime"`
	Email    string `json:"email"`
	IsMale   *bool  `json:"isMale"`
}
