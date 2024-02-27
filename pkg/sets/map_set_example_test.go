package sets_test

import (
	"fmt"

	"utilgo/pkg/sets"
)

func ExampleSet() {
	// Initialize a new set of integers
	s := sets.NewSet(1, 2, 3)

	// Add an element to the set
	s.Add(4)

	// Display the set
	fmt.Println("Set:", s)

	// Create another set
	other := sets.NewSet(3, 4, 5)

	// UnionWith of s and other
	UnionWith := s.Union(other)
	fmt.Println("UnionWith:", UnionWith)

	// IntersectionWith of s and other
	IntersectionWith := s.Intersection(other)
	fmt.Println("IntersectionWith:", IntersectionWith)

	// DifferenceWith of s and other
	DifferenceWith := s.Difference(other)
	fmt.Println("DifferenceWith:", DifferenceWith)

	// Check if an element is in the set
	fmt.Println("Contains 3:", s.Contains(3))
	fmt.Println("Contains 6:", s.Contains(6))

	// Output:
	// Set: { 1, 2, 3, 4 }
	// UnionWith: { 1, 2, 3, 4, 5 }
	// IntersectionWith: { 3, 4 }
	// DifferenceWith: { 1, 2 }
	// Contains 3: true
	// Contains 6: false
}
