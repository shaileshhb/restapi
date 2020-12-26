package repository

import "github.com/jinzhu/gorm"

type UnitOfWork struct {
	DB        *gorm.DB
	Committed bool
	ReadOnly  bool
}

func NewUnitOfWork(db *gorm.DB, readOnly bool) *UnitOfWork {

	if readOnly {
		return &UnitOfWork{
			DB:        db.New(),
			Committed: false,
			ReadOnly:  readOnly,
		}
	}
	return &UnitOfWork{
		DB:        db.New().Begin(),
		Committed: false,
		ReadOnly:  readOnly,
	}
}

func (uow *UnitOfWork) Commit() {
	if !uow.ReadOnly {
		uow.DB.Commit()
	}

	uow.Committed = true
}

func (uow *UnitOfWork) Complete() {
	if !uow.Committed && !uow.ReadOnly {
		uow.DB.Rollback()
	}
}
