// Package errs provides an error and validation helpers.
// It defines a [Validator] interface for types that can validate themselves and various other error types for aggregating and detailing multiple validation errors.
//
// The [Validator] interface requires implementing types to provide a Validate method that checks for internal consistency or correctness, returning an error if the validation fails.
// This allows for self-validating models and other structures, making it easier to ensure data integrity throughout the application.
package errs

import (
	"errors"
	"io/fs"
	"syscall"
)

// Validator is an interface implemented by types capable of self-validation.
// The Validate method is intended to check the implementing type for internal consistency or correctness, returning an error if validation fails.
//
// Implementors of [Validator] should return nil if the object is considered valid, or any error if there are specific validation errors to report.
//
// You're free to return any error type, [Errors] is provided as a convenience.
type Validator interface {
	// Validate checks the object for validity, returning an error if it is not valid.
	// The error returned can be cast to a Errors type to inspect individual errors.
	Validate() error
}

// Must is a helper that takes a comma-error idiom and returns just the value, panicking if an error occurred.
func Must[T any](ret T, err error) T {
	if err != nil {
		panic(Wrap("must panic for", err))
	}
	return ret
}

// CheckPathError is used to deconstruct a [pkg/fs.PathError] to check its content.
func CheckPathError(err error, op, path string, innerErr syscall.Errno) bool {
	var errAsTarget *fs.PathError
	if !errors.As(err, &errAsTarget) {
		return false
	}

	if errAsTarget.Op != op {
		return false
	}
	if errAsTarget.Path != path {
		return false
	}
	return errors.Is(errAsTarget.Err, innerErr)
}
