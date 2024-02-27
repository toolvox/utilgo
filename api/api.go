// Package api defines interfaces implemented by library structs and common types.
package api

// Unit represents an empty struct for efficient representation of nothing.
type Unit = struct{}

// U is the unit instance, representing the single instance of the type nothing.
var U = Unit{}

// BasicSet defines basic set operations for elements of comparable type TCmp.
type BasicSet[TCmp comparable] interface {
	Add(elems ...TCmp)
	Contains(element TCmp) bool
	String() string
}

// Set extends BasicSet with union, intersection, and difference operations for sets of comparable type TCmp.
type Set[TCmp comparable] interface {
	BasicSet[TCmp]

	Union(other Set[TCmp]) Set[TCmp]
	Intersection(other Set[TCmp]) Set[TCmp]
	Difference(other Set[TCmp]) Set[TCmp]
	ThreeWay(other Set[TCmp]) [3]Set[TCmp]
}

// ElementSet extends BasicSet with operations to union, intersect, and differentiate with individual elements.
type ElementSet[TCmp comparable] interface {
	BasicSet[TCmp]
	UnionWith(elements ...TCmp) Set[TCmp]
	IntersectionWith(elements ...TCmp) Set[TCmp]
	DifferenceWith(elements ...TCmp) Set[TCmp]
}
