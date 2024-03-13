package stringutil

import (
	"slices"

	"utilgo/pkg/sliceutil"
)

// String represents a string as a slice of runes.
// String typifies a string through a rune slice, facilitating direct Unicode character manipulation.
type String []rune

// String converts the [String] (slice of runes) back into a standard Go string.
func (s String) String() string {
	return string([]rune(s))
}

// Len returns the number of runes in the [String].
func (s String) Len() int {
	return len(s)
}

// Bytes converts the [String] into a slice of bytes.
// This is useful for operations or functions that require raw byte data.
func (s String) Bytes() []byte {
	return []byte(string(s))
}

// ByteLen returns the length of the [String] in bytes.
// This might differ from [String.Len]() for strings containing multi-byte characters.
func (s String) ByteLen() int {
	return len([]byte(string(s)))
}

// ByteSlice returns a new [String] consisting of a subset of the original [String]'s bytes, starting from 'from' and ending at 'to'.
// It handles out-of-range indices gracefully.
func (s String) ByteSlice(from, to int) String {
	slice := sliceutil.Slice(s.Bytes(), from, to)
	return String(string(slice))
}

// Slice returns a new [String] consisting of a subset of the original [String]'s runes, starting from 'from' and ending at 'to'.
// It handles out-of-range indices gracefully.
func (s String) Slice(from, to int) String {
	return sliceutil.Slice(s, from, to)
}

// ByteSplice returns a new [String] consisting of a splicing of the `splice` [String] into original [String]'s bytes, starting from 'from' and ending at 'to'.
// It handles out-of-range indices gracefully.
func (s String) ByteSplice(from, to int, splice String) String {
	slice := sliceutil.Splice(s.Bytes(), from, to, splice.Bytes())
	return String(string(slice))
}

// Splice returns a new [String] consisting of a splicing of the `splice` [String] into original [String]'s runes, starting from 'from' and ending at 'to'.
// It handles out-of-range indices gracefully.
func (s String) Splice(from, to int, splice String) String {
	return sliceutil.Splice([]rune(s), from, to, splice)
}

// Append the runes of another [String] to this one.
func (s String) Append(runes String) String {
	return append(s, runes...)
}

// CountPrefix returns the number of runes, starting from the 0th, that return true for isPrefix.
func (s String) CountPrefix(isPrefix func(rune) bool) int {
	for i, r := range s {
		if !isPrefix(r) {
			return i
		}
	}
	return len(s)
}

// Equals returns true when both [String]s have the same runes.
func (s String) Equals(other String) bool {
	return slices.Equal(s, other)
}

// Clone makes a deep copy of the [String].
func (s String) Clone() String {
	var res String = make(String, s.Len())
	copy(res, s)
	return res
}
