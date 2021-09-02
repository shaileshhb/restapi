package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
)

// Base will parent of all the models which consist of master fields.
type Base struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key" json:"id" structs:"id"`
	CreatedAt time.Time  `json:"-" structs:"-"`
	UpdatedAt time.Time  `json:"-" structs:"-"`
	DeletedAt *time.Time `sql:"index" json:"-" structs:"-"`
}

// BeforeCreate will assign UUID before adding new record to table.
func (base *Base) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)
}
