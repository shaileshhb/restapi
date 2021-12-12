package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/restapi/errors"
)

// UnmarshalJSON parses data from request and return otherwise error return.
func UnmarshalJSON(request *http.Request, out interface{}) error {
	if request.Body == nil {
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.NewHTTPError(errors.ErrorCodeReadWriteFailure, http.StatusBadRequest)
	}

	if len(body) == 0 {
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return errors.NewHTTPError(errors.ErrorCodeInvalidJSON, http.StatusBadRequest)
	}
	return nil
}

// GetLimitAndOffset gets and returns the limit and offset from the given request.
func GetLimitAndOffset(r *http.Request) (limit int, offset int) {
	queryparam := mux.Vars(r)
	limitparam := queryparam["limit"]
	offsetparam := queryparam["offset"]
	var err error
	limit = 5
	if len(limitparam) > 0 {
		limit, err = strconv.Atoi(limitparam)
		if err != nil {
			limit = 5
		}
	}
	offset = 0
	if len(offsetparam) > 0 {
		offset, err = strconv.Atoi(offsetparam)
		if err != nil {
			offset = 0
		}
	}
	return limit, offset
}
