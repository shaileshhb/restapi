package service

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/repository"
	"github.com/shaileshhb/restapi/utility"
)

type Service struct {
	repo *repository.GormRepository
	DB   *gorm.DB
}

func NewService(repo *repository.GormRepository, db *gorm.DB) *Service {
	return &Service{
		repo: repo,
		DB:   db,
	}
}

// GetAll gives all students
func (s *Service) GetAll(students *[]model.Student) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor

	if err := s.repo.Get(uow, students, queryProcessors); err != nil {
		uow.Complete()
		return err
	} else {
		utility.ConvertDateTime(students)
	}
	uow.Commit()

	return nil

}

// Get returns students as per the id
func (s *Service) Get(students *[]model.Student, id string) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.GetStudentByID(id))

	if err := s.repo.Get(uow, students, queryProcessors); err != nil {
		uow.Complete()
		return err
	} else {
		utility.ConvertDateTime(students)
	}
	uow.Commit()
	return nil
}

func (s *Service) AddNewStudent(student *model.Student) error {

	// create unit of work
	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Add(uow, student); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	return nil

}

func (s *Service) Update(student *model.Student, id string) error {

	uow := repository.NewUnitOfWork(s.DB, false)

	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.GetStudentByID(id))

	if err := s.repo.Update(uow, student, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

// Delete the student
func (s *Service) Delete(student *model.Student, id string) error {

	uow := repository.NewUnitOfWork(s.DB, false)

	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.GetStudentByID(id))

	if err := s.repo.Delete(uow, student, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) ValidateJsonFields(student model.Student) error {

	if err := s.ValidateStringValues([]string{student.Name, student.Date, student.Email}); err != nil {
		return err
	}
	return nil
}

func (s *Service) ValidateStringValues(fields []string) error {

	for _, field := range fields {
		for _, str := range field {
			if (str < 'a' || str > 'z') && (str < 'A' || str > 'Z') {
				return errors.New("Invalid String")
			}
		}

	}
	return nil
}
