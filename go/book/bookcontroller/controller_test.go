package bookcontroller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shaileshhb/restapi/model/book"
)

func TestAddBook(t *testing.T) {

	bookStock1 := 10
	bookStock2 := 0

	tt := []struct {
		name string
		book book.Book
		err  string
	}{
		{name: "add book", book: book.Book{
			Name:  "BookOne",
			Stock: &bookStock1,
		}},
		{name: "zero book", book: book.Book{
			Name:  "BookTwo",
			Stock: &bookStock2,
		}, err: "Stock should atleast be 1"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			// data := url.Values{}
			// data.Set("name", tc.book.Name)
			// data.Set("stock", strconv.Itoa(*tc.book.Stock))

			body, err := json.Marshal(&tc.book)
			if err != nil {
				t.Errorf("err while marshaling data: %v", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
			if err != nil {
				t.Errorf("could not send request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Token", token)

			wr := httptest.NewRecorder()

			controller.AddBook(wr, req)

			res := wr.Result()
			defer res.Body.Close()

			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request but got %v", res.Status)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status %v but got %v", http.StatusOK, res.Status)
			}

		})
	}

}

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
