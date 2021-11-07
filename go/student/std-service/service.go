package stdservice

import (
	"errors"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/model/bookissue"
	"github.com/shaileshhb/restapi/model/student"
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
func (s *Service) GetAll(students *[]student.Student) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.Preload([]string{"BookIssues"}))

	if err := s.repo.Get(uow, students, queryProcessors); err != nil {
		uow.Commit()
		return err
	}
	uow.Commit()

	utility.TrimDates(students)

	return nil
}

// Get returns students as per the id
func (s *Service) Get(students *student.Student, id string) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.Preload([]string{"BookIssues"}))
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Get(uow, students, queryProcessors); err != nil {
		uow.Commit()
		return err
	}
	uow.Commit()

	utility.TrimDate(students)
	return nil
}

func (s *Service) AddNewStudent(student *student.Student) error {

	if err := s.Validate(student); err != nil {
		return err
	}

	var queryProcessors []repository.QueryProcessor
	checkName := "name = ?"
	queryProcessors = append(queryProcessors, repository.Search(checkName, student.Name, student))
	utility.EmptyToNull(student)

	// create unit of work
	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Add(uow, student, queryProcessors); err != nil {
		uow.Commit()
		return err
	}
	uow.Commit()

	return nil

}

func (s *Service) Update(student *student.Student, id string) error {

	if err := s.Validate(student); err != nil {
		return err
	}

	var queryProcessors []repository.QueryProcessor
	checkID := "id = ?"
	queryProcessors = append(queryProcessors, repository.Where(checkID, id))

	uow := repository.NewUnitOfWork(s.DB, false)

	if err := s.repo.Save(uow, student, queryProcessors); err != nil {
		uow.Commit()
		return err
	}
	uow.Commit()

	return nil
}

// Delete the student
func (s *Service) Delete(student *student.Student, id string) error {

	uow := repository.NewUnitOfWork(s.DB, false)
	var totalCount int
	var queryProcessors []repository.QueryProcessor

	query := "student_id=? AND returned_flag=?"
	queryProcessors = append(queryProcessors, repository.Where(query, id, false))

	if err := s.repo.GetCount(uow, bookissue.BookIssue{}, &totalCount, queryProcessors); err != nil {
		uow.Commit()
		return err
	}

	if totalCount > 0 {
		return errors.New("please return all issued books")
	}

	queryProcessors = nil
	queryCondition := "id=?"
	queryProcessors = append(queryProcessors, repository.Where(queryCondition, id))

	if err := s.repo.Delete(uow, student, queryProcessors); err != nil {
		uow.Commit()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) GetSum(students *student.Student, sum *student.Sum) error {

	uow := repository.NewUnitOfWork(s.DB, true)
	var queryProcessors []repository.QueryProcessor

	queryProcessors = append(queryProcessors, repository.Model(&student.Student{}))

	query := "sum(age + roll_no) as n"
	queryProcessors = append(queryProcessors, repository.Select(query))

	if err := s.repo.Scan(uow, sum, queryProcessors); err != nil {
		uow.Commit()
		return err
	}
	uow.Commit()
	return nil

}

func (s *Service) GetDiff(students *student.Student, sum *student.Sum) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.Model(&student.Student{}))
	query := "abs(sum(age - roll_no)) as n"
	queryProcessors = append(queryProcessors, repository.Select(query))

	if err := s.repo.Scan(uow, sum, queryProcessors); err != nil {
		uow.Commit()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) GetDiffOfAgeAndRecord(students *student.Student, sum *student.Sum) error {

	uow := repository.NewUnitOfWork(s.DB, true)

	query := "sum(age) - count(*) as n"
	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.Model(&student.Student{}))

	queryProcessors = append(queryProcessors, repository.Select(query))

	if err := s.repo.Scan(uow, sum, queryProcessors); err != nil {
		uow.Commit()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) Search(students *[]student.Student, params map[string][]string) error {

	uow := repository.NewUnitOfWork(s.DB, true)
	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.Preload([]string{"BookIssues"}))
	s.createSearchQueries(params, &queryProcessors)

	// utility.CreateSearchQuery(params, &queryProcessors)

	if err := s.repo.Get(uow, students, queryProcessors); err != nil {
		uow.Commit()
		return err
	}
	uow.Commit()

	utility.TrimDates(students)
	return nil
}

func (s *Service) createSearchQueries(params map[string][]string, queryProcessors *[]repository.QueryProcessor) {

	var columnNames []string
	var conditions []string
	var values []interface{}
	var operators []string

	if name := params["name"]; len(name) > 0 {
		utility.AddToSlice("name", "LIKE ?", "OR", "%"+name[0]+"%", &columnNames, &conditions, &operators, &values)
	}
	if email := params["email"]; len(email) > 0 {
		utility.AddToSlice("email", "LIKE ?", "OR", "%"+email[0]+"%", &columnNames, &conditions, &operators, &values)
	}
	if age := params["age"]; len(age) > 0 {
		utility.AddToSlice("age", ">= ?", "OR", age[0], &columnNames, &conditions, &operators, &values)
	}
	if dateFrom := params["dateFrom"]; len(dateFrom) > 0 {
		utility.AddToSlice("date", ">= ?", "OR", dateFrom[0], &columnNames, &conditions, &operators, &values)
	}
	if dateTo := params["dateTo"]; len(dateTo) > 0 {
		if len(params["dateFrom"]) > 0 {
			utility.AddToSlice("date", "<= ?", "AND", dateTo[0], &columnNames, &conditions, &operators, &values)
		} else {
			utility.AddToSlice("date", "<= ?", "OR", dateTo[0], &columnNames, &conditions, &operators, &values)
		}
	}
	if books := params["books"]; len(params["books"]) > 0 {
		var query string
		selectQuery := "*"
		*queryProcessors = append(*queryProcessors, repository.Select(selectQuery))

		query = "LEFT JOIN book_issues on students.id = book_issues.student_id"
		*queryProcessors = append(*queryProcessors, repository.Join(query))

		groupBy := "students.id"
		*queryProcessors = append(*queryProcessors, repository.GroupBy([]string{groupBy}))

		query = "returned_flag = false"
		*queryProcessors = append(*queryProcessors, repository.Where(query))

		bookID := strings.Split(books[0], ",")

		utility.AddToSlice("book_id", "IN (?)", "OR", bookID, &columnNames, &conditions, &operators, &values)
	}

	*queryProcessors = append(*queryProcessors, repository.FilterWithOperator(columnNames, conditions, operators, values))

}

func (s *Service) Validate(student *student.Student) error {

	namePattern := regexp.MustCompile(`^[a-zA-Z_ ]*$`)
	// emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$`)
	phonePattern := regexp.MustCompile(`^[0-9]*$`)

	if student.Name == "" || !namePattern.MatchString(student.Name) {
		return errors.New("name is required")
	}

	if student.RollNo != nil && (*student.RollNo) < 0 {
		return errors.New("roll number is invalid")
	}

	if student.Email == "" {
		// || !emailPattern.MatchString(student.Email)
		return errors.New("email is invalid")
	}

	if student.PhoneNumber != nil && len(*student.PhoneNumber) <= 10 && !phonePattern.MatchString(*student.PhoneNumber) {
		return errors.New("phone number should be only numbers and atleast 10 digits")
	}
	return nil
}
