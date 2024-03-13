package errs

import (
	"fmt"
	"slices"
	"strings"

	"utilgo/pkg/reflectutil"
)

// [Errors] is a slice of errors intended for easily aggregating errors into a single error.
//
// The default value of [Errors] or an empty [Errors] is treated as a nil error.
//
//	> nil, default, or empty error elements will be discarded
type Errors []error

// Clean removes all the zero values (read: nil error or empty error structs).
func (e Errors) Clean() []error { return slices.DeleteFunc(e, reflectutil.IsZero) }

// Error formats the [Errors] slice into a single string, showing all aggregated errors.
func (e Errors) Error() string {
	es := e.Clean()
	if len(es) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("errors: [")
	sb.WriteString(es[0].Error())
	for _, err := range es[1:] {
		sb.WriteString(", ")
		sb.WriteString(err.Error())
	}
	sb.WriteString("]")
	return sb.String()
}

// Unwrap returns the slice of errors contained within [Errors], allowing individual errors to be examined.
func (e Errors) Unwrap() []error {
	return e.Clean()
}

// OrNil checks if the [Errors] slice is empty and returns nil if true; otherwise, it returns the [Errors] slice itself as an error.
func (e Errors) OrNil() error {
	if len(e.Clean()) == 0 {
		return nil
	}
	return Errors(e.Clean())
}

// WithError appends a new error to the [Errors] slice.
//
// If the error is nil, nothing is appended.
func (e *Errors) WithError(err error) {
	if err == nil {
		return
	}
	*e = append(*e, err)
}

// WithErrorf appends a new error, formatted according to a format specifier, to the [Errors] slice.
//
// "%u" is supported as this is a wrapper for [fmt.Errorf].
func (e *Errors) WithErrorf(format string, args ...interface{}) {
	if len(format) == 0 {
		return
	}
	e.WithError(fmt.Errorf(format, args...))
}
