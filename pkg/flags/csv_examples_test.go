package flags_test

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"utilgo/pkg/flags"
)

// Typical use-cases for [utilgo/pkg/CSVValue]
func ExampleCSVValue() {
	// Setup: Simulate command line flag input
	os.Args = []string{"cmd", "-csv", "1,2,3,4,5"}

	var csv flags.CSVValue
	flag.Var(&csv, "csv", "Comma-separated values")
	flag.Parse()

	// Examples:
	fmt.Println("String values:", csv.Get())

	ints, _ := csv.Ints()
	fmt.Println("Int values:", ints)

	floats, _ := csv.Floats()
	fmt.Println("Float values:", floats)

	// Output:
	// String values: [1 2 3 4 5]
	// Int values: [1 2 3 4 5]
	// Float values: [1 2 3 4 5]
}

// Using the [utilgo/pkg/CSVValue.ParseCSV] function to convert CSV string values to a custom type using a custom parser.
func ExampleParseCSV() {
	// Setup:
	var csv flags.CSVValue
	csv.Set("true,false,true")

	// Example:
	bools, _ := flags.ParseCSV(csv, strconv.ParseBool)
	fmt.Println(bools)
	// Output: [true false true]
}
