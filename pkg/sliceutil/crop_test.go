package sliceutil_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/pkg/sliceutil"
)

func TestCropIndex(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		index    int
		expected []int
	}{
		{"Remove middle element", []int{1, 2, 3}, 1, []int{1, 3}},
		{"Remove first element", []int{1, 2, 3}, 0, []int{2, 3}},
		{"Remove last element", []int{1, 2, 3}, 2, []int{1, 2}},
		{"Index out of range", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"Negative index", []int{1, 2, 3}, -1, []int{1, 2, 3}},
		{"Single element 1", []int{1}, 0, []int{}},
		{"Single element 2", []int{2}, 1, []int{2}},
		{"Empty slice", []int{}, 0, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := sliceutil.CropIndex(c.input, c.index)
			require.ElementsMatch(t, result, c.expected)
		})
	}
}

func TestCropIndices(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		indices  []int
		expected []int
	}{
		{"Remove multiple indices", []int{1, 2, 3, 4, 5}, []int{1, 3}, []int{1, 3, 5}},
		{"Indices out of range", []int{1, 2, 3}, []int{3, 4}, []int{1, 2, 3}},
		{"Negative index", []int{1, 2, 3}, []int{-1, 0}, []int{2, 3}},
		{"Empty slice", []int{}, []int{0, 1}, []int{}},
		{"Remove all elements", []int{1, 2, 3}, []int{0, 1, 2}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := sliceutil.CropIndices(c.input, c.indices...)
			require.ElementsMatch(t, result, c.expected)
		})
	}
}

func TestCropElement(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		element  string
		expected []string
	}{
		{"Remove existing element", []string{"a", "b", "c", "d"}, "b", []string{"a", "c", "d"}},
		{"Element not found", []string{"a", "b", "c"}, "d", []string{"a", "b", "c"}},
		{"Empty slice", []string{}, "a", []string{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := sliceutil.CropElement(c.input, c.element)
			require.ElementsMatch(t, result, c.expected)
		})
	}
}

func TestCropElements(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		elements []int
		expected []int
	}{
		{"Remove multiple elements", []int{1, 2, 3, 4, 5}, []int{2, 4}, []int{1, 3, 5}},
		{"Elements not found", []int{1, 2, 3}, []int{4, 5}, []int{1, 2, 3}},
		{"Empty slice", []int{}, []int{1, 2}, []int{}},
		{"Remove all elements", []int{1, 1, 1}, []int{1}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := sliceutil.CropElements(c.input, c.elements...)
			require.ElementsMatch(t, result, c.expected)
		})
	}
}
