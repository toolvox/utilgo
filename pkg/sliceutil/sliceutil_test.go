package sliceutil_test

import (
	"fmt"
	"reflect"
	"testing"

	"utilgo/pkg/sliceutil"

	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	// Test case for Select function
	input := []int{1, 2, 3}
	expected := []string{"1", "2", "3"}
	result := sliceutil.Select(input, func(n int) string {
		return fmt.Sprintf("%d", n)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Select: got %v, want %v", result, expected)
	}
}

func TestSelectNonZero(t *testing.T) {
	// Test case for SelectNonZero function
	input := []int{1, 0, 3}
	expected := []string{"1", "3"}
	result := sliceutil.SelectNonZero(input, func(n int) string {
		if n == 0 {
			return ""
		}
		return fmt.Sprintf("%d", n)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SelectNonZero: got %v, want %v", result, expected)
	}
}

var N = 1000

func Benchmark_Sliceutil_Prepend(b *testing.B) {
	b.Run("Half_Half", func(b *testing.B) {
		listHalf := make([]int, N, N*2)
		for i := 0; i < b.N; i++ {
			sliceutil.Prepend(listHalf, make([]int, N)...)
		}
	})
	b.Run("Full_Full", func(b *testing.B) {
		listFull := make([]int, N*2)
		for i := 0; i < b.N; i++ {
			sliceutil.Prepend(listFull, make([]int, N*2)...)
		}
	})
}


func Test_Prepend(t *testing.T) {
	testList := []int{4, 5, 6, 7}
	resList := sliceutil.Prepend(testList, 1, 2, 3)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, resList)
}
