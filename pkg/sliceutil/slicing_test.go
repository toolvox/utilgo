package sliceutil_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/pkg/sliceutil"
	"github.com/toolvox/utilgo/pkg/stringutil"
)

type String = stringutil.String

func TestSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		from  int
		to    int
		want  []int
	}{
		{"Positive indices", []int{1, 2, 3, 4, 5}, 1, 3, []int{2, 3}},
		{"Negative start index", []int{1, 2, 3, 4, 5}, -5, 3, []int{1, 2, 3}},
		{"Negative start index2", []int{1, 2, 3, 4, 5}, -10, 3, []int{1, 2, 3}},
		{"Out-of-range end index", []int{1, 2, 3, 4, 5}, 2, 10, []int{3, 4, 5}},
		{"Reversed indices", []int{1, 2, 3, 4, 5}, 3, -5, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceutil.Slice(tt.slice, tt.from, tt.to); !reflect.DeepEqual(got, tt.want) {
				require.Equal(t, tt.want, got, tt.name)
			}
		})
	}
}

func TestSplice(t *testing.T) {
	tests := []struct {
		name   string
		s      stringutil.String
		from   int
		to     int
		splice stringutil.String
		want   stringutil.String
	}{
		// {"Reversed indices", String("Hello, world!"), 7, -2, String("universe"), String("Hello, universe!")},
		{"Insert in the middle", String("Hello, world!"), 7, 7, String("wonderful "), String("Hello, wonderful world!")},
		// {"Replace in the middle", String("Hello, world!"), 7, 12, String("universe"), String("Hello, universe!")},
		// {"Negative start index", String("Hello, world!"), -100, 5, String("Hi"), String("Hi, world!")},
		// {"Out-of-range end index", String("Hello, world!"), 3, 100, String("p"), String("Help")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceutil.Splice(tt.s, tt.from, tt.to, tt.splice); !reflect.DeepEqual(got, tt.want) {
				require.Equal(t, tt.want.String(), got.String(), tt.name)

			}
		})
	}
}
