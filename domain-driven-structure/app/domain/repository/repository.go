package repository

import "github.com/jinzhu/gorm"

type Repository interface {
	Get(uow *UnitOfWork, out interface{}) error
	GetCount(uow *UnitOfWork, out interface{}, count *int, queryProcessors []QueryProcessor) error
	Add(uow *UnitOfWork, entity interface{}) error
	Update(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
	Save(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
	Delete(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
}

type GormRepository struct{}

func NewGormRepository() *GormRepository {
	return &GormRepository{}
}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)
