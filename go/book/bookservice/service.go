package bookservice

import (
	"errors"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/repository"
)

type BookService struct {
	repo *repository.GormRepository
	DB   *gorm.DB
}

func NewBookService(repo *repository.GormRepository, db *gorm.DB) *BookService {
	return &BookService{
		repo: repo,
		DB:   db,
	}
}

func (s *BookService) GetAllBooks(books *[]model.Book) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor

	if err := s.repo.Get(uow, books, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

func (s *BookService) GetBook(book *model.Book, id string) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Get(uow, book, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *BookService) AddNewBook(book *model.Book) error {

	if err := s.Validate(book); err != nil {
		return err
	}

	var queryProcessors []repository.QueryProcessor

	checkName := "name = ?"
	queryProcessors = append(queryProcessors, repository.Search(checkName, book.Name, book))

	// create unit of work
	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Add(uow, book, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

// func (s *BookService) Update(book *model.Book, id string) error {

// 	if err := s.Validate(book); err != nil {
// 		return err
// 	}

// 	var queryProcessors []repository.QueryProcessor
// 	checkID := "id = ?"
// 	queryProcessors = append(queryProcessors, repository.Where(checkID, id))

// 	checkName := "name = ?"
// 	queryProcessors = append(queryProcessors, repository.Search(checkName, book.Name, book))

// 	uow := repository.NewUnitOfWork(s.DB, false)

// 	if err := s.repo.Update(uow, book, ?????, queryProcessors); err != nil {
// 		uow.Complete()
// 		return err
// 	}
// 	uow.Commit()

// 	return nil
// }

func (s *BookService) Delete(book *model.Book, id string) error {

	uow := repository.NewUnitOfWork(s.DB, false)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Delete(uow, book, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *BookService) Validate(book *model.Book) error {
	namePattern := regexp.MustCompile("^[a-zA-Z_ ]*$")

	if book.Name == "" || !namePattern.MatchString(book.Name) {
		return errors.New("Name is required")
	}

	if book.Stock == nil {
		return errors.New("Stock should atleast be 1")
	}

	return nil
}
