// Package api provides the interfaces which the structs in the library implement, and common type definitions.
package api

// Unit is an empty struct. An efficient way to represent nothing.
type Unit = struct{}

// U is the unit. The single instance of the type nothing.
var U = Unit{}

type BasicSet[TCmp comparable] interface {
	Add(elems ...TCmp)
	Contains(element TCmp) bool
	String() string
}

type Set[TCmp comparable] interface {
	BasicSet[TCmp]

	Union(other Set[TCmp]) Set[TCmp]
	Intersection(other Set[TCmp]) Set[TCmp]
	Difference(other Set[TCmp]) Set[TCmp]
	ThreeWay(other Set[TCmp]) [3]Set[TCmp]
}

type ElementSet[TCmp comparable] interface {
	BasicSet[TCmp]
	UnionWith(elements ...TCmp) Set[TCmp]
	IntersectionWith(elements ...TCmp) Set[TCmp]
	DifferenceWith(elements ...TCmp) Set[TCmp]
}
