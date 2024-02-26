package flags

import (
	"fmt"
	"strconv"
	"strings"
)

// CSVValue holds a slice of strings that are parsed from comma-separated values provided via flags.
// Provides [CSVValue.Ints]() and [CSVValue.Floats]() functions.
//
// See the [ParseCSV]() function if you want to parse to a custom type.
type CSVValue struct {
	Values []string
}

// String returns a string representation of the CSV values.
// This method implements the [flag.Value] interface.
func (csv CSVValue) String() string {
	return strings.Join(csv.Values, ",")
}

// Set parses the input string for comma-separated values and stores them.
// This method implements the [flag.Value] interface.
func (csv *CSVValue) Set(value string) error {
	csv.Values = strings.Split(value, ",")
	return nil
}

// Get returns the slice of values as an interface{}.
// This method implements the [flag.Getter] interface.
// The return type is always []string.
func (csv CSVValue) Get() any {
	return csv.Values
}

// Ints converts the stored CSV values to a slice of ints.
func (csv CSVValue) Ints() ([]int, error) {
	var ints []int
	for _, v := range csv.Values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("parsing '%s' to int: %w", v, err)
		}
		ints = append(ints, i)
	}
	return ints, nil
}

// Floats converts the stored CSV values to a slice of floats.
func (csv CSVValue) Floats() ([]float64, error) {
	var floats []float64
	for _, v := range csv.Values {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing '%s' to float: %w", v, err)
		}
		floats = append(floats, f)
	}
	return floats, nil
}

// ParseCSV converts the stored CSV value to a []T using a custom parser.
func ParseCSV[T any](csv CSVValue, parser func(string) (T, error)) ([]T, error) {
	var res []T
	for _, val := range csv.Values {
		parsedVal, err := parser(val)
		if err != nil {
			return nil, fmt.Errorf("parsing '%s': %w", val, err)
		}
		res = append(res, parsedVal)
	}
	return res, nil
}
