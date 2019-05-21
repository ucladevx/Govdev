package errors

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestIs(t *testing.T) {
	assert := assert.New(t)

	err := io.EOF
	err1 := &govError{
		err:       err,
		errorType: BadRequest,
	}
	err2 := xerrors.Errorf("error occured while reading file: %w", err1)
	err3 := xerrors.Errorf("read config failed: %w", err2)
	err4 := xerrors.Errorf("error occured while reading file: %w", err1)

	var gerr *govError
	assert.Equal(xerrors.As(err3, &gerr), true)
	assert.Equal(gerr.errorType, BadRequest)
	assert.Equal(gerr.err, io.EOF)

	assert.Equal(xerrors.Is(err3, err4), false)
	assert.Equal(xerrors.Is(err3, err3), true)
	assert.Equal(xerrors.Is(err3, err2), true)
	assert.Equal(xerrors.Is(err3, err1), true)
	assert.Equal(xerrors.Is(err3, io.EOF), true)
	assert.Equal(xerrors.Is(err3, &govError{}), false)
	assert.Equal(xerrors.Is(err3, err), true)
}

func TestXWarpf(t *testing.T) {
	assert := assert.New(t)

	err1 := xerrors.New("no such file")
	err2 := XWrap(err1, "open failed")
	err3 := XWrap(err2, "read config failed")

	assert.Equal(err3.Error(), "read config failed: open failed: no such file")
	assert.Equal(xerrors.Unwrap(err3), err2)
	assert.Equal(xerrors.Is(err3, err2), true)
	assert.Equal(xerrors.Is(err3, err1), true)
	assert.Equal(xerrors.Is(err3, xerrors.New("no such file")), false)
}

func TestXWrapfgovError(t *testing.T) {
	assert := assert.New(t)

	err := xerrors.New("no such file")
	g := &govError{
		err:       err,
		errorType: BadRequest,
	}
	err1 := XWrap(g, "open failed")
	err2 := XWrap(err1, "read config failed")

	assert.Equal(err2.Error(), "read config failed: open failed: no such file")
	var gov *govError
	assert.Equal(xerrors.As(err2, &gov), true)
	assert.Equal(gov.err.Error(), "no such file")
	assert.Equal(gov.errorType, BadRequest)
}
