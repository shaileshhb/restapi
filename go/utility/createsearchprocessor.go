package utility

import (
	"log"
	"strings"

	"github.com/shaileshhb/restapi/repository"
)

func CreateSearchProcessor(params map[string][]string, queryProcessors *[]repository.QueryProcessor) {

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

			query = ""
			for i := 0; i < len(bookID); i++ {
				if i == 0 {
					query = "book_id = " + "'" + bookID[i] + "'"
					continue
				}
				query += " OR book_id = " + "'" + bookID[i] + "'"
			}
			log.Println("Book IDs -> ", bookID)
			log.Println("Query -> ", query)
			*queryProcessors = append(*queryProcessors, repository.Where(query))
		}
	}
}

/*
Search query generated
SELECT * FROM `students`  WHERE `students`.`deleted_at` IS NULL AND ((name LIKE '%m%') AND (age >= '21') AND (email LIKE '%com%'))

Select query with books
SELECT * FROM `students` LEFT JOIN book_issues
on students.id = book_issues.student_id
WHERE `students`.`deleted_at` IS NULL AND ((name LIKE '%a%')
AND (date >= '2020-12-15') AND (date <= '2021-01-21') AND (book_id = '195eb692-9fc3-4bf7-acba-f6e36c17b041' OR book_id = '444a5f39-0477-480a-a991-899968eabb27')))
GROUP BY students.id
*/
