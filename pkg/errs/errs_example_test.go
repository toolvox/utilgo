package errs_test

import (
	"fmt"
	"strings"

	"github.com/toolvox/utilgo/pkg/errs"
)

// ExampleType shows how a type can implement the Validator interface.
type ExampleType struct {
	Field1 string
	Field2 int
}

// Validate implements the Validator interface for ExampleType.
// It checks multiple conditions and aggregates any errors found using Errorf.
func (et ExampleType) Validate() error {
	var errors errs.Errors

	if strings.TrimSpace(et.Field1) == "" {
		errors.WithErrorf("Field1 must not be empty")
	}

	if et.Field2 <= 0 {
		errors.WithErrorf("Field2 must be positive")
	}

	// an accidental nil
	errors = append(errors, nil)

	if len(errors.Unwrap()) > 0 {
		errors.WithErrorf("%s error: %w", "another", errs.New("inner error"))
	}

	return errs.Wrap("validation", errors.OrNil())
}

func ExampleValidator() {
	exampleInvalid := ExampleType{Field1: "", Field2: -1}
	fmt.Printf("Validating: %#v\n", exampleInvalid)

	if err := exampleInvalid.Validate(); err != nil {
		fmt.Println("Validation failed:\n", err)
	} else {
		fmt.Println("Validation passed")
	}

	exampleValid := ExampleType{Field1: "Bob", Field2: 13}
	fmt.Printf("Validating: %+v\n", exampleValid)

	if err := exampleValid.Validate(); err != nil {
		fmt.Println("Validation failed:\n", err)
	} else {
		fmt.Println("Validation passed")
	}

	// Output:
	// Validating: errs_test.ExampleType{Field1:"", Field2:-1}
	// Validation failed:
	//  validation: errors: [Field1 must not be empty, Field2 must be positive, another error: inner error]
	// Validating: {Field1:Bob Field2:13}
	// Validation passed
}
