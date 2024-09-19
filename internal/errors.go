package internal

// ValidationError represents an error that will be thrown when validating entity fields
type ValidationError struct {
	msg string
}

// Error returns the validation error message
func (err ValidationError) Error() string {
	return err.msg
}

// NewValidationError returns a new validation error
func NewValidationError(msg string) ValidationError {
	return ValidationError{
		msg: msg,
	}
}

func IsValidationError(err error) bool {
	switch err.(type) {
	case ValidationError:
		return true
	default:
		return false
	}
}
