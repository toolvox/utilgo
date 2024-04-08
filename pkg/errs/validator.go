package errs

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
