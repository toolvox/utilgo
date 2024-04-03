package sets

import (
	"fmt"
	"maps"
	"sort"
	"strings"

	"github.com/toolvox/utilgo/api"
)

// Unit type for the set implementation.
type Unit = api.Unit

// U is the unit value used to populate the set.
var U = api.U

// Set represents a set of elements of type C.
type Set[C comparable] map[C]Unit

// NewSet initializes a new Set with the provided elements.
func NewSet[C comparable](elements ...C) Set[C] {
	result := make(Set[C], len(elements))
	for _, e := range elements {
		result[e] = U
	}
	return result
}

// Add inserts elements into the set.
func (s Set[C]) Add(elems ...C) {
	for _, elem := range elems {
		s[elem] = U
	}
}

// Contains checks if the set contains the specified element.
func (s Set[C]) Contains(element C) bool {
	_, ok := s[element]
	return ok
}

// Union combines two sets into a new one containing elements from both.
func (s Set[C]) Union(other Set[C]) Set[C] {
	result := make(Set[C], max(len(s), len(other)))
	maps.Copy(result, s)
	maps.Copy(result, other)
	return result
}

// UnionWith adds multiple elements to the set and returns the resulting set.
func (s Set[C]) UnionWith(elements ...C) Set[C] {
	result := make(Set[C], max(len(s), len(elements)))
	maps.Copy(result, s)
	for _, e := range elements {
		result[e] = U
	}
	return result
}

// Intersection creates a set of elements common to both sets.
func (s Set[C]) Intersection(other Set[C]) Set[C] {
	result := make(Set[C], min(len(s), len(other)))
	for k := range s {
		if _, ok := other[k]; ok {
			result[k] = U
		}
	}
	return result
}

// IntersectionWith forms a set from common elements of the set and the provided elements.
func (s Set[C]) IntersectionWith(elements ...C) Set[C] {
	result := make(Set[C], min(len(s), len(elements)))
	for _, k := range elements {
		if _, ok := s[k]; ok {
			result[k] = U
		}
	}
	return result
}

// Difference creates a set of elements in the first set but not in the second.
func (s Set[C]) Difference(other Set[C]) Set[C] {
	result := maps.Clone(s)
	for k := range other {
		delete(result, k)
	}
	return result
}

// DifferenceWith removes specified elements from the set.
func (s Set[C]) DifferenceWith(elements ...C) Set[C] {
	result := maps.Clone(s)
	for _, k := range elements {
		delete(result, k)
	}
	return result
}

// ThreeWay splits elements into three sets: common, only in the first set, and only in the second set.
func (s Set[C]) ThreeWay(other Set[C]) [3]Set[C] {
	blr := [3]Set[C]{NewSet[C](), NewSet[C](), NewSet[C]()}
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
func (s Set[C]) String() string {
	var result []string
	for k := range s {
		result = append(result, fmt.Sprint(k))
	}
	sort.Strings(result)
	return fmt.Sprintf("{ %s }", strings.Join(result, ", "))
}
