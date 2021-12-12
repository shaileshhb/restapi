package errors

// HTTPError Represents HTTP type error
type HTTPError struct {
	HTTPStatus int    `example:"400" json:"httpStatus"`
	ErrorKey   string `example:"unable to fetch data" json:"errorKey"`
}

// Implementing Error interface
func (httpError HTTPError) Error() string {
	return httpError.ErrorKey
}

// NewHTTPError returns new instance of HTTPError
func NewHTTPError(key string, statuscode int) *HTTPError {
	return &HTTPError{
		HTTPStatus: statuscode,
		ErrorKey:   key,
	}
}
