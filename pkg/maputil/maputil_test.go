package maputil_test

import (
	"reflect"
	"testing"

	"utilgo/pkg/maputil"
)

func TestSortedKeys(t *testing.T) {
	tests := []struct {
		name string
		self map[int]string
		want []int
	}{
		{
			name: "empty map",
			self: map[int]string{},
			want: []int{},
		},
		{
			name: "single element",
			self: map[int]string{1: "a"},
			want: []int{1},
		},
		{
			name: "multiple elements",
			self: map[int]string{2: "b", 1: "a", 3: "c"},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maputil.SortedKeys(tt.self); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
