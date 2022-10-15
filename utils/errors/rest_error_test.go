package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is internal server error")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is internal server error", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is not found server error")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "this is not found server error", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is bad request error")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is bad request error", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewError(t *testing.T) {
	err := NewError("this is new error")
	assert.NotNil(t, err)
}