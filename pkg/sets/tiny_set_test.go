package sets_test

import (
	"reflect"
	"testing"

	"utilgo/pkg/sets"
)

func TestTinySet(t *testing.T) {
	tests := []struct {
		name          string
		operation     string
		set1          sets.TinySet[int]
		set2          sets.TinySet[int]
		addElements   []int
		expected      interface{}
		expectedThree [3]sets.TinySet[int]
	}{
		{
			name:        "Union with Set",
			operation:   "SetUnion",
			set1:        sets.NewTinySet(1, 2),
			set2:        sets.NewTinySet(2, 3),
			expected:    sets.NewTinySet(1, 2, 3),
		},
		{
			name:        "Union with Elements",
			operation:   "Union",
			set1:        sets.NewTinySet(1, 2),
			addElements: []int{3, 2},
			expected:    sets.NewTinySet(1, 2, 3),
		},
		{
			name:        "Intersection with Set",
			operation:   "SetIntersection",
			set1:        sets.NewTinySet(1, 2),
			set2:        sets.NewTinySet(2, 3),
			expected:    sets.NewTinySet(2),
		},
		{
			name:        "Intersection with Elements",
			operation:   "Intersection",
			set1:        sets.NewTinySet(1, 2),
			addElements: []int{2, 3},
			expected:    sets.NewTinySet(2),
		},
		{
			name:        "Difference with Set",
			operation:   "SetDifference",
			set1:        sets.NewTinySet(1, 2),
			set2:        sets.NewTinySet(2, 3),
			expected:    sets.NewTinySet(1),
		},
		{
			name:        "Difference with Elements",
			operation:   "Difference",
			set1:        sets.NewTinySet(1, 2),
			addElements: []int{2, 3},
			expected:    sets.NewTinySet(1),
		},
		{
			name:      "ThreeWay comparison",
			operation: "ThreeWay",
			set1:      sets.NewTinySet(1, 2),
			set2:      sets.NewTinySet(2, 3),
			expectedThree: [3]sets.TinySet[int]{
				sets.NewTinySet(2), // Common
				sets.NewTinySet(1), // Only in set1
				sets.NewTinySet(3), // Only in set2
			},
		},
		{
			name:        "Contains",
			operation:   "Contains",
			set1:        sets.NewTinySet(1, 2, 3),
			addElements: []int{2},
			expected:    true,
		},
		{
			name:        "String representation",
			operation:   "String",
			set1:        sets.NewTinySet(1, 2, 3),
			expected:    "{1, 2, 3}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var result interface{}

			switch tc.operation {
			case "SetUnion":
				result = tc.set1.SetUnion(tc.set2)
			case "Union":
				result = tc.set1.Union(tc.addElements...)
			case "SetIntersection":
				result = tc.set1.SetIntersection(tc.set2)
			case "Intersection":
				result = tc.set1.Intersection(tc.addElements...)
			case "SetDifference":
				result = tc.set1.SetDifference(tc.set2)
			case "Difference":
				result = tc.set1.Difference(tc.addElements...)
			case "ThreeWay":
				result = tc.set1.ThreeWay(tc.set2)
				if !reflect.DeepEqual(result, tc.expectedThree) {
					t.Errorf("Expected %v, got %v", tc.expectedThree, result)
				}
				return
			case "Contains":
				result = tc.set1.Contains(tc.addElements[0])
			case "String":
				result = tc.set1.String()
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
