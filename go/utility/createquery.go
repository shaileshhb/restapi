package utility

import (
	"log"
	"strings"

	"github.com/shaileshhb/restapi/repository"
)

func CreateSearchQuery(params map[string][]string, queryProcessors *[]repository.QueryProcessor) {

	var query string

	for key, value := range params {

		log.Println("KEY -> ", key, "VALUE -> ", value, "length ->", len(value))

		if key == "age" && value[0] != "" {
			query = key + " >= ?"
			*queryProcessors = append(*queryProcessors, repository.Where(query, value[0]))
			continue
		}

		if key == "dateFrom" && value[0] != "" {
			query = "date >= ?"
			*queryProcessors = append(*queryProcessors, repository.Where(query, value[0]))
			continue
		}

		if key == "dateTo" && value[0] != "" {
			query = "date <= ?"
			*queryProcessors = append(*queryProcessors, repository.Where(query, value[0]))
			continue
		}

		if key == "name" || key == "email" {
			query = key + " LIKE ?"
			queryValue := "%" + value[0] + "%"
			*queryProcessors = append(*queryProcessors, repository.Where(query, queryValue))
			continue
		}

		if key == "books" {
			selectQuery := "*"
			*queryProcessors = append(*queryProcessors, repository.Select(selectQuery))

			query = "LEFT JOIN book_issues on students.id = book_issues.student_id"
			*queryProcessors = append(*queryProcessors, repository.Join(query))

			groupBy := "students.id"
			*queryProcessors = append(*queryProcessors, repository.GroupBy([]string{groupBy}))

			query = "returned_flag = false"
			*queryProcessors = append(*queryProcessors, repository.Where(query))

			log.Println("value[0] -> ", value[0])
			bookID := strings.Split(value[0], ",")

			query = "book_id IN (?)"

			log.Println("Book IDs -> ", bookID)
			log.Println("Query -> ", query)

			*queryProcessors = append(*queryProcessors, repository.Where(query, bookID))
		}
	}
}

/*
Search query generated
SELECT * FROM `students`  WHERE `students`.`deleted_at` IS NULL AND ((name LIKE '%m%') AND (age >= '21') AND (email LIKE '%com%'))

Select query with books
SELECT * FROM `students` LEFT JOIN book_issues on students.id = book_issues.student_id
WHERE `students`.`deleted_at` IS NULL AND ((returned_flag = false) AND
(book_id in ( '4b0aa6eb-e851-4cf5-8cb6-837d2a63c0dd', '2ae5f77c-96aa-4191-8f48-9fc6cd4a6083' )))
GROUP BY students.id
*/

func AddToSlice(columnName string, condition string, operator string, value interface{},
	columnNames *[]string, conditions *[]string, operators *[]string, values *[]interface{}) {

	if len(*columnNames) != 0 {
		*operators = append(*operators, operator)
	}
	*columnNames = append(*columnNames, columnName)
	*conditions = append(*conditions, condition)
	*values = append(*values, value)
}
