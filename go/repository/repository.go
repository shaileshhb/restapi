package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/model"
)

type Repository interface {
	Get(uow *UnitOfWork, out interface{}) error
	AddCustomer(uow *UnitOfWork, cust model.Student) error
	UpdateCustomer(uow *UnitOfWork, cust model.Student, newCust model.Student) error
	DeleteCustomer(uow *UnitOfWork, cust model.Student) error
}

type GormRepository struct{}

func NewGormRepository() *GormRepository {
	return &GormRepository{}
}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

func Where(condition string, id string) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if id != "" {
			db = db.Model(out).Where("id=?", id)
		}
		return db, nil
	}
}

func (*GormRepository) Get(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {

	db := uow.DB
	var err error

	if queryProcessors != nil {
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return err
			}
		}
	}
	if err = db.Debug().Find(out).Error; err != nil {
		return err
	}

	return nil
}

func (g *GormRepository) Add(uow *UnitOfWork, entity interface{}) error {

	db := uow.DB

	if err := db.Debug().Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) Delete(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error {

	db := uow.DB
	var err error

	if queryProcessors != nil {
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
			if err != nil {
				return err
			}
		}
	}
	if err := db.Debug().Delete(entity).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) Update(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error {

	db := uow.DB
	var err error

	if queryProcessors != nil {
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
			if err != nil {
				return err
			}
		}
	}
	if err := db.Debug().Model(entity).Update(entity).Error; err != nil {
		return err
	}
	return nil
}
