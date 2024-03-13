package errs

import (
	"fmt"
	"slices"

	"utilgo/pkg/reflectutil"
)

// WrappedError represents an error that wraps another error with an additional message.
// It allows for adding context to the original error, making it easier to understand the error's origin and nature.
type WrappedError struct {
	Message string
	Wrapped error
}

// Error returns a formatted string combining the message and the wrapped error.
// The format is "<Message>: <Wrapped Error>", providing a clear indication of the context followed by the original error.
func (e WrappedError) Error() string { return fmt.Sprintf("%s: %s", e.Message, e.Wrapped) }

// Unwrap returns the original error that was wrapped. This allows [errors.Is] and [errors.As] to work with [WrappedError],
// facilitating error inspection and handling.
func (e WrappedError) Unwrap() error { return e.Wrapped }

func (e WrappedError) Is(target error) bool {
	return e.Error() == target.Error()
}

// Wraps takes a string wrapper and a variadic slice of errors to create a new [WrappedError].
// Any passed errors that are nil are discarded.
// Returns a [WrappedError] containing filtered error(s) or nil or no valid errors were passed.
func Wrap(wrapper string, errs ...error) error {
	errs = slices.DeleteFunc(errs, reflectutil.IsZero)
	switch len(errs) {
	case 0:
		return nil

	case 1:
		return WrappedError{
			Message: wrapper,
			Wrapped: errs[0],
		}

	default:
		return WrappedError{
			Message: wrapper,
			Wrapped: Errors(errs),
		}
	}
}

func Wrapf(wrapperFormat string, argsAndErrs ...any) error {
	split := len(argsAndErrs)
	var errs []error
	for ; split > 0; split-- {
		if err, ok := argsAndErrs[split-1].(error); ok {
			errs = append(errs, err)
		} else {
			break
		}
	}
	args := argsAndErrs[:max(0, split)]
	return Wrap(fmt.Sprintf(wrapperFormat, args...), errs...)
}
