package errors

// ValidationError Represents Validation type error
type ValidationError struct {
	Err string `json:"error" example:"Name must be specified"`
}

// Error Implements error interface
func (vError ValidationError) Error() string {
	return vError.Err
}

// NewValidationError returns new instance of Validation error.
func NewValidationError(error string) *ValidationError {
	return &ValidationError{
		Err: error,
	}
}
