package utility

import (
	"log"

	"github.com/shaileshhb/restapi/repository"
)

func CreateQueryProcessor(params map[string][]string, queryProcessors *[]repository.QueryProcessor) {

	var query string

	for key, value := range params {

		log.Println("KEY -> ", key, "VALUE -> ", value, "length ->", len(value))

		if key == "age" && value[0] != "" {
			query = key + " >= ?"
			*queryProcessors = append(*queryProcessors, repository.Where(query, value[0]))
			continue
		}

		if key == "start" && value[0] != "" {
			query = "date >= ?"
			*queryProcessors = append(*queryProcessors, repository.Where(query, value[0]))
			continue
		}

		if key == "end" && value[0] != "" {
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
	}
}

/*
Search query generated
SELECT * FROM `students`  WHERE `students`.`deleted_at` IS NULL AND ((name LIKE '%m%') AND (age >= '21') AND (email LIKE '%com%'))
*/
