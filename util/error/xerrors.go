package errors

import "golang.org/x/xerrors"

type govError struct {
	errorType ErrorType
	err       error
}

func (e *govError) Error() string {
	return e.err.Error()
}

func (e *govError) Unwrap() error {
	return e.err
}

var _ xerrors.Wrapper = &govError{}
var _ error = &govError{}

// XWrap is a convenience function around XWrapf
func XWrap(err error, msg string) error {
	return XWrapf(err, msg)
}

// XWrapf takes an error, and wraps it using Errorf by appending a ": %w" string
// to the end of the message, and appending the errors to the list of args. The
// resulting error implements the xerrors.Wrapper interface
func XWrapf(err error, msg string, args ...interface{}) error {
	errMsg := msg + ": %w"
	args = append(args, err)
	return xerrors.Errorf(errMsg, args...)
}
