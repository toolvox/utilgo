package sets

import (
	"fmt"
	"maps"
	"sort"
	"strings"

	"utilgo/api"
)

// Unit type for the set implementation.
type Unit = api.Unit

// U is the unit value used to populate the set.
var U = api.U

// Set represents a set of elements of type T.
type Set[TCmp comparable] map[TCmp]Unit

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

// Contains checks if the set contains the specified element.
func (s Set[T]) Contains(element T) bool {
	_, ok := s[element]
	return ok
}

// Union combines two sets into a new one containing elements from both.
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T], max(len(s), len(other)))
	maps.Copy(result, s)
	maps.Copy(result, other)
	return result
}

// UnionWith adds multiple elements to the set and returns the resulting set.
func (s Set[T]) UnionWith(elements ...T) Set[T] {
	result := make(Set[T], max(len(s), len(elements)))
	maps.Copy(result, s)
	for _, e := range elements {
		result[e] = U
	}
	return result
}

// Intersection creates a set of elements common to both sets.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T], min(len(s), len(other)))
	for k := range s {
		if _, ok := other[k]; ok {
			result[k] = U
		}
	}
	return result
}

// IntersectionWith forms a set from common elements of the set and the provided elements.
func (s Set[T]) IntersectionWith(elements ...T) Set[T] {
	result := make(Set[T], min(len(s), len(elements)))
	for _, k := range elements {
		if _, ok := s[k]; ok {
			result[k] = U
		}
	}
	return result
}

// Difference creates a set of elements in the first set but not in the second.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := maps.Clone(s)
	for k := range other {
		delete(result, k)
	}
	return result
}

// DifferenceWith removes specified elements from the set.
func (s Set[T]) DifferenceWith(elements ...T) Set[T] {
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

// String returns a string representation of the set.
func (s Set[T]) String() string {
	var result []string
	for k := range s {
		result = append(result, fmt.Sprint(k))
	}
	sort.Strings(result)
	return fmt.Sprintf("{ %s }", strings.Join(result, ", "))
}
