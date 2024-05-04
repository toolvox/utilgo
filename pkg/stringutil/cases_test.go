package stringutil_test

import (
	"testing"

	"github.com/toolvox/utilgo/pkg/stringutil"
)

const (
	FIRST_CASE = stringutil.CamelCase
	LAST_CASE  = stringutil.TitleCase
)

func TestStringCaseConversions(t *testing.T) {
	testCases := [][4]string{
		{"helloWorld", "HelloWorld", "hello_world", "Hello World"},
		{"myTestString", "MyTestString", "my_test_string", "My Test String"},
		{"sampleInput", "SampleInput", "sample_input", "Sample Input"},
		{"goIsGreat", "GoIsGreat", "go_is_great", "Go Is Great"},
		{"vesselShotGlass666", "VesselShotGlass666", "vessel_shot_glass_666", "Vessel Shot Glass 666"},
		{"420Fast69FuriousAgain", "420Fast69FuriousAgain", "420_fast_69_furious_again", "420 Fast 69 Furious Again"},
	}

	// Generate all possible conversions using a nested loop over integer values
	for from := FIRST_CASE; from <= LAST_CASE; from++ {
		for to := FIRST_CASE; to <= LAST_CASE; to++ {
			for _, tc := range testCases {
				input, expected := tc[from], tc[to]
				t.Run(input+"|"+expected, func(t *testing.T) {
					result := stringutil.TransCase(input, from, to)
					if result != expected {
						t.Errorf("expected %s but got %s", expected, result)
					}
				})
			}
		}
	}
}
