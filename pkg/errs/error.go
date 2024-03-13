package errs

import "fmt"

// Error defines a simple string-based error type.
// It implements the error interface, allowing instances of Error to be used wherever Go errors are expected.
type Error string

// Error returns the string representation of the error.
// This method satisfies the error interface for the Error type, making it compatible with Go's built-in error handling.
func (e Error) Error() string { return string(e) }

// New creates a new error with the given text.
// It provides a simple way to create errors without needing to use [fmt.Errorf] or [errors.New] for basic cases.
func New(s string) error { return Error(s) }

// Newf creates a plain text error from a formatted string (using [Error] as the backing type).
// It's a convenient wrapper around [fmt.Sprintf], allowing for the creation of formatted error messages without needing [fmt.Errorf]'s '%u'.
func Newf(format string, args ...any) error {
	return Error(fmt.Sprintf(format, args...))
}
