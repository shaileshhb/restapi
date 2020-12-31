package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	Get(uow *UnitOfWork, out interface{}) error
	Add(uow *UnitOfWork, entity interface{}) error
	Update(uow *UnitOfWork, entity interface{}, entityMap map[string]interface{}, queryProcessors []QueryProcessor) error
	Delete(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
}

type GormRepository struct{}

func NewGormRepository() *GormRepository {
	return &GormRepository{}
}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

func Where(condition string, value string) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if value != "" {
			db = db.Model(out).Where(condition, value)
		}
		return db, nil
	}
}

func Search(condition string, value string, entity interface{}) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		// student := &model.Student{}
		if value != "" {
			if !db.Debug().Where(condition, value).First(entity).RecordNotFound() {
				return nil, errors.New("Same name already exists")
			}
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

func (g *GormRepository) Add(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error {

	db := uow.DB
	var err error

	log.Println("Inside Repo Add", queryProcessors == nil)

	if queryProcessors != nil {
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
			if err != nil {
				return err
			}
		}
	}

	if err = db.Debug().Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) Update(uow *UnitOfWork, entity interface{}, entityMap map[string]interface{}, queryProcessors []QueryProcessor) error {

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

	if err := db.Debug().Updates(entityMap).Error; err != nil {
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

func (g *GormRepository) GetSum(uow *UnitOfWork, entity interface{}) (int64, error) {

	db := uow.DB
	var err error
	var result int64

	row := db.Debug().Model(entity).Select("sum(age+roll_no)").Row()
	if err = row.Scan(&result); err != nil {
		return result, nil
	}
	// db.Debug().Raw("SELECT SUM(age) FROM students").Scan(&age)

	return result, nil
}
