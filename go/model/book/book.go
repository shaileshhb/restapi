package book

import (
	"encoding/json"
	"io"

	"github.com/shaileshhb/restapi/model/general"
)

type Book struct {
	general.Base
	Name  string `gorm:"type:varchar(30)" json:"name"`
	Stock *int   `gorm:"type:int" json:"stock"`
}

type BookAvailability struct {
	general.Base
	Name       string `json:"name"`
	InStock    *int   `json:"inStock"`
	TotalStock *int   `json:"totalStock"`
}

// FromJSON will decode json and return error
func (b *Book) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}
