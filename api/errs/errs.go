// Package errs provides an error handling and validation utility.
// It defines a [Validator] interface for types that can validate themselves and a [ValidationErrors] type
// for aggregating and detailing multiple validation errors within a single error context.
//
// The [Validator] interface requires implementing types to provide a Validate method that checks
// for internal consistency or correctness, returning an error if the validation fails. This allows
// for self-validating models and other structures, making it easier to ensure data integrity throughout
// the application.
//
// [ValidationErrors] is designed to collect multiple errors encountered during the validation process,
// enabling detailed and user-friendly error reporting. It provides methods for formatting error messages,
// accessing individual errors, and conditionally returning nil if no errors are present.
package errs
