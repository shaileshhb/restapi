package student

type SearchStudent struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Age      *string `json:"age"`
	DateFrom *string `json:"dateFrom"`
	DateTo   *string `json:"dateTo"`
	BookID   *string `json:"bookID"`
}
