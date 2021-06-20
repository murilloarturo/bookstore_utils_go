package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is a message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is a bad request")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "this is a bad request", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)

	assert.Nil(t, err.Causes)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is a not found error")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "this is a not found error", err.Message)
	assert.EqualValues(t, "not_found", err.Error)

	assert.Nil(t, err.Causes)
}

func TestNewError(t *testing.T) {
	err := NewError("error message")

	assert.NotNil(t, err)
	assert.EqualValues(t, "error message", err.Error())
}
