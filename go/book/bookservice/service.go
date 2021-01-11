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

func (s *BookService) GetAllBooks(books *[]model.BookAvailability) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var book = model.Book{}
	var queryProcessors []repository.QueryProcessor

	queryProcessors = append(queryProcessors, repository.Model(book))

	selectQuery := "books.id as id, books.name as name, if(sum(returned_flag=0)>0, abs(stock - sum(returned_flag=0)), stock) as in_stock, books.stock as total_stock"
	// selectQuery := "books.id as id, books.name as name, abs(stock - sum(returned_flag=0)) as in_stock, books.stock as total_stock"

	queryProcessors = append(queryProcessors, repository.Select(selectQuery))

	joinQuery := "left join book_issues on books.id = book_issues.book_id"
	queryProcessors = append(queryProcessors, repository.Join(joinQuery))

	groupBy := "books.id"
	queryProcessors = append(queryProcessors, repository.GroupBy([]string{groupBy}))

	if err := s.repo.Scan(uow, books, queryProcessors); err != nil {
		uow.Complete()
		return err
	}

	uow.Commit()
	return nil

	// s.DB.Debug().Model(books).Select("books.id as id, books.name as name, if(returned_flag = false, abs(stock - count(book_id)), stock) as total_stock, books.stock as stock, book_issues.returned_flag as returned_flag").Joins("left join book_issues on books.id = book_issues.book_id").Group("books.id").Scan(joinBookIssue)
	// // SELECT books.id as id, books.name as name, books.stock as total_stock, books.stock as stock, book_issues.returned_flag as returned_flag FROM `books` inner join book_issues on books.id = book_issues.book_id WHERE `books`.`deleted_at` IS NULL GROUP BY books.id

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

	// create unit of work
	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.Validate(book); err != nil {
		return err
	}

	var queryProcessors []repository.QueryProcessor

	checkName := "name = ?"
	queryProcessors = append(queryProcessors, repository.Search(checkName, book.Name, book))

	if err := s.repo.Add(uow, book, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

func (s *BookService) Update(book *model.Book, id string) error {

	if err := s.Validate(book); err != nil {
		return err
	}

	var queryProcessors []repository.QueryProcessor
	checkID := "id = ?"
	queryProcessors = append(queryProcessors, repository.Where(checkID, id))

	// checkName := "name = ?"
	// queryProcessors = append(queryProcessors, repository.Search(checkName, book.Name, book))

	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Update(uow, book, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

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

	if book.Name == "" {
		return errors.New("Name is required")
	}

	if !namePattern.MatchString(book.Name) {
		return errors.New("Name is invalid")
	}

	if book.Stock == nil {
		return errors.New("Stock should atleast be 1")
	}

	return nil
}
