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
