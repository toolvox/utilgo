package errs_test

import (
	"fmt"
	"strings"

	"utilgo/api/errs"
)

// ExampleType demonstrates a type that implements the Validator interface.
type ExampleType struct {
	Field1 string
	Field2 int
}

// Validate implements the Validator interface for ExampleType.
// It checks multiple conditions and aggregates any errors found using Errorf.
func (et ExampleType) Validate() error {
	var errs errs.ValidationErrors

	if strings.TrimSpace(et.Field1) == "" {
		errs.Errorf("Field1 must not be empty")
	}

	if et.Field2 <= 0 {
		errs.Errorf("Field2 must be positive")
	}

	return errs.OrNil()
}

func ExampleValidator() {
	exampleInvalid := ExampleType{Field1: "", Field2: -1}
	fmt.Printf("Validating: %#v\n", exampleInvalid)
	if err := exampleInvalid.Validate(); err != nil {
		fmt.Println("Validation failed:", err)
		// To inspect individual errors:
		if valErrs, ok := err.(errs.ValidationErrors); ok {
			for _, e := range valErrs.Unwrap() {
				fmt.Println("Detail:", e)
			}
		}
	} else {
		fmt.Println("Validation passed")
	}

	exampleValid := ExampleType{Field1: "Bob", Field2: 13}
	fmt.Printf("Validating: %+v\n", exampleValid)
	if err := exampleValid.Validate(); err != nil {
		fmt.Println("Validation failed:", err)
		// To inspect individual errors:
		if valErrs, ok := err.(errs.ValidationErrors); ok {
			for _, e := range valErrs.Unwrap() {
				fmt.Println("Detail:", e)
			}
		}
	} else {
		fmt.Println("Validation passed")
	}

	// Output:
	// Validating: errs_test.ExampleType{Field1:"", Field2:-1}
	// Validation failed: validation errors: [Field1 must not be empty, Field2 must be positive]
	// Detail: Field1 must not be empty
	// Detail: Field2 must be positive
	// Validating: {Field1:Bob Field2:13}
	// Validation passed
}
