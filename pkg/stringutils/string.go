package stringutils

// String represents a string as a slice of runes.
// This type allows for more direct manipulation of Unicode characters.
type String []rune

// String converts the String (slice of runes) back into a standard Go string.
func (s String) String() string {
	return string([]rune(s))
}

// Len returns the number of runes in the String, representing its length.
func (s String) Len() int {
	return len(s)
}

// Bytes converts the String into a slice of bytes.
// This is useful for operations or functions that require raw byte data.
func (s String) Bytes() []byte {
	return []byte(string(s))
}

// ByteLen returns the length of the String in bytes.
// This might differ from Len() for strings containing multi-byte characters.
func (s String) ByteLen() int {
	return len([]byte(string(s)))
}

// ByteSlice returns a new String consisting of a subset of the original String's bytes,
// starting from 'from' and ending at 'to'. It handles out-of-range indices gracefully.
func (s String) ByteSlice(from, to int) String {
	slice := safeSlice(s.Bytes(), from, to)
	return String([]rune(string(slice)))
}

// Slice returns a new String consisting of a subset of the original String's runes,
// starting from 'from' and ending at 'to'. It handles out-of-range indices gracefully.
func (s String) Slice(from, to int) String {
	return String(safeSlice([]rune(s), from, to))
}
