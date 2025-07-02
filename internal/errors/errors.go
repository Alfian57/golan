package errs

import "errors"

type FieldError struct {
	Field string
	Error string
}

type ErrValidationErrors struct {
	Errors []FieldError
}

func (e ErrValidationErrors) Error() string {
	if len(e.Errors) == 0 {
		return "validation errors"
	}
	return e.Errors[0].Error
}

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrUsernameExist = errors.New("username already exists")
)
