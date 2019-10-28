package model

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ErrorServer defines default message for system error
const ErrorServer = "Internal server error"

// Error defines error details
type Error struct {
	Message    string `json:"message"`
	Status     uint32 `json:"-"`
	messageLog string
}

// NewErrorNil creates empty Error
func NewErrorNil() Error {
	return Error{}
}

// NewErrorBadRequest creates error with HTTP code 400
func NewErrorBadRequest(m string) Error {
	return newError(m, http.StatusBadRequest)
}

// NewErrorServer creates error with HTTP code 500
func NewErrorServer(m string) Error {
	return newError(m, http.StatusInternalServerError)
}

// String returns Error converted to JSON
func (e Error) String() (string, error) {
	bytes, err := json.Marshal(e)
	return string(bytes[:]), err
}

// WithError adds Golang error message to Error log message
func (e Error) WithError(err error) Error {
	e.messageLog = e.Message + ": " + err.Error()
	return e
}

// WithErrorMessage adds message to Error log message
func (e Error) WithErrorMessage(msg string) Error {
	e.messageLog = fmt.Sprintf("%s with message: %s", e.Message, msg)
	return e
}

// WithErrorObject copy log message from past Error
func (e Error) WithErrorObject(err Error) Error {
	e.messageLog = err.messageLog
	return e
}

// newError creates Error struct
func newError(m string, s uint32) Error {
	return Error{
		Message: m,
		Status:  s,
	}
}
