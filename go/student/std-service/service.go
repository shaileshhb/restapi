package stdservice

import (
	"errors"
	"log"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/model"
	"github.com/shaileshhb/restapi/repository"
	"github.com/shaileshhb/restapi/utility"
	"github.com/shaileshhb/restapi/utility/calculatepenalty"
	"github.com/shaileshhb/restapi/utility/structToMap"
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
	queryProcessors = append(queryProcessors, repository.Preload([]string{"BookIssues"}))

	if err := s.repo.Get(uow, students, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	utility.TrimDates(students)
	// utility.TrimDateTime(students)

	return nil

}

// Get returns students as per the id
func (s *Service) Get(students *model.Student, id string) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.Preload([]string{"BookIssues"}))
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Get(uow, students, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()

	utility.TrimDate(students)
	s.UpdatePenalty(&students.BookIssues)
	// utility.TrimDateTime(students)

	return nil
}

func (s *Service) UpdatePenalty(bookIssues *[]model.BookIssue) error {

	log.Println("BookID From UPDATE PENALTY")

	uow := repository.NewUnitOfWork(s.DB, false)
	var queryProcessors []repository.QueryProcessor

	calculatepenalty.Penalty(bookIssues)

	if err := s.repo.Update(uow, bookIssues, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) AddNewStudent(student *model.Student) error {

	if err := s.Validate(student); err != nil {
		return err
	}

	var queryProcessors []repository.QueryProcessor

	checkName := "name = ?"
	queryProcessors = append(queryProcessors, repository.Search(checkName, student.Name, student))

	structToMap.EmptyToNull(student)

	// create unit of work
	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Add(uow, student, queryProcessors); err != nil {
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

	log.Println("Book Issues -> ", student.BookIssues)

	// studentMap := structToMap.ConvertStructToMap(student, id)

	var queryProcessors []repository.QueryProcessor
	checkID := "id = ?"
	queryProcessors = append(queryProcessors, repository.Where(checkID, id))

	checkName := "name = ?"
	queryProcessors = append(queryProcessors, repository.Search(checkName, student.Name, student))

	// bookAlreadyIssuedSearch := ""

	uow := repository.NewUnitOfWork(s.DB, false)

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

	namePattern := regexp.MustCompile("^[a-zA-Z_ ]*$")
	emailPattern := regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")
	phonePattern := regexp.MustCompile("^[0-9]*$")
	// datePattern := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	// dateTimePattern := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`)

	if student.Name == "" || !namePattern.MatchString(student.Name) {
		return errors.New("Name is required")
	}

	if student.RollNo != nil && (*student.RollNo) < 0 {
		return errors.New("Roll number is invalid")
	}

	if student.Email == "" || !emailPattern.MatchString(student.Email) {
		return errors.New("Email is invalid")
	}

	if student.PhoneNumber != nil && len(*student.PhoneNumber) <= 10 && !phonePattern.MatchString(*student.PhoneNumber) {
		// log.Println(len(*student.PhoneNumber) >= 10, phonePattern.MatchString(*student.PhoneNumber))
		return errors.New("Phone number should be only numbers and atleast 10 digits")
	}

	// if student.DateTime != nil && !dateTimePattern.MatchString((*student.DateTime)) {
	// 	return errors.New("Date time is invalid")

	// }

	// if student.Date != nil && !datePattern.MatchString((*student.Date)) {
	// 	return errors.New("Date is invalid")

	// }

	// if student.IsMale == nil {
	// 	return errors.New("Gender is required")
	// }

	return nil
}

func (s *Service) GetSum(students *model.Student, sum *model.Sum) error {

	// if err := s.repo.GetSum()
	uow := repository.NewUnitOfWork(s.DB, true)

	query := "sum(age + roll_no) as n"

	err := s.repo.SelectQuery(uow, students, sum, query)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) GetDiff(students *model.Student, sum *model.Sum) error {

	// if err := s.repo.GetSum()
	uow := repository.NewUnitOfWork(s.DB, true)

	query := "abs(sum(age - roll_no)) as n"

	err := s.repo.SelectQuery(uow, students, sum, query)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) GetDiffOfAgeAndRecord(students *model.Student, sum *model.Sum) error {

	// if err := s.repo.GetSum()
	uow := repository.NewUnitOfWork(s.DB, true)

	query := "sum(age) - count(*) as n"

	err := s.repo.SelectQuery(uow, students, sum, query)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) GetAge(students *[]model.Student) error {

	// if err := s.repo.GetSum()
	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryCondition := "age > ?"
	minAge := 18
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, minAge))

	if err := s.repo.Get(uow, students, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}
