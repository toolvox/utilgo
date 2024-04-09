package sets

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/toolvox/utilgo/api"
)

// Unit type for the set implementation.
type Unit = api.Unit

// U is the unit value used to populate the set.
var U = api.U

// Set uses a go map to represent a set of elements of type C as keys with no value.
type Set[C comparable] map[C]Unit

// NewSet initializes a new [Set] with the given elements, ensuring uniqueness.
func NewSet[C comparable](elements ...C) Set[C] {
	result := make(Set[C], len(elements))
	for _, e := range elements {
		result[e] = U
	}
	return result
}

// String returns the string representation of the [Set].
func (set Set[C]) String() string {
	var sb strings.Builder
	sb.WriteRune('{')
	var notFirst bool
	for e := range set {
		if notFirst {
			sb.WriteRune(',')
		} else {
			notFirst = true
		}
		sb.WriteRune(' ')
		fmt.Fprint(&sb, e)
	}
	sb.WriteRune(' ')
	sb.WriteRune('}')
	return sb.String()
}

// Len counts the elements in the [Set].
func (set Set[C]) Len() int { return len(set) }

// Elements returns the unique elements of the [Set].
func (set Set[C]) Elements() []C {
	var elements []C = make([]C, 0, set.Len())
	for e := range set {
		elements = append(elements, e)
	}
	return elements
}

// Add unique elements to the [Set].
// Repeated elements will be discarded.
func (set Set[C]) Add(elements ...C) {
	for _, elem := range elements {
		set[elem] = U
	}
}

// Contains checked whether all elements are in the [Set].
func (set Set[C]) Contains(elements ...C) bool {
outer:
	for _, element := range elements {
		if _, ok := set[element]; ok {
			continue outer
		}
		return false
	}
	return true
}

// Remove any existing elements from the [Set].
func (set Set[C]) Remove(elements ...C) {
	for _, e := range elements {
		delete(set, e)
	}
}

// Union combines two sets into a new one containing elements from both.
func (set Set[C]) Union(other Set[C]) Set[C] {
	result := make(Set[C], max(len(set), len(other)))
	maps.Copy(result, set)
	maps.Copy(result, other)
	return result
}

// UnionWith adds multiple elements to the set and returns the resulting set.
func (set Set[C]) UnionWith(elements ...C) Set[C] {
	result := make(Set[C], max(len(set), len(elements)))
	maps.Copy(result, set)
	for _, e := range elements {
		result[e] = U
	}
	return result
}

// Intersection creates a set of elements common to both sets.
func (set Set[C]) Intersection(other Set[C]) Set[C] {
	result := make(Set[C], min(len(set), len(other)))
	for k := range set {
		if _, ok := other[k]; ok {
			result[k] = U
		}
	}
	return result
}

// IntersectionWith forms a set from common elements of the set and the provided elements.
func (set Set[C]) IntersectionWith(elements ...C) Set[C] {
	result := make(Set[C], min(len(set), len(elements)))
	for _, k := range elements {
		if _, ok := set[k]; ok {
			result[k] = U
		}
	}
	return result
}

// Difference creates a set of elements in the first set but not in the second.
func (set Set[C]) Difference(other Set[C]) Set[C] {
	result := maps.Clone(set)
	for k := range other {
		delete(result, k)
	}
	return result
}

// DifferenceWith removes specified elements from the set.
func (set Set[C]) DifferenceWith(elements ...C) Set[C] {
	result := maps.Clone(set)
	for _, k := range elements {
		delete(result, k)
	}
	return result
}

// ThreeWay splits elements into three sets: common, only in the first set, and only in the second set.
func (set Set[C]) ThreeWay(other Set[C]) [3]Set[C] {
	blr := [3]Set[C]{NewSet[C](), NewSet[C](), NewSet[C]()}
	for k := range set {
		if other.Contains(k) {
			blr[0].Add(k)
		} else {
			blr[1].Add(k)
		}
	}

	for ok := range other {
		if set.Contains(ok) {
			continue
		}
		blr[2].Add(ok)
	}
	return blr
}

// MarshalJSON converts the [Set] to a JSON array.
func (set Set[T]) MarshalJSON() ([]byte, error) {
	slice := make([]T, 0, len(set))
	for key := range set {
		slice = append(slice, key)
	}

	return json.Marshal(slice)
}

// UnmarshalJSON converts the JSON []byte to a [Set].
func (set *Set[T]) UnmarshalJSON(data []byte) error {
	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	*set = NewSet(slice...)
	return nil
}

// MarshalYAML converts the [Set] to a YAML array.
func (set Set[T]) MarshalYAML() (interface{}, error) {
	slice := make([]T, 0, len(set))
	for key := range set {
		slice = append(slice, key)
	}
	return slice, nil
}

// UnmarshalYAML converts the YAML [yaml.Node] (representing a YAML list) to a [Set].
func (set *Set[T]) UnmarshalYAML(value *yaml.Node) error {
	var slice []T
	err := value.Decode(&slice)
	if err != nil {
		return err
	}
	if *set == nil {
		n := NewSet[T](slice...)
		*set = n
	} else {
		set.Add(slice...)
	}
	return nil
}
