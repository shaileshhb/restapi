package bookcontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestAddBook(t *testing.T) {

// 	bookStock := 10

// 	arg := book.Book{
// 		Name:  "Test book",
// 		Stock: &bookStock,
// 	}

// }

func TestGetAllBooks(t *testing.T) {

	tt := []struct {
		name string
		err  string
	}{
		{name: "get all books"},
		// {name: "get all books"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/books", nil)

			wr := httptest.NewRecorder()
			controller.GetAllBooks(wr, req)

			res := wr.Result()
			// defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request but got %v", res.StatusCode)
				}
				return
			}

		})
	}

}
