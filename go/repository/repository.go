package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	Get(uow *UnitOfWork, out interface{}) error
	GetCount(uow *UnitOfWork, out interface{}, count *int, queryProcessors []QueryProcessor) error
	Add(uow *UnitOfWork, entity interface{}) error
	Update(uow *UnitOfWork, entity interface{}, entityMap map[string]interface{}, queryProcessors []QueryProcessor) error
	Delete(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
}

type GormRepository struct{}

func NewGormRepository() *GormRepository {
	return &GormRepository{}
}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

func Where(condition string, value interface{}) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if value != "" {
			db = db.Debug().Model(out).Where(condition, value)
		}
		return db, nil
	}
}

// Filter will filter the results
func Filter(condition string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Where(condition, args...)
		return db, nil
	}
}

func Search(condition string, value string, entity interface{}) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if value != "" {
			if !db.Debug().Where(condition, value).First(entity).RecordNotFound() {
				return nil, errors.New("Same name already exists")
			}
		}
		return db, nil
	}
}

func Model() QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Model(out)

		return db, nil
	}
}

func Preload(preloadAssociation []string) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if preloadAssociation != nil {
			for _, association := range preloadAssociation {
				db = db.Debug().Preload(association)
			}
		}
		return db, nil
	}
}

func GroupBy(groupBy []string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		for _, entity := range groupBy {
			db = db.Group(entity)
		}
		return db, nil
	}
}

func Select(query string) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Select(query)
		return db, nil
	}
}

func Join(query string) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Joins(query)
		return db, nil
	}
}

func (g *GormRepository) Get(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {

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

func (g *GormRepository) SelectQuery(uow *UnitOfWork, entity interface{}, out interface{}, query string) error {

	db := uow.DB
	var err error

	if err = db.Debug().Model(entity).Select(query).Scan(out).Error; err != nil {
		return err
	}

	return nil
}

func (g *GormRepository) Scan(uow *UnitOfWork, entity interface{}, out interface{}, queryProcessors []QueryProcessor) error {

	db := uow.DB
	var err error

	if queryProcessors != nil {
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
		}
	}

	if err = db.Debug().Scan(out).Error; err != nil {
		return err
	}
	return nil

}
