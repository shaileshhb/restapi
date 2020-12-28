package service

import (
	"errors"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/repository"
)

type UserService struct {
	repo *repository.GormRepository
	DB   *gorm.DB
}

func NewUserService(repo *repository.GormRepository, db *gorm.DB) *UserService {
	return &UserService{
		repo: repo,
		DB:   db,
	}
}

func (s *UserService) Get(user *model.User, username string) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "username=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, username))

	if err := s.repo.Get(uow, user, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil

}

func (s *UserService) Add(user *model.User) error {

	if err := s.Validate(user); err != nil {
		return err
	}

	var queryProcessors []repository.QueryProcessor

	checkName := "username = ?"
	queryProcessors = append(queryProcessors, repository.Search(checkName, user.Username))

	// create unit of work
	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Add(uow, user, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil

}

// func (s *UserService) Update(user *model.User, id string) error {

// if err := s.Validate(user); err != nil {
// 	return err
// }

// 	var queryProcessors []repository.QueryProcessor
// 	queryCondition := "id=?"
// 	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

// 	uow := repository.NewUnitOfWork(s.DB, false)

// 	if err := s.repo.Update(uow, user, queryProcessors); err != nil {
// 		uow.Complete()
// 		return err
// 	}
// 	uow.Commit()

// 	return nil
// }

// Delete the student
// func (s *UserService) Delete(user *model.User, id string) error {

// 	uow := repository.NewUnitOfWork(s.DB, false)

// 	var queryProcessors []repository.QueryProcessor
// 	queryCondition := "id=?"
// 	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

// 	if err := s.repo.Delete(uow, user, queryProcessors); err != nil {
// 		uow.Complete()
// 		return err
// 	}
// 	uow.Commit()
// 	return nil
// }

func (s *UserService) Validate(user *model.User) error {

	emailPattern := regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")

	if user.Username == "" {
		return errors.New("Username is required")
	}

	if user.Password == "" {
		return errors.New("Password is required")
	}

	if user.Email == "" || !emailPattern.MatchString(user.Email) {
		return errors.New("Email is invalid")
	}

	return nil
}
