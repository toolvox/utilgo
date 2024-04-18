package sets_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/toolvox/utilgo/pkg/sets"
	"github.com/toolvox/utilgo/test"
	test_sets "github.com/toolvox/utilgo/test/sets_test"
)

func Test_TinySet(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		t.Run("new", func(t *testing.T) {
			testSet := sets.NewTinySet(0, 1, 2, 3, 1, 3, 5)
			require.True(t, testSet.Contains(1, 2, 3, 5))
		})
		test_sets.Run_Test_Set(t,
			test.TestConstructorFor[*sets.TinySet[int]]{
				NewFunc: func() *sets.TinySet[int] {
					return sets.NewTinySet[int]()
				},
			},
			test.TestDataFor[int]{
				ElementFunc: func(i int) int { return i },
			},
			100,
		)
	})
	t.Run("string", func(t *testing.T) {
		t.Run("new", func(t *testing.T) {
			testSet := sets.NewTinySet("0", "1", "2", "3", "1", "3", "5")
			require.True(t, testSet.Contains("1", "2", "3", "5"))
		})
		test_sets.Run_Test_Set(t,
			test.TestConstructorFor[*sets.TinySet[string]]{
				NewFunc: func() *sets.TinySet[string] {
					return sets.NewTinySet[string]()
				},
			},
			test.TestDataFor[string]{
				ElementFunc: func(i int) string { return fmt.Sprint(i) },
			},
			100,
		)
	})
}

func TestTinySet(t *testing.T) {
	tests := []struct {
		name          string
		operation     string
		set1          *sets.TinySet[int]
		set2          *sets.TinySet[int]
		addElements   []int
		expected      *sets.TinySet[int]
		expectedThree [3]sets.TinySet[int]
	}{
		{
			name:      "UnionWith with Set",
			operation: "Union",
			set1:      sets.NewTinySet(1, 2),
			set2:      sets.NewTinySet(2, 3),
			expected:  sets.NewTinySet(1, 2, 3),
		},
		{
			name:        "UnionWith with Elements",
			operation:   "UnionWith",
			set1:        sets.NewTinySet(1, 2),
			addElements: []int{3, 2},
			expected:    sets.NewTinySet(1, 2, 3),
		},
		{
			name:      "IntersectionWith with Set",
			operation: "Intersection",
			set1:      sets.NewTinySet(1, 2),
			set2:      sets.NewTinySet(2, 3),
			expected:  sets.NewTinySet(2),
		},
		{
			name:        "IntersectionWith with Elements",
			operation:   "IntersectionWith",
			set1:        sets.NewTinySet(1, 2),
			addElements: []int{2, 3},
			expected:    sets.NewTinySet(2),
		},
		{
			name:      "DifferenceWith with Set",
			operation: "Difference",
			set1:      sets.NewTinySet(1, 2),
			set2:      sets.NewTinySet(2, 3),
			expected:  sets.NewTinySet(1),
		},
		{
			name:        "DifferenceWith with Elements",
			operation:   "DifferenceWith",
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
				*sets.NewTinySet(2), // Common
				*sets.NewTinySet(1), // Only in set1
				*sets.NewTinySet(3), // Only in set2
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var result interface{}

			switch tc.operation {
			case "Union":
				result = tc.set1.Union(*tc.set2)
			case "UnionWith":
				result = tc.set1.UnionWith(tc.addElements...)
			case "Intersection":
				result = tc.set1.Intersection(*tc.set2)
			case "IntersectionWith":
				result = tc.set1.IntersectionWith(tc.addElements...)
			case "Difference":
				result = tc.set1.Difference(*tc.set2)
			case "DifferenceWith":
				result = tc.set1.DifferenceWith(tc.addElements...)
			case "ThreeWay":
				result = tc.set1.ThreeWay(*tc.set2)
				if !reflect.DeepEqual(result, tc.expectedThree) {
					t.Errorf("Expected %v, got %v", tc.expectedThree, result)
				}
				return
			case "Contains":
				result = tc.set1.Contains(tc.addElements[0])
			case "String":
				result = tc.set1.String()
			}

			if !reflect.DeepEqual(result, *tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func Test_TinySet_Encoding(t *testing.T) {
	testSet := sets.NewTinySet("hello", "goodbye", "salute")
	t.Run("YAML", func(t *testing.T) {
		var bs []byte
		var err error
		t.Run("Marshal", func(t *testing.T) {
			bs, err = yaml.Marshal(testSet)
			require.NoError(t, err)
			for _, v := range *testSet {
				require.Contains(t, string(bs), v)
			}
		})
		t.Run("Unmarshal", func(t *testing.T) {
			var cycleSet *sets.TinySet[string]
			err = yaml.Unmarshal(bs, &cycleSet)
			require.NoError(t, err)
			require.Equal(t, testSet, cycleSet)
		})
	})

	t.Run("JSON", func(t *testing.T) {
		var bs []byte
		var err error
		t.Run("Marshal", func(t *testing.T) {
			bs, err = json.Marshal(testSet)
			require.NoError(t, err)
			for _, v := range *testSet {
				require.Contains(t, string(bs), v)
			}
		})
		t.Run("Unmarshal", func(t *testing.T) {
			var cycleSet *sets.TinySet[string]
			err = json.Unmarshal(bs, &cycleSet)
			require.NoError(t, err)
			require.Equal(t, testSet, cycleSet)
		})
	})

}
