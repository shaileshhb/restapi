package repository

import "github.com/jinzhu/gorm"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

type Repository interface {
	Get(uow *UnitOfWork, out interface{}) error
	Add(uow *UnitOfWork, entity interface{}) error
	Update(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
	Delete(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

func Where(condition string, entity string) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if entity != "" {
			db = db.Model(out).Where(condition, entity)
		}
		return db, nil
	}
}

func (g *UserRepository) Get(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {

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
	if err = db.Debug().First(out).Error; err != nil {
		return err
	}

	return nil
}

func (g *UserRepository) Add(uow *UnitOfWork, entity interface{}) error {

	db := uow.DB

	if err := db.Debug().Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// func (g *UserRepository) Update(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error {

// 	db := uow.DB
// 	var err error

// 	if queryProcessors != nil {
// 		for _, queryProcessor := range queryProcessors {
// 			db, err = queryProcessor(db, entity)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	if err := db.Debug().Updates(entity).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (g *UserRepository) Delete(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error {

// 	db := uow.DB
// 	var err error

// 	if queryProcessors != nil {
// 		for _, queryProcessor := range queryProcessors {
// 			db, err = queryProcessor(db, entity)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	if err := db.Debug().Delete(entity).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
