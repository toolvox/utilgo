package sets_test

import (
	"fmt"
	"sort"

	"github.com/toolvox/utilgo/pkg/sets"
)

func ExampleSet() {
	// Initialize a new set of integers
	s := sets.NewSet(1, 2, 3)

	// Add an element to the set
	s.Add(4)
	var resultValues []int = []int{1, 2, 3, 4}

	// Display the set
	fmt.Println("Set:", resultValues)

	// Create another set
	other := sets.NewSet(3, 4, 5)

	// UnionWith of s and other
	UnionWith := s.Union(other)
	resultValues = UnionWith.Elements()
	sort.Ints(resultValues)
	fmt.Println("UnionWith:", resultValues)

	// IntersectionWith of s and other
	IntersectionWith := s.Intersection(other)
	resultValues = IntersectionWith.Elements()
	sort.Ints(resultValues)
	fmt.Println("IntersectionWith:", resultValues)

	// DifferenceWith of s and other
	DifferenceWith := s.Difference(other)
	resultValues = DifferenceWith.Elements()
	sort.Ints(resultValues)
	fmt.Println("DifferenceWith:", resultValues)

	// Check if an element is in the set
	fmt.Println("Contains 3:", s.Contains(3))
	fmt.Println("Contains 6:", s.Contains(6))

	// Output:
	// Set: [1 2 3 4]
	// UnionWith: [1 2 3 4 5]
	// IntersectionWith: [3 4]
	// DifferenceWith: [1 2]
	// Contains 3: true
	// Contains 6: false
}
