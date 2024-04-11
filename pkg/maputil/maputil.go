// Package maputil provides utilities and helpers for scenarios not covered by the [pkg/maps] package.
package maputil

import (
	"cmp"
	"maps"
	"sort"
)

// Merge all the maps into a new map.
func Merge[M ~map[K]V, K comparable, V any](all ...M) M {
	var count int
	for _, one := range all {
		count += len(one)
	}
	var result M = make(M, count)
	for _, one := range all {
		maps.Copy(result, one)
	}
	return result
}

// Keys gets a slice of all the keys in a map, in arbitrary order.
func Keys[M ~map[K]V, K comparable, V any](self M) []K {
	var result []K = make([]K, len(self))
	id := 0
	for k := range self {
		result[id] = k
		id++
	}
	return result
}

// SortedKeys gets a slice of all the keys in a map, sorted.
func SortedKeys[M ~map[K]V, K cmp.Ordered, V any](self M) []K {
	return SortKeys(self, cmp.Less)
}

// SortKeys gets a slice of all the keys in a map, sorted by a predicate func.
func SortKeys[M ~map[K]V, K comparable, V any](self M, pred func(a, b K) bool) []K {
	result := Keys(self)
	sort.Slice(result, func(i, j int) bool {
		return pred(result[i], result[j])
	})
	return result
}

// Values gets a slice of all the values in a map, in arbitrary order.
func Values[M ~map[K]V, K comparable, V any](self M) []V {
	r := make([]V, 0, len(self))
	for _, v := range self {
		r = append(r, v)
	}
	return r
}

// SortedValues gets a slice of all the values in a map, ordered by the keys.
func SortedValues[M ~map[K]V, K cmp.Ordered, V any](self M) []V {
	return SortValues(self, cmp.Less)
}

// SortValues gets a slice of all the values in a map, ordered by the keys, sorted by a predicate func.
func SortValues[M ~map[K]V, K comparable, V any](self M, pred func(a, b K) bool) []V {
	result := make([]V, len(self))
	keys := SortKeys(self, pred)
	for i, k := range keys {
		result[i] = self[k]
	}
	return result
}
