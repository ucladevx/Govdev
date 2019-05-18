// Package errors wraps Dave Cheney's errors package, by adding an error type
// and context to the error. Our custom errors implement the errors.Wrap
// interface, allowing a Cause function to work properly and allow the
// application to retrieve the cause of the error, and allows us to add context
// to errors by wrapping them.
//
//	err := fmt.Errorf("no such file")
//	err = Wrap(err, "open failed")
//	err = Wrap(err, "read config failed")
//	fmt.Println(err.Error()) // "read config failed: open failed: no such file"
//
// If we wanted to find the original error, then we can use the Cause function,
// which returns the first error that doesn't implement a Cause() function.
//
//	originalError := Cause(err)
//	fmt.Println(originalError.Error()) // "no such file"
//
// Additionally, we can add error types to the custom error, of a type
// ErrorType. By default, a NoType error is created, but we would also like to
// be able to create other types of errors, which can be done using
//
//	err = BadRequest.New("no such file")
// and
//	err = BadRequest.Wrap(err, "open failed")
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
}

// Error implements the Error interface
func (ce customError) Error() string {
	return ce.originalError.Error()
}

// Cause implements the causer interface in github.com/pkg/errors
func (ce customError) Cause() error {
	return ce.originalError
}

func New(msg string, args ...interface{}) error {
	return customError{
		errorType:     NoType,
		originalError: fmt.Errorf(msg, args...),
	}
}

func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

func Wrapf(err error, msg string, args ...interface{}) error {
	cerr, ok := err.(customError)
	if ok {
		return customError{
			errorType:     cerr.errorType,
			originalError: errors.Wrapf(err, msg, args...),
		}
	} else {
		return customError{
			errorType:     NoType,
			originalError: errors.Wrapf(err, msg, args...),
		}
	}

}

// Cause returns the very original error
func Cause(err error) error {
	return errors.Cause(err)
}

// New Creates a new custom error, with a defined error type. Function signature
// allows
func (errorType ErrorType) New(msg string, args ...interface{}) error {
	return customError{
		errorType:     errorType,
		originalError: fmt.Errorf(msg, args...),
	}
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

func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}
