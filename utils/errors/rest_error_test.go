package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is internal server error")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.GetStatus())
	assert.EqualValues(t, "this is internal server error", err.GetMessage())
	assert.EqualValues(t, "internal_server_error", err.GetError())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is not found server error")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.GetStatus())
	assert.EqualValues(t, "this is not found server error", err.GetMessage())
	assert.EqualValues(t, "not_found", err.GetError())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is bad request error")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is bad request error", err.GetMessage())
	assert.EqualValues(t, "bad_request", err.GetError())
}

func TestNewError(t *testing.T) {
	err := NewError("this is new error")
	assert.NotNil(t, err)
}