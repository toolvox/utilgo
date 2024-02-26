// Package stringutils provides utilities for string manipulation,
// focusing on operations with runes and bytes for Unicode strings.
package stringutils

// safeSlice is a generic helper function to safeSlice any safeSlice type TS,
// from a starting index 'from' to an ending index 'to',
// handling negative and out-of-range indices.
func safeSlice[TS ~[]T, T any](slice TS, from, to int) TS {
	if from < 0 {
		from = 0
	}
	if to > len(slice) || to < 0 {
		to = len(slice)
	}
	if from >= to {
		return TS{}
	}
	return slice[from:to]
}
