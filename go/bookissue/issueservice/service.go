package issueservice

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

func (s *IssueService) Get(issue *model.BookIssue, id string) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Get(uow, issue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

func (s *IssueService) AddNewBook(issue *model.BookIssue) error {

	var queryProcessors []repository.QueryProcessor

	// create unit of work
	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Add(uow, issue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

func (s *IssueService) Delete(issue *model.BookIssue, id string) error {

	uow := repository.NewUnitOfWork(s.DB, false)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Delete(uow, issue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}
