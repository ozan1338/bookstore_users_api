package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type RestError interface {
	GetMessage() string
	GetStatus() int
	GetError() string
}

type restError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Err   string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func (e restError) GetError() string {
	return fmt.Sprintf(e.Err)
}


func (e restError) GetMessage() string {
	return e.Message
}


func (e restError) GetStatus() int {
	return e.Status
}

func NewRestError(message string, status int, err string) RestError {
	return restError{
		Message: message,
		Status:  status,
		Err: err,
	}
}

func NewBadRequestError(message string) RestError {
	return restError{
		Message: message,
		Status:  http.StatusBadRequest,
		Err: "bad_request",
	}
}

func NewNotFoundError(message string) RestError {
	return restError{
		Message: message,
		Status:  http.StatusNotFound,
		Err: "not_found",
	}
}

func NewInternalServerError(message string) RestError {
	return restError{
		Message: message,
		Status: http.StatusInternalServerError,
		Err: "internal_server_error",
	}
}

func NewUnauthorizedError(message string) RestError {
	return restError{
		Message: message,
		Status: http.StatusUnauthorized,
		Err: "unauthorized",
	}
}