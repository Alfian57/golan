package errs

import (
	"fmt"
	"net/http"
)

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

type ValidationError struct {
	Errors []FieldError `json:"errors"`
}

func (e *ValidationError) Error() string {
	if len(e.Errors) == 0 {
		return "validation errors"
	}
	return e.Errors[0].Error
}

var (
	ErrUserNotFound   = &AppError{Code: http.StatusNotFound, Message: "user not found"}
	ErrUsernameExist  = &AppError{Code: http.StatusConflict, Message: "username already exists"}
	ErrInternalServer = &AppError{Code: http.StatusInternalServerError, Message: "internal server error"}
	ErrBadRequest     = &AppError{Code: http.StatusBadRequest, Message: "bad request"}
	ErrUnauthorized   = &AppError{Code: http.StatusUnauthorized, Message: "unauthorized"}
)

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewValidationError(fieldErrors []FieldError) *ValidationError {
	return &ValidationError{Errors: fieldErrors}
}

func NewFieldError(field, message string) FieldError {
	return FieldError{Field: field, Error: message}
}
