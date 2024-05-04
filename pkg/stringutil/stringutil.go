// Package stringutil provides utilities for string manipulation,
// focusing on operations with runes and bytes for Unicode strings.
package stringutil

import (
	"fmt"
	"strings"
)

// CountPrefix returns the number of runes, starting from the 0th, that return true for isPrefix.
func CountPrefix(s string, isPrefix func(rune) bool) int {
	for i, r := range s {
		if !isPrefix(r) {
			return i
		}
	}
	return len(s)
}

func EnsurePrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s
	}
	return fmt.Sprint(prefix, s)
}
