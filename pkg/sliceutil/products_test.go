package sliceutil

import (
	"reflect"
	"testing"
)

func TestPowerSet(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want [][]int
	}{
		{
			name: "Empty set",
			src:  []int{},
			want: [][]int{{}},
		},
		{
			name: "Single element",
			src:  []int{1},
			want: [][]int{{}, {1}},
		},
		{
			name: "Two elements",
			src:  []int{1, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "Three elements",
			src:  []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowerSet(tt.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PowerSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct(t *testing.T) {
	tests := []struct {
		name   string
		slices [][]int
		want   [][]int
	}{
		{
			name:   "Empty input",
			slices: [][]int{},
			want:   [][]int{},
		},
		{
			name:   "Single slice",
			slices: [][]int{{1, 2}},
			want:   [][]int{{1}, {2}},
		},
		{
			name:   "Two slices",
			slices: [][]int{{1, 2}, {3, 4}},
			want:   [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}},
		},
		{
			name:   "Three slices",
			slices: [][]int{{1}, {2, 3}, {4, 5}},
			want:   [][]int{{1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}},
		},
		{
			name:   "Three slices 2",
			slices: [][]int{{1, 6}, {2, 3}, {4, 5}},
			want:   [][]int{{1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {6, 2, 4}, {6, 2, 5}, {6, 3, 4}, {6, 3, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slices := make([][]int, len(tt.slices))
			copy(slices, tt.slices)
			if got := Product(slices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Product() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrefixes(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  [][]int
	}{
		{"empty slice", []int{}, [][]int{{}}},
		{"single element", []int{1}, [][]int{{}, {1}}},
		{"multiple elements", []int{1, 2, 3}, [][]int{{}, {1}, {1, 2}, {1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prefixes(tt.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuffixes(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  [][]int
	}{
		{"empty slice", []int{}, [][]int{{}}},
		{"single element", []int{1}, [][]int{{}, {1}}},
		{"multiple elements", []int{1, 2, 3}, [][]int{{}, {3}, {2, 3}, {1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Suffixes(tt.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Suffixes() = %v, want %v", got, tt.want)
			}
		})
	}
}
