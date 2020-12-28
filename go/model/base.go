package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key" json:"id" structs:"id"`
	CreatedAt time.Time  `json:"-" structs:"-"`
	UpdatedAt time.Time  `json:"-" structs:"-"`
	DeletedAt *time.Time `sql:"index" json:"-" structs:"-"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)
}
