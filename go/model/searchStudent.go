package model

type SearchStudent struct {
	name     *string `json: name`
	email    *string `json: email`
	age      *string `json: age`
	dateFrom *string `json: dateFrom`
	dateTo   *string `json: dateTo`
	bookID   *string `json: bookID`
}
