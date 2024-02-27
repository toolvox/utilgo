package sets

import (
	"fmt"
	"strings"
)

// TinySet represents a set of elements of type TCmp using a slice for storage.
// It is optimized for small sets where the overhead of a map is not justified.
type TinySet[TCmp comparable] []TCmp

// NewTinySet initializes a new [TinySet] with the given elements, ensuring uniqueness.
func NewTinySet[TCmp comparable](elements ...TCmp) TinySet[TCmp] {
	ts := TinySet[TCmp]{}
	for _, elem := range elements {
		if !ts.Contains(elem) {
			ts = append(ts, elem)
		}
	}
	return ts
}

// Add inserts new unique elements into the [TinySet].
func (ts *TinySet[TCmp]) Add(elements ...TCmp) {
	for _, elem := range elements {
		if !ts.Contains(elem) {
			*ts = append(*ts, elem)
		}
	}
}

// Contains checks if an element exists in the [TinySet].
func (ts TinySet[TCmp]) Contains(elem TCmp) bool {
	for _, e := range ts {
		if e == elem {
			return true
		}
	}
	return false
}

// Union merges the current set with another set and returns the result as a new [TinySet].
func (ts TinySet[TCmp]) Union(other TinySet[TCmp]) TinySet[TCmp] {
	result := NewTinySet(ts...) // Start with a copy of the current set
	for _, elem := range other {
		if !result.Contains(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// UnionWith adds multiple elements to the set and returns the resulting set.
// This operation does not modify the original set but returns a new one.
func (ts TinySet[TCmp]) UnionWith(elements ...TCmp) TinySet[TCmp] {
	result := NewTinySet(ts...) // Start with a copy of the current set
	for _, elem := range elements {
		if !result.Contains(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// Intersection returns a new [TinySet] containing elements that exist in both sets.
func (ts TinySet[TCmp]) Intersection(other TinySet[TCmp]) TinySet[TCmp] {
	var result TinySet[TCmp]
	for _, elem := range ts {
		if other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// IntersectionWith returns a new [TinySet] containing elements that exist in both this set and the provided elements.
func (ts TinySet[TCmp]) IntersectionWith(elements ...TCmp) TinySet[TCmp] {
	other := NewTinySet(elements...)
	return ts.Intersection(other)
}

// Difference returns a new [TinySet] containing elements present in this set but not in the other set.
func (ts TinySet[TCmp]) Difference(other TinySet[TCmp]) TinySet[TCmp] {
	var result TinySet[TCmp]
	for _, elem := range ts {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// DifferenceWith returns a new [TinySet] containing elements present in this set but not among the provided elements.
func (ts TinySet[TCmp]) DifferenceWith(elements ...TCmp) TinySet[TCmp] {
	other := NewTinySet(elements...)
	return ts.Difference(other)
}

// ThreeWay splits elements into three [TinySets]: common to both, only in the first, and only in the second.
func (ts TinySet[TCmp]) ThreeWay(other TinySet[TCmp]) [3]TinySet[TCmp] {
	common := ts.Intersection(other)
	onlyInTs := ts.Difference(other)
	onlyInOther := other.Difference(ts)
	return [3]TinySet[TCmp]{common, onlyInTs, onlyInOther}
}

// String provides a string representation of the [TinySet].
func (ts TinySet[TCmp]) String() string {
	var elems []string
	for _, elem := range ts {
		elems = append(elems, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("{%s}", strings.Join(elems, ", "))
}
