package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	err := BadRequest.New("New error")
	switch v := err.(type) {
	case error:
		// err is an error
	default:
		t.Fatalf("New should return an error type, type is %s", v)
	}
	cerr, ok := err.(customError)
	if !ok {
		t.Fatal("Error should be of type customError")
	}
	assert.Equal(cerr.errorType, BadRequest, "ErrorType should be BadRequest")

	wrappedError := BadRequest.Wrapf(err, "Error Wrap %d", 1)

	assert.Equal(GetType(wrappedError), BadRequest, "Error type of wrapped request should be BadRequest, was %d", GetType(wrappedError))
	assert.EqualError(wrappedError, "Error Wrap 1: New error")

}

func TestWrap(t *testing.T) {
	assert := assert.New(t)
	err := fmt.Errorf("no such file")
	err = Wrap(err, "open failed")
	err = Wrap(err, "read config failed")

	assert.Equal(GetType(err), NoType, "Error type of wrapped request should be NoType, was %d", GetType(err))
	assert.EqualError(err, "read config failed: open failed: no such file")

	assert.EqualError(Cause(err), "no such file")
}
