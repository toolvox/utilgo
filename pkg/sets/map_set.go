package sets

import (
	"fmt"
	"sort"
	"strings"

	"utilgo/api"

	"maps"
)

// Unit type for the set implementation.
type Unit = api.Unit

// U is the unit value used to populate the set.
var U = api.U

// Set represents a set of elements of type T.
type Set[T comparable] map[T]Unit

// NewSet initializes a new Set with the provided elements.
func NewSet[T comparable](elements ...T) Set[T] {
	result := make(Set[T], len(elements))
	for _, e := range elements {
		result[e] = U
	}
	return result
}

// Add inserts elements into the set.
func (s Set[T]) Add(elems ...T) {
	for _, elem := range elems {
		s[elem] = U
	}
}

// SetUnion combines two sets into a new one containing elements from both.
func (s Set[T]) SetUnion(other Set[T]) Set[T] {
	result := make(Set[T], max(len(s), len(other)))
	maps.Copy(result, s)
	maps.Copy(result, other)
	return result
}

// Union adds multiple elements to the set and returns the resulting set.
func (s Set[T]) Union(elements ...T) Set[T] {
	result := make(Set[T], max(len(s), len(elements)))
	maps.Copy(result, s)
	for _, e := range elements {
		result[e] = U
	}
	return result
}

// SetIntersection creates a set of elements common to both sets.
func (s Set[T]) SetIntersection(other Set[T]) Set[T] {
	result := make(Set[T], min(len(s), len(other)))
	for k := range s {
		if _, ok := other[k]; ok {
			result[k] = U
		}
	}
	return result
}

// Intersection forms a set from common elements of the set and the provided elements.
func (s Set[T]) Intersection(elements ...T) Set[T] {
	result := make(Set[T], min(len(s), len(elements)))
	for _, k := range elements {
		if _, ok := s[k]; ok {
			result[k] = U
		}
	}
	return result
}

// SetDifference creates a set of elements in the first set but not in the second.
func (s Set[T]) SetDifference(other Set[T]) Set[T] {
	result := maps.Clone(s)
	for k := range other {
		delete(result, k)
	}
	return result
}

// Difference removes specified elements from the set.
func (s Set[T]) Difference(elements ...T) Set[T] {
	result := maps.Clone(s)
	for _, k := range elements {
		delete(result, k)
	}
	return result
}

// ThreeWay splits elements into three sets: common, only in the first set, and only in the second set.
func (s Set[T]) ThreeWay(other Set[T]) [3]Set[T] {
	blr := [3]Set[T]{NewSet[T](), NewSet[T](), NewSet[T]()}
	for k := range s {
		if other.Contains(k) {
			blr[0].Add(k)
		} else {
			blr[1].Add(k)
		}
	}

	for ok := range other {
		if s.Contains(ok) {
			continue
		}
		blr[2].Add(ok)
	}
	return blr
}

// Contains checks if the set contains the specified element.
func (s Set[T]) Contains(element T) bool {
	_, ok := s[element]
	return ok
}

// String returns a string representation of the set.
func (s Set[T]) String() string {
	var result []string
	for k := range s {
		result = append(result, fmt.Sprint(k))
	}
	sort.Strings(result)
	return fmt.Sprintf("{ %s }", strings.Join(result, ", "))
}
