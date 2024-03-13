// Package maputil provides utilities and helpers for scenarios not covered by the [pkg/maps] package.
package maputil

import (
	"cmp"
	"sort"
)

// SortedKeys gets a slice of all the keys in a map, sorted
func SortedKeys[M ~map[K]V, K cmp.Ordered, V any](self M) []K {
	var result []K = make([]K, len(self))
	id := 0
	for k := range self {
		result[id] = k
		id++
	}

	sort.Slice(result, func(i, j int) bool {
		return cmp.Less(result[i], result[j])
	})
	return result
}
