package bookissueservice

import (
	"log"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/repository"
	"github.com/shaileshhb/restapi/utility/validateBookIssue"
	"github.com/shaileshhb/restapi/utility/calculatepenalty"

)

type BookIssueService struct {
	repo *repository.GormRepository
	DB   *gorm.DB
}

func NewIssueService(repo *repository.GormRepository, db *gorm.DB) *BookIssueService {
	return &BookIssueService{
		repo: repo,
		DB:   db,
	}
}

// func (s *BookIssueService) GetAll(bookIssue *[]model.BookIssue) error {

// 	uow := repository.NewUnitOfWork(s.DB, true)

// 	var queryProcessors []repository.QueryProcessor

// 	if err := s.repo.Get(uow, bookIssue, queryProcessors); err != nil {
// 		uow.Complete()
// 		return err
// 	}
// 	uow.Commit()

// 	return nil
// }

func (s *BookIssueService) GetAll(bookIssues *[]model.BookIssue, id string) error {

	log.Println("Get all called, id ->", id)

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "student_id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Get(uow, bookIssues, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	log.Println("BookIssues->", bookIssues == nil, len(*bookIssues))
	if len(*bookIssues) != 0 {
		s.UpdatePenalty(bookIssues)
	}
	return nil
}

func (s *BookIssueService) UpdatePenalty(bookIssues *[]model.BookIssue) error {

	// log.Println("BookID From UPDATE PENALTY")

	var queryProcessors []repository.QueryProcessor

	log.Println("Book issues Before ->", (*bookIssues)[0].Penalty)
	err := calculatepenalty.Penalty(bookIssues)
	if err != nil {
		return err
	}
	log.Println("Book issues After ->", (*bookIssues)[0].Penalty)
	log.Println("Issue Date->", (*bookIssues)[0].IssueDate)

	for i := 0; i < len((*bookIssues)); i++ {
		uow := repository.NewUnitOfWork(s.DB, false)
		queryProcessors = nil
		queryProcessors = append(queryProcessors, repository.Where("student_id=?", (*bookIssues)[i].StudentID))

		if err := s.repo.Update(uow, (*bookIssues)[i], queryProcessors); err != nil {
			uow.Complete()
			return err
		}
		uow.Commit()

	}

	log.Println("Book penalty after final update ->", (*bookIssues)[0].Penalty)
	return nil
}

func (s *BookIssueService) AddNewBookIssue(bookIssue *model.BookIssue) error {


	uow := repository.NewUnitOfWork(s.DB, false)
	var queryProcessors []repository.QueryProcessor

	var issue = []model.BookIssue{}
	condition := "student_id=?"
	queryProcessors = append(queryProcessors, repository.Where(condition, bookIssue.StudentID))

	if err := s.repo.Get(uow, &issue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}

	if err := validateBookIssue.ValidateBookIssue(bookIssue, issue); err != nil {
		uow.Complete()
		return err
	}

	if err := s.repo.Add(uow, bookIssue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil
}

func (s *BookIssueService) UpdateBook(bookIssue *model.BookIssue, bookID string) error {

	var queryProcessors []repository.QueryProcessor
	
	var queryBookID = "book_id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryBookID, bookID))

	var queryStudentID = "student_id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryStudentID, bookIssue.StudentID))

	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Update(uow, bookIssue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil

}

func (s *BookIssueService) Delete(bookIssue *model.BookIssue, id string) error {

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

func (s *BookIssueService) Validate(bookIssue *model.BookIssue) error {

	if bookIssue.BookID.String() == "" {
		return errors.New("Book ID is required")
	}

	if bookIssue.StudentID.String() == "" {
		return errors.New("Student ID is required")
	}

	return nil
}