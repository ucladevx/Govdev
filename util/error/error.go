// Package errors wraps Dave Cheney's errors package, by adding an error type
// and context to the error. Our custom errors implement the errors.Wrap
// interface, allowing a Cause function to work properly and allow the
// application to retrieve the cause of the error.
package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// ErrorType is the type of an error.
type ErrorType uint

// Each ErrorType can be mapped to the corresponding HTTP.Status error. For
// example, BadRequest ErrorType maps to http.StatusBadRequest.
const (
	NoType             ErrorType = iota // NoType error
	BadRequest                          // BadRequest error, data formatted badly
	Unauthorized                        // Unauthorized error, unauthenticated error
	Forbidden                           // Forbidden error, authentcated user without proper rights
	NotFound                            // NotFound error, resource is not found
	ServiceUnavailable                  // ServiceUnavailable error, Service is down
)

type customError struct {
	errorType     ErrorType
	originalError error
	context       errorContext
}

type errorContext struct {
	Field   string
	Message string
}

// New Creates a new custom error, with a defined error type. Function signature
// allows
func (errorType ErrorType) New(msg string, args ...interface{}) error {
	return customError{
		errorType:     errorType,
		originalError: fmt.Errorf(msg, args...),
	}
}

// NewWithContext creates a new custom error, but adds a context field to the
// custom error.
func (errorType ErrorType) NewWithContext(msg, ctxField, ctxMsg string, args ...interface{}) error {
	cError := errorType.New(msg, args).(customError)
	context := errorContext{
		Field:   ctxField,
		Message: ctxMsg,
	}
	cError.context = context
	return cError
}

// Wrap implements the errors.Wrap interface, allow us to wrap a custom error in
// another custom error. Calls Wrapf under the hood.
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrapf(err, msg)
}

// Wrapf calls is like Wrap but allows additional arguments.
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{
		errorType:     errorType,
		originalError: errors.Wrapf(err, msg, args...),
	}
}

// Cause returns the very original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Error implements the Error interface
func (ce customError) Error() string {
	return ce.originalError.Error()
}
