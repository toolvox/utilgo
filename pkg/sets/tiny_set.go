package sets

import (
	"fmt"
	"strings"
)

// TinySet represents a set of elements of type T using a slice for storage.
// It is optimized for small sets where the overhead of a map is not justified.
type TinySet[T comparable] []T

// NewTinySet initializes a new TinySet with the given elements, ensuring uniqueness.
func NewTinySet[T comparable](elements ...T) TinySet[T] {
	ts := TinySet[T]{}
	for _, elem := range elements {
		if !ts.Contains(elem) {
			ts = append(ts, elem)
		}
	}
	return ts
}

// Add inserts new unique elements into the TinySet.
func (ts *TinySet[T]) Add(elements ...T) {
	for _, elem := range elements {
		if !ts.Contains(elem) {
			*ts = append(*ts, elem)
		}
	}
}

// Contains checks if an element exists in the TinySet.
func (ts TinySet[T]) Contains(elem T) bool {
	for _, e := range ts {
		if e == elem {
			return true
		}
	}
	return false
}

// SetUnion merges the current set with another set and returns the result as a new TinySet.
func (ts TinySet[T]) SetUnion(other TinySet[T]) TinySet[T] {
	result := NewTinySet(ts...) // Start with a copy of the current set
	for _, elem := range other {
		if !result.Contains(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// Union adds multiple elements to the set and returns the resulting set.
// This operation does not modify the original set but returns a new one.
func (ts TinySet[T]) Union(elements ...T) TinySet[T] {
	result := NewTinySet(ts...) // Start with a copy of the current set
	for _, elem := range elements {
		if !result.Contains(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// SetIntersection returns a new TinySet containing elements that exist in both sets.
func (ts TinySet[T]) SetIntersection(other TinySet[T]) TinySet[T] {
	var result TinySet[T]
	for _, elem := range ts {
		if other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Intersection returns a new TinySet containing elements that exist in both this set and the provided elements.
func (ts TinySet[T]) Intersection(elements ...T) TinySet[T] {
	other := NewTinySet(elements...)
	return ts.SetIntersection(other)
}

// SetDifference returns a new TinySet containing elements present in this set but not in the other set.
func (ts TinySet[T]) SetDifference(other TinySet[T]) TinySet[T] {
	var result TinySet[T]
	for _, elem := range ts {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Difference returns a new TinySet containing elements present in this set but not among the provided elements.
func (ts TinySet[T]) Difference(elements ...T) TinySet[T] {
	other := NewTinySet(elements...)
	return ts.SetDifference(other)
}

// ThreeWay splits elements into three TinySets: common to both, only in the first, and only in the second.
func (ts TinySet[T]) ThreeWay(other TinySet[T]) [3]TinySet[T] {
	common := ts.SetIntersection(other)
	onlyInTs := ts.SetDifference(other)
	onlyInOther := other.SetDifference(ts)
	return [3]TinySet[T]{common, onlyInTs, onlyInOther}
}

// String provides a string representation of the TinySet.
func (ts TinySet[T]) String() string {
	var elems []string
	for _, elem := range ts {
		elems = append(elems, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("{%s}", strings.Join(elems, ", "))
}
