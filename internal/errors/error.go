package errors

import "fmt"

// Error is a base error that implement scw.SdkError
type Error struct {
	str string
	err error
}

// Error implement standard xerror.Wrapper interface
func (e *Error) Unwrap() error {
	return e.err
}

// Error implement standard error interface
func (e *Error) Error() string {
	str := "scaleway-sdk: " + e.str
	if e.err != nil {
		str += ": " + e.err.Error()
	}
	return str
}

// isScwSdkError implement SdkError interface
func (e *Error) isScwSdkError() {}

// New creates a new error with that same interface as fmt.Errorf
func New(format string, args ...interface{}) error {
	return &Error{
		str: fmt.Sprintf(format, args...),
	}
}

// Wrap an error with additional information
func Wrap(err error, format string, args ...interface{}) error {
	return &Error{
		err: err,
		str: fmt.Sprintf(format, args...),
	}
}
