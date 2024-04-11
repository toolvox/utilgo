package maputil_test

import (
	"cmp"
	"sort"
	"testing"

	"github.com/toolvox/utilgo/pkg/maputil"
)

// TestMerge ensures that merging maps works as expected for various scenarios.
func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		maps []map[string]int
		want map[string]int
	}{
		{
			name: "empty input",
			maps: []map[string]int{},
			want: map[string]int{},
		},
		{
			name: "single map",
			maps: []map[string]int{{"a": 1}},
			want: map[string]int{"a": 1},
		},
		{
			name: "multiple maps with unique keys",
			maps: []map[string]int{{"a": 1}, {"b": 2}, {"c": 3}},
			want: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name: "multiple maps with overlapping keys",
			maps: []map[string]int{{"a": 1, "b": 2}, {"b": 3, "c": 4}},
			want: map[string]int{"a": 1, "b": 3, "c": 4}, // Expect last value to overwrite previous
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maputil.Merge(tt.maps...)
			if len(got) != len(tt.want) {
				t.Errorf("Merge() got len = %v, want len %v", len(got), len(tt.want))
			}
			for k, v := range tt.want {
				if gv, ok := got[k]; !ok || gv != v {
					t.Errorf("Merge() got[%v] = %v, want %v", k, gv, v)
				}
			}
		})
	}
}

// TestKeys ensures that retrieving keys from a map works correctly.
func TestKeys(t *testing.T) {
	testMap := map[int]string{1: "a", 2: "b", 3: "c"}
	gotKeys := maputil.Keys(testMap)
	if len(gotKeys) != len(testMap) {
		t.Fatalf("Keys() returned %d keys, want %d", len(gotKeys), len(testMap))
	}
	wantKeysMap := make(map[int]bool)
	for k := range testMap {
		wantKeysMap[k] = true
	}
	for _, k := range gotKeys {
		if !wantKeysMap[k] {
			t.Errorf("Keys() returned unexpected key: %v", k)
		}
	}
}

// TestSortedKeys verifies the SortedKeys function sorts keys of various types.
func TestSortedKeys(t *testing.T) {
	tests := []struct {
		name string
		self map[int]string
		want []int
	}{
		{
			name: "int keys",
			self: map[int]string{3: "c", 1: "a", 2: "b"},
			want: []int{1, 2, 3},
		},
		{
			name: "empty map",
			self: map[int]string{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maputil.SortedKeys(tt.self)
			if !equalSlices(got, tt.want) {
				t.Errorf("SortedKeys() got %v, want %v", got, tt.want)
			}
		})
	}
}

// TestValues checks that values are correctly retrieved from a map.
func TestValues(t *testing.T) {
	testMap := map[string]int{"a": 1, "b": 2}
	gotValues := maputil.Values(testMap)
	if len(gotValues) != 2 {
		t.Fatalf("Values() returned %d values, want 2", len(gotValues))
	}
	wantValuesMap := map[int]bool{1: true, 2: true}
	for _, v := range gotValues {
		if !wantValuesMap[v] {
			t.Errorf("Values() returned unexpected value: %v", v)
		}
	}
}

// TestSortedValues verifies that values are sorted according to the keys.
func TestSortedValues(t *testing.T) {
	tests := []struct {
		name string
		self map[string]int
		want []int
	}{
		{
			name: "string keys",
			self: map[string]int{"b": 2, "a": 1, "c": 3},
			want: []int{1, 2, 3},
		},
		{
			name: "empty map",
			self: map[string]int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maputil.SortedValues(tt.self)
			if !equalSlices(got, tt.want) {
				t.Errorf("SortedValues() got %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to compare slices for equality, regardless of order.
func equalSlices[T cmp.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool { return less(a[i], a[j]) })
	sort.Slice(b, func(i, j int) bool { return less(b[i], b[j]) })
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Generic less function to compare two values.
func less[T cmp.Ordered](a, b T) bool {
	return a < b
}
