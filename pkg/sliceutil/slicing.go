package sliceutil

import (
	"slices"
)


// Slice between two indices. Support negative numbers >= -len to count from the end.
func Slice[TS ~[]T, T any](slice TS, from, to int) TS {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return []T{}
	}

	if from < 0 {
		from += sliceLen
	}
	if to < 0 {
		to += sliceLen + 1
	}

	from = max(0, min(from, sliceLen))
	to = min(max(from, to), sliceLen)

	if from >= to {
		return []T{}
	}

	return slice[from:to]
}

// Splice inserts the string `splice` into `s` between the specified `from` and `to` indices.
func Splice[TS ~[]T, T any](slice TS, from, to int, splice TS) TS {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return []T{}
	}

	if from < 0 {
		from += sliceLen
	}
	if to < 0 {
		to += sliceLen + 1
	}

	from = max(0, min(from, sliceLen))
	to = min(max(from, to), sliceLen)

	return slices.Concat(slice[:from], splice, slice[to:])
}
