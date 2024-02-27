package errs

import (
	"fmt"
	"strings"
)

// Validator is an interface implemented by types capable of self-validation.
// The Validate method is intended to check the implementing type for internal
// consistency or correctness, returning an error if validation fails.
//
// Implementors of [Validator] should return nil if the object is considered valid,
// or an instance of [ValidationErrors] if there are specific validation errors to report.
//
// You're free to return any error type, [ValidationErrors] is provided as a convenience.
type Validator interface {
	// Validate checks the object for validity, returning an error if it is not valid.
	// The error returned can be cast to a [ValidationErrors] type to inspect individual errors.
	Validate() error
}

// ValidationErrors is a slice of errors intended for aggregating multiple validation errors into a single error.
type ValidationErrors []error

// Error formats the [ValidationErrors] slice into a single string, showing all validation errors.
func (e ValidationErrors) Error() string {
	var errs []string
	for _, err := range e {
		errs = append(errs, err.Error())
	}
	return fmt.Sprintf("validation errors: [%s]", strings.Join(errs, ", "))
}

// Unwrap returns the slice of errors contained within [ValidationErrors], allowing individual errors to be examined.
func (e ValidationErrors) Unwrap() []error {
	return e
}

// OrNil checks if the [ValidationErrors] slice is empty and returns nil if true; otherwise, it returns the [ValidationErrors] slice itself as an error.
func (e ValidationErrors) OrNil() error {
	if len(e) == 0 {
		return nil
	}
	return e
}

// Errorf appends a new error, formatted according to a format specifier, to the [ValidationErrors] slice.
func (e *ValidationErrors) Errorf(format string, args ...interface{}) {
	*e = append(*e, fmt.Errorf(format, args...))
}
