package errors

import (
	e "github.com/pkg/errors"
	"net/http"
)

const (
	BadRequest          = "BAD_REQUEST"
	NotFound            = "NOT_FOUND"
	InternalServerError = "INTERNAL_SERVER_ERROR"
)

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(message string) error {
	return e.New(message)
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   BadRequest,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   NotFound,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   InternalServerError,
	}
}
