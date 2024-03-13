package sliceutil

import (
	"slices"
	"sort"
)

// SliceCropIndex removes an element from a slice at a specified index and returns the modified slice.
// If the index is out of bounds, the original slice is returned unmodified.
// Original slice may be modified by this operation.
func SliceCropIndex[S ~[]E, E any](slice S, i int) S {
	switch {
	case i < 0 || i >= len(slice):
		return slice

	case len(slice) == 1 && i == 0:
		return S{}

	case i == 0:
		return slice[1:]

	case i == len(slice)-1:
		return slice[:i]
	}

	left := slice[:i]
	right := slice[i+1:]
	return append(left, right...)
}

// CropIndex removes an element from a slice at a specified index and returns the modified slice.
// It ensures the slice is clipped to remove any potential nil or zero-value elements at the end.
// This function clones the original slice before modification, preserving the input slice's integrity.
func CropIndex[S ~[]E, E any](slice S, index int) S {
	return slices.Clip(SliceCropIndex(slices.Clone(slice), index))
}

// CropIndices removes elements from a slice at specified indices and returns the modified slice.
// Indices are processed in reverse order to maintain correct indexing after each removal.
// Like CropIndex, it clones the original slice before modification and clips the result.
func CropIndices[S ~[]E, E any](slice S, indices ...int) S {
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))
	slice = slices.Clone(slice)
	for _, i := range indices {
		slice = SliceCropIndex(slice, i)
	}

	return slices.Clip(slice)
}

// CropElement removes all occurrences of a specified element from a slice and returns the modified slice.
// It searches for the element using linear search and removes each occurrence found.
// This function clones the slice before modification and clips the result to ensure integrity.
func CropElement[S ~[]C, C comparable](slice S, element C) S {
	slice = slices.Clone(slice)
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			slice = SliceCropIndex(slice, i)
		}
	}
	return slices.Clip(slice)
}

// CropElements removes all occurrences of any of the specified elements from a slice and returns the modified slice.
// Each element in the input slice is checked against the specified elements, and matches are removed.
// The function clones the original slice before modification and ensures the result is clipped.
func CropElements[S ~[]C, C comparable](slice S, elements ...C) S {
	slice = slices.Clone(slice)
	for i := 0; i < len(slice); i++ {
		if slices.Contains(elements, slice[i]) {
			slice = SliceCropIndex(slice, i)
			i--
		}
	}
	return slices.Clip(slice)
}
