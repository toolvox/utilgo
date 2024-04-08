package errs_test

import (
	"errors"
	"fmt"

	"github.com/toolvox/utilgo/pkg/errs"
)

func ExampleWrappedError_Error() {
	// Creating a base error
	baseErr := errors.New("this is a base error")

	// Wrapping the base error with additional context
	wrappedErr := errs.WrappedError{
		Message: "failed due to external system",
		Wrapped: baseErr,
	}

	// Printing the error
	fmt.Println(wrappedErr.Error())

	// Output: failed due to external system: this is a base error
}

func ExampleWrappedError_Unwrap() {
	// Creating a base error and wrapping it
	baseErr := errors.New("base error")
	wrappedErr := errs.Wrap("context added", baseErr)

	// Unwrapping the error
	unwrappedErr := errors.Unwrap(wrappedErr)
	fmt.Println(unwrappedErr)

	// Output: base error
}

func ExampleWrappedError_Is() {
	// Creating a target error to check against
	targetErr := errors.New("target error")

	// Creating a wrapped error containing the target error
	wrappedErr := errs.Wrap("additional context", targetErr)

	// Checking if the wrapped error contains the target error
	fmt.Println(errors.Is(wrappedErr, targetErr))

	// Output: true
}

func ExampleWrappedError_As() {
	// Creating a custom error type
	customErr := errs.Error("custom error")

	// Wrapping the custom error
	wrappedErr := errs.Wrap("context", customErr)

	// Attempting to type-assert the wrapped error back to the custom type
	var targetErr errs.Error
	if errors.As(wrappedErr, &targetErr) {
		fmt.Println(targetErr)
	} else {
		fmt.Println("error does not match")
	}

	// Output: custom error
}

func ExampleWrap() {
	// Creating two errors to wrap
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	// Wrapping the errors with additional context
	wrappedErr := errs.Wrap("multiple errors occurred", err1, err2)

	fmt.Println(wrappedErr.Error())

	// Output: multiple errors occurred: errors: [error 1, error 2]
}

func ExampleWrapf() {
	// Creating an error to wrap
	err := errors.New("specific error")

	// Wrapping the error with a formatted message
	wrappedErr := errs.Wrapf("failed at %d%%", 50, err)

	fmt.Println(wrappedErr.Error())

	// Output: failed at 50%: specific error
}
