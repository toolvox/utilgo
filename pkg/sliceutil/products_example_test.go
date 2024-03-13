package sliceutil_test

import (
	"fmt"

	"utilgo/pkg/sliceutil"
)

func ExamplePowerSet() {
	fmt.Println("PowerSet with []int{1, 2, 3}:")
	for _, set := range sliceutil.PowerSet([]int{1, 2, 3}) {
		fmt.Println(set)
	}

	// Unordered output:
	// PowerSet with []int{1, 2, 3}:
	// []
	// [1]
	// [2]
	// [1 2]
	// [3]
	// [1 3]
	// [2 3]
	// [1 2 3]
}

func ExampleProduct() {
	fmt.Println("Product of [][]int{{1, 2}, {3, 4}, {5, 6, 7}}:")
	for _, prod := range sliceutil.Product([][]int{{1, 2}, {3, 4}, {5, 6, 7}}...) {
		fmt.Println(prod)
	}

	// Output:
	// Product of [][]int{{1, 2}, {3, 4}, {5, 6, 7}}:
	// [1 3 5]
	// [1 3 6]
	// [1 3 7]
	// [1 4 5]
	// [1 4 6]
	// [1 4 7]
	// [2 3 5]
	// [2 3 6]
	// [2 3 7]
	// [2 4 5]
	// [2 4 6]
	// [2 4 7]
}
