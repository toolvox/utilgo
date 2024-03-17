package sliceutil

import (
	"slices"

	"utilgo/pkg/reflectutil"
)

// Select transforms each element of the input slice using the provided selector function and returns a new slice of the transformed elements.
//
// E1 is the element type of the input slice and should be a slice type.
//
// E2 is the element type of the resulting slice and should be a slice type.
//
// T1 is the type of elements in the input slice.
//
// T2 is the type of elements in the resulting slice.
func Select[E1 ~[]T1, E2 []T2, T1, T2 any](slice E1, selector func(T1) T2) E2 {
	var res E2 = make(E2, len(slice))
	for i, v := range slice {
		res[i] = selector(v)
	}
	return res
}

// SelectNonZero transforms each element of the input slice using the provided selector function, then filters out any zero values from the resulting slice.
// It leverages the Select function to perform the transformation and then uses [pkg/slices.DeleteFunc] from the slices package to remove zero values, with zero-ness determined by the [pkg/utilgo/pkg/reflectutil.IsZero] function.
func SelectNonZero[E1 ~[]T1, E2 []T2, T1, T2 any](slice E1, selector func(T1) T2) E2 {
	return slices.DeleteFunc(
		Select[E1, E2](slice, selector),
		reflectutil.IsZero)
}

// Prepend inserts the items before the list.
func Prepend[E ~[]T, T any](list E, items ...T) E {
	countItems := len(items)
	list = append(list, make(E, countItems)...)
	copy(list[countItems:], list)
	copy(list, items)
	return list
}
