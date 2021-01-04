package bookissueservice

import (
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/repository"
)

type IssueService struct {
	repo *repository.GormRepository
	DB   *gorm.DB
}

func NewIssueService(repo *repository.GormRepository, db *gorm.DB) *IssueService {
	return &IssueService{
		repo: repo,
		DB:   db,
	}
}

func (s *IssueService) GetAll(issues *[]model.BookIssue) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor

	if err := s.repo.Get(uow, issues, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

func (s *IssueService) Penalty(penalty *[]model.BookIssueWithPenalty) error {

	return nil
}

func (s *IssueService) Get(bookIssue *model.BookIssue, id string) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "book_id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Get(uow, bookIssue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

func (s *IssueService) AddNewBook(bookIssue *model.BookIssue) error {

	var queryProcessors []repository.QueryProcessor

	// create unit of work
	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Add(uow, bookIssue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

func (s *IssueService) UpdateBook(bookIssue *model.BookIssue, bookID string) error {

	var queryProcessors []repository.QueryProcessor
	var queryCondition = "book_id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, bookID))

	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Update(uow, bookIssue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil

}

func (s *IssueService) Delete(bookIssue *model.BookIssue, id string) error {

	uow := repository.NewUnitOfWork(s.DB, false)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Delete(uow, bookIssue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}
