package api

// BasicSet is an interface for a basic implementation of a set of comparables.
type BasicSet[C comparable] interface {
	// String returns the string representation of the set.
	// Example:
	//  { (), (), () }
	String() string
	// Len counts the elements in the set.
	Len() int
	// Elements returns the unique elements of the set.
	Elements() []C
	// Add unique elements to the set.
	// Repeated elements will be discarded.
	Add(elements ...C)
	// Contains checked whether all elements are in the set.
	Contains(elements ...C) bool
	// Remove any existing elements from the set.
	Remove(elements ...C)
}

// ElementSet is an interface for a complete implementation of a set of comparables.
type ElementSet[C comparable] interface {
	BasicSet[C]

	Union(other ElementSet[C]) ElementSet[C]
	Intersection(other ElementSet[C]) ElementSet[C]
	Difference(other ElementSet[C]) ElementSet[C]

	ThreeWay(other ElementSet[C]) [3]ElementSet[C]

	UnionWith(elements ...C) ElementSet[C]
	IntersectionWith(elements ...C) ElementSet[C]
	DifferenceWith(elements ...C) ElementSet[C]
}
