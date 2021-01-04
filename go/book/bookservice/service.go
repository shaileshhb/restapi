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

	queryProcessors = append(queryProcessors, repository.Model())

	selectQuery := "books.id as id, books.name as name, if(returned_flag = false, abs(stock - count(book_id)), stock) as in_stock, books.stock as total_stock"
	queryProcessors = append(queryProcessors, repository.Select(selectQuery))

	joinQuery := "left join book_issues on books.id = book_issues.book_id"
	queryProcessors = append(queryProcessors, repository.Join(joinQuery))

	groupBy := "books.id"
	queryProcessors = append(queryProcessors, repository.GroupBy([]string{groupBy}))

	if err := s.repo.Scan(uow, &book, books, queryProcessors); err != nil {
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

// func (s *BookService) GetTotalBooksAvailable(joinBookIssue *[]model.BookAvailability) error {

// 	// s.DB.Debug().Model(book).Select("*").Joins("left join book_issues on id = book_id").Scan(joinBookIssue)
// 	// SELECT * FROM `books` left join book_issues on id = book_id WHERE `books`.`deleted_at` IS NULL

// 	uow := repository.NewUnitOfWork(s.DB, true)

// 	var book = model.Book{}
// 	var queryProcessors []repository.QueryProcessor

// 	queryProcessors = append(queryProcessors, repository.Model())

// 	selectQuery := "books.id as id, books.name as name, if(returned_flag = false, abs(stock - count(book_id)), stock) as total_stock, books.stock as stock, book_issues.returned_flag as returned_flag"
// 	queryProcessors = append(queryProcessors, repository.Select(selectQuery))

// 	joinQuery := "inner join book_issues on books.id = book_issues.book_id"
// 	queryProcessors = append(queryProcessors, repository.Join(joinQuery))

// 	groupBy := "books.id"
// 	queryProcessors = append(queryProcessors, repository.GroupBy([]string{groupBy}))

// 	if err := s.repo.Scan(uow, &book, joinBookIssue, queryProcessors); err != nil {
// 		uow.Complete()
// 		return err
// 	}

// 	uow.Commit()
// 	return nil

// 	// s.DB.Debug().Model(books).Select("books.id as id, books.name as name, if(returned_flag = false, abs(stock - count(book_id)), stock) as total_stock, books.stock as stock, book_issues.returned_flag as returned_flag").Joins("left join book_issues on books.id = book_issues.book_id").Group("books.id").Scan(joinBookIssue)
// 	// // SELECT books.id as id, books.name as name, books.stock as total_stock, books.stock as stock, book_issues.returned_flag as returned_flag FROM `books` inner join book_issues on books.id = book_issues.book_id WHERE `books`.`deleted_at` IS NULL GROUP BY books.id
// 	// return nil

// }

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

	if book.Name == "" || !namePattern.MatchString(book.Name) {
		return errors.New("Name is required")
	}

	if book.Stock == nil {
		return errors.New("Stock should atleast be 1")
	}

	return nil
}
