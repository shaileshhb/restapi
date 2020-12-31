package model

type Book struct {
	Base
	Name  string `gorm:"type:varchar(30)" json:"name"`
	Stock *int   `gorm:"type:int" json:"stock"`
}
