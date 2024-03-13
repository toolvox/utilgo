package flags_test

import (
	"fmt"
	"reflect"
	"testing"

	"utilgo/pkg/flags"
)

func TestCSVValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{} // Use interface{} to handle different types of expected results
		testFunc func(*flags.CSVValue) (interface{}, error)
		wantErr  bool
	}{
		{
			name:     "SetAndGet valid input",
			input:    "1,2,3,4,5",
			expected: []string{"1", "2", "3", "4", "5"},
			testFunc: func(csv *flags.CSVValue) (interface{}, error) {
				return csv.Get().([]string), nil
			},
			wantErr: false,
		},
		{
			name:     "Ints valid input",
			input:    "1,2,3,-4,5",
			expected: []int{1, 2, 3, -4, 5},
			testFunc: func(csv *flags.CSVValue) (interface{}, error) {
				return csv.Ints()
			},
			wantErr: false,
		},
		{
			name:     "Floats valid input",
			input:    "1.1,2.2,3.3,-4.4,5.5",
			expected: []float64{1.1, 2.2, 3.3, -4.4, 5.5},
			testFunc: func(csv *flags.CSVValue) (interface{}, error) {
				return csv.Floats()
			},
			wantErr: false,
		},
		{
			name:     "Ints invalid input",
			input:    "not,a,number",
			expected: nil,
			testFunc: func(csv *flags.CSVValue) (interface{}, error) {
				return csv.Ints()
			},
			wantErr: true,
		},
		{
			name:     "Floats invalid input",
			input:    "not,a,float",
			expected: nil,
			testFunc: func(csv *flags.CSVValue) (interface{}, error) {
				return csv.Floats()
			},
			wantErr: true,
		},
		{
			name:     "ParseCSV with custom parser and invalid input",
			input:    "1,true,not_boolean",
			expected: nil, // Expected is nil since we anticipate an error
			testFunc: func(csv *flags.CSVValue) (interface{}, error) {
				// Custom parser function that converts string to bool but fails on invalid boolean values
				return flags.ParseCSV(*csv, func(val string) (bool, error) {
					if val == "true" {
						return true, nil
					} else if val == "false" {
						return false, nil
					}
					// Return an error for any input that is not a valid boolean string
					return false, fmt.Errorf("invalid boolean value")
				})
			},
			wantErr: true, // Indicating that we expect this test case to trigger an error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var csv flags.CSVValue
			err := csv.Set(tt.input)
			if err != nil {
				t.Fatalf("CSVValue.Set() failed: %v", err)
			}

			got, err := tt.testFunc(&csv)
			if (err != nil) != tt.wantErr {
				t.Errorf("testFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("testFunc() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
