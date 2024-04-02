package sliceutil

import (
	"slices"

	"utilgo/pkg/reflectutil"
)

// SelectFunc transforms each element of the input slice using the provided selector function and returns a new slice of the transformed elements.
//
// E1 is the element type of the input slice and should be a slice type.
//
// E2 is the element type of the resulting slice and should be a slice type.
//
// T1 is the type of elements in the input slice.
//
// T2 is the type of elements in the resulting slice.
func SelectFunc[E1 ~[]T1, E2 []T2, T1, T2 any](slice E1, selector func(T1) T2) E2 {
	var res E2 = make(E2, len(slice))
	for i, v := range slice {
		res[i] = selector(v)
	}
	return res
}

// SelectNonZeroFunc transforms each element of the input slice using the provided selector function, then filters out any zero values from the resulting slice.
// It leverages the Select function to perform the transformation and then uses [pkg/slices.DeleteFunc] from the slices package to remove zero values, with zero-ness determined by the [pkg/utilgo/pkg/reflectutil.IsZero] function.
func SelectNonZeroFunc[E1 ~[]T1, E2 []T2, T1, T2 any](slice E1, selector func(T1) T2) E2 {
	return slices.DeleteFunc(
		SelectFunc[E1, E2](slice, selector),
		reflectutil.IsZero)
}

// Prepend inserts the items before the slice.
func Prepend[E ~[]T, T any](slice E, elements ...T) E {
	countItems := len(elements)
	slice = append(slice, make(E, countItems)...)
	copy(slice[countItems:], slice)
	copy(slice, elements)
	return slice
}

// Interleave inserts the element into the slice every step elements.
func Interleave[E ~[]T, T any](slice E, element T, step int) E {
	if step <= 0 {
		step = 1
	}

	tail := len(slice)
	if step >= tail {
		return slice
	}

	extras := tail/step - 1
	slice = append(slice, make([]T, extras)...)
	end := tail + extras

	for tail > 0 && tail < end {
		copy(slice[end-step:end], slice[tail-step:tail])
		end -= (step + 1)
		tail -= step
		slice[end] = element
	}

	return slice
}

// IndexFuncs returns the first index i satisfying f(s[i]) for any f in fs, or nil if none do.
func IndexFuncs[E ~[]T, T any](slice E, funcs ...func(T) bool) []int {
	var indices []int
	for i := range slice {
		for _, f := range funcs {
			if f(slice[i]) {
				indices = append(indices, i)
				break
			}
		}
	}

	return indices
}

// DeleteFuncs removes any elements from s for which any of the funcs returns true, returning the modified slice.
// DeleteFuncs zeroes the elements between the new length and the original length.
func DeleteFuncs[E ~[]T, T any](slice E, funcs ...func(T) bool) E {
    indices := IndexFuncs(slice, funcs...)
    if len(indices) == 0 {
        return slice
    }

    delCount := 0
    for i, idx := range indices {
        adjIdx := idx - i
        copy(slice[adjIdx:], slice[adjIdx+1:])
        delCount++
    }

    newLength := len(slice) - delCount
    clear(slice[newLength:])
    return slice[:newLength]
}