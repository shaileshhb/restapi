package bookcontroller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shaileshhb/restapi/model/book"
	"github.com/shaileshhb/restapi/utility"
	"github.com/stretchr/testify/require"
)

func TestAddBook(t *testing.T) {

	tt := []struct {
		name string
		book book.Book
		err  string
	}{
		{name: "add book", book: book.Book{
			Name:  utility.GenerateRandomString(6),
			Stock: utility.GenerateBookStock(1, 20),
		}},
		{name: "zero book", book: book.Book{
			Name:  utility.GenerateRandomString(6),
			Stock: utility.GenerateBookStock(0, 0),
		}, err: "Stock should atleast be 1"},
		{name: "invalid name", book: book.Book{
			Name:  "Book-3",
			Stock: utility.GenerateBookStock(1, 20),
		}, err: "Name is invalid"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			body, err := json.Marshal(&tc.book)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
			require.NoError(t, err)

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
			require.NotEmpty(t, res.Body)

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
