package service

import (
	"errors"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/repository"
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
	}
	//  else {
	// 	utility.ConvertDateTime(students)
	// 	utility.TrimDateTime(students)
	// }
	uow.Commit()

	return nil

}

// Get returns students as per the id
func (s *Service) Get(students *[]model.Student, id string) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Get(uow, students, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	// else {
	// 	utility.ConvertDateTime(students)
	// 	utility.TrimDateTime(students)
	// }
	uow.Commit()
	return nil
}

func (s *Service) AddNewStudent(student *model.Student) error {

	if err := s.Validate(student); err != nil {
		return err
	}

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

	if err := s.Validate(student); err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(s.DB, false)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

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
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Delete(uow, student, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) Validate(student *model.Student) error {

	emailPattern := regexp.MustCompile("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")
	// datePattern := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	// dateTimePattern := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)

	if student.Name == "" {
		return errors.New("Name is required")
	}

	// if student.RollNo < 0 {
	// 	return errors.New("Roll number is invalid")
	// }

	if student.Email == "" || !emailPattern.MatchString(student.Email) {
		return errors.New("Email is invalid")
	}

	// if student.Age < 18 {
	// 	return errors.New("Age cannot be less than 18")
	// }

	// if !dateTimePattern.MatchString(student.DateTime) {
	// 	return errors.New("Date time is invalid")

	// }

	// if !datePattern.MatchString(student.Date) {
	// 	return errors.New("Date is invalid")
	// }

	// if student.IsMale == nil {
	// 	return errors.New("Gender is required")
	// }

	return nil
}
