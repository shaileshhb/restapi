package entity

// Book consist of name and stock of a particular book.
type Book struct {
	Base
	Name  string `gorm:"type:varchar(30)" json:"name"`
	Stock *int   `gorm:"type:int" json:"stock"`
}

// BookAvailability will consist of stock details about a book.
type BookAvailability struct {
	Base
	Name       string `json:"name"`
	InStock    *int   `json:"inStock"`
	TotalStock *int   `json:"totalStock"`
}
