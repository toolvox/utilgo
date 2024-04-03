// Package sets provides several implementation of a Set and related helpers.
package sets

import (
	"fmt"
	"strings"

	"github.com/toolvox/utilgo/pkg/sliceutil"
)

// TinySet represents a set of elements of type C using a slice for storage.
// It is optimized for small sets where the overhead of a map is not justified.
type TinySet[C comparable] []C

// NewTinySet initializes a new [TinySet] with the given elements, ensuring uniqueness.
func NewTinySet[C comparable](elements ...C) TinySet[C] {
	result := TinySet[C]{}
	for _, e := range elements {
		result.Add(e)
	}
	return result
}

// String provides a string representation of the [TinySet].
func (tinySet TinySet[C]) String() string {
	var elems []string
	for _, elem := range tinySet {
		elems = append(elems, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("{%s}", strings.Join(elems, ", "))
}

// Len counts the elements in the [TinySet].
func (tinySet TinySet[C]) Len() int { return len(tinySet) }

// Elements returns the contained elements.
func (tinySet TinySet[C]) Elements() []C { return tinySet }

// Contains checks if an element exists in the [TinySet].
func (tinySet TinySet[C]) Contains(elements ...C) bool {
outer:
	for _, e := range elements {
		for _, i := range tinySet {
			if i == e {
				continue outer
			}
		}
		return false
	}
	return true
}

// Add inserts new unique elements into the [TinySet].
func (tinySet *TinySet[C]) Add(elements ...C) {
	for _, elem := range elements {
		if !tinySet.Contains(elem) {
			*tinySet = append(*tinySet, elem)
		}
	}
}

// Remove elements from the [TinySet].
func (tinySet *TinySet[C]) Remove(elements ...C) {
	*tinySet = sliceutil.CropElements(*tinySet, elements...)
}

// Union merges the current set with another set and returns the result as a new [TinySet].
func (tinySet TinySet[C]) Union(other TinySet[C]) TinySet[C] {
	result := NewTinySet(tinySet...) // Start with a copy of the current set
	for _, elem := range other {
		if !result.Contains(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// Intersection returns a new [TinySet] containing elements that exist in both sets.
func (tinySet TinySet[C]) Intersection(other TinySet[C]) TinySet[C] {
	var result TinySet[C]
	for _, elem := range tinySet {
		if other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Difference returns a new [TinySet] containing elements present in this set but not in the other set.
func (tinySet TinySet[C]) Difference(other TinySet[C]) TinySet[C] {
	var result TinySet[C]
	for _, elem := range tinySet {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// ThreeWay splits elements into three [TinySets]: common to both, only in the first, and only in the second.
func (tinySet TinySet[C]) ThreeWay(other TinySet[C]) [3]TinySet[C] {
	return [3]TinySet[C]{
		tinySet.Intersection(other),
		tinySet.Difference(other),
		other.Difference(tinySet),
	}
}

// UnionWith adds multiple elements to the set and returns the resulting set.
// This operation does not modify the original set but returns a new one.
func (tinySet TinySet[C]) UnionWith(elements ...C) TinySet[C] {
	result := NewTinySet(tinySet...) // Start with a copy of the current set
	for _, elem := range elements {
		if !result.Contains(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// IntersectionWith returns a new [TinySet] containing elements that exist in both this set and the provided elements.
func (tinySet TinySet[C]) IntersectionWith(elements ...C) TinySet[C] {
	other := NewTinySet(elements...)
	return tinySet.Intersection(other)
}

// DifferenceWith returns a new [TinySet] containing elements present in this set but not among the provided elements.
func (tinySet TinySet[C]) DifferenceWith(elements ...C) TinySet[C] {
	other := NewTinySet(elements...)
	return tinySet.Difference(other)
}
