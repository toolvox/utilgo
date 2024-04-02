// Package randutil provides various random-related helpers and utilities.
package randutil

import "math/rand"

// Derange returns a derangement of ints of size n.
func Derange(rng *rand.Rand, n int) []int {
	arr := make([]int, n)
	for i := range n {
		arr[i] = i
	}

	for i := n - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		for j == i {
			j = rng.Intn(i + 1)
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i, val := range arr {
		if i == val {
			if i == 0 {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			} else {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
	}

	return arr
}

