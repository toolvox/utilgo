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

	// Union of s and other
	union := s.SetUnion(other)
	fmt.Println("Union:", union)

	// Intersection of s and other
	intersection := s.SetIntersection(other)
	fmt.Println("Intersection:", intersection)

	// Difference of s and other
	difference := s.SetDifference(other)
	fmt.Println("Difference:", difference)

	// Check if an element is in the set
	fmt.Println("Contains 3:", s.Contains(3))
	fmt.Println("Contains 6:", s.Contains(6))

	// Output:
	// Set: { 1, 2, 3, 4 }
	// Union: { 1, 2, 3, 4, 5 }
	// Intersection: { 3, 4 }
	// Difference: { 1, 2 }
	// Contains 3: true
	// Contains 6: false
}
