package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shaileshhb/restapi/errors"
)

// RespondJSON Make response with json formate.
func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, error := json.Marshal(payload)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(response))
}

// RespondJSONWithXTotalCount Make response with json format and add X-Total-Count header.
func RespondJSONWithXTotalCount(w http.ResponseWriter, code int, count int, payload interface{}) {
	response, error := json.Marshal(payload)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	SetNewHeader(w, "X-Total-Count", strconv.Itoa(count))
	w.WriteHeader(code)
	w.Write([]byte(response))
}

// RespondErrorMessage make error response with payload.
func RespondErrorMessage(w http.ResponseWriter, code int, msg string) {
	RespondJSON(w, code, map[string]string{"error": msg})
}

// RespondError check error type and Write to ResponseWriter.
func RespondError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *errors.ValidationError:
		RespondJSON(w, http.StatusBadRequest, err)
	case *errors.HTTPError:
		httpError := err.(*errors.HTTPError)
		RespondJSON(w, httpError.HTTPStatus, httpError.ErrorKey)
	default:
		RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
	}
}

// SetNewHeader will expose and set the given headerName and value
// 	SetNewHeader(w,"total","10") will set header "total" : "10"
func SetNewHeader(w http.ResponseWriter, headerName, value string) {
	w.Header().Add("Access-Control-Expose-Headers", headerName)
	w.Header().Set(headerName, value)
}
