package sliceutil_test

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"github.com/toolvox/utilgo/pkg/reflectutil"
	"github.com/toolvox/utilgo/pkg/sliceutil"

	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	// Test case for Select function
	input := []int{1, 2, 3}
	expected := []string{"1", "2", "3"}
	result := sliceutil.SelectFunc(input, func(n int) string {
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
	result := sliceutil.SelectNonZeroFunc(input, func(n int) string {
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

func Test_Interleave(t *testing.T) {
	tests := []struct {
		list    []int
		element int
		step    int
		expect  []int
	}{
		{
			list:    []int{1, 2, 3},
			element: 9,
			step:    0,
			expect:  []int{1, 9, 2, 9, 3},
		},
		{
			list:    []int{1, 2, 3},
			element: 9,
			step:    1,
			expect:  []int{1, 9, 2, 9, 3},
		},
		{
			list:    []int{1, 2, 3, 4},
			element: 5,
			step:    2,
			expect:  []int{1, 2, 5, 3, 4},
		},
	}

	for ti, tt := range tests {
		t.Run(fmt.Sprint(ti), func(t *testing.T) {
			actual := sliceutil.Interleave(tt.list, tt.element, tt.step)
			require.Equal(t, tt.expect, actual)
		})
	}
}

func Benchmark_IndexFuncs(b *testing.B) {
	b.Run("___int", func(b *testing.B) {
		RunBench_IndexFuncs(b, 1, 2,
			func(i int) bool { return i == 2 },
			func(i int) bool { return i > 1 },
			func(i int) bool { return i%2 == 0 },
		)
	})
	b.Run("string", func(b *testing.B) {
		RunBench_IndexFuncs(b, "keep", "drop",
			func(s string) bool { return s == "drop" },
			func(s string) bool { return len(s) == 4 && s[0] == 'd' },
			func(s string) bool { return s != "keep" },
		)
	})
	type test struct {
		Id   int
		Name string
		Test *test
	}
	testKeep := test{
		Id:   123,
		Name: "Asd.",
		Test: nil,
	}
	testDrop := test{
		Id:   -321,
		Name: ".dsA",
		Test: &testKeep,
	}
	b.Run("__test", func(b *testing.B) {
		RunBench_IndexFuncs(b, testKeep, testDrop,
			func(t test) bool { return t.Id != 123 && t.Test == &testKeep },
			func(t test) bool { return t.Name != "Asd." && t.Test != nil },
			func(t test) bool { return t.Name == ".dsA" || t.Test != nil },
		)
	})
}

func RunBench_IndexFuncs[T any](b *testing.B, falseE, trueE T, extraFuncs ...func(T) bool) {
	for n := 2; n <= len(extraFuncs)+1; n++ {
		funcs := sliceutil.Prepend(extraFuncs, reflectutil.IsZero[T])[:n]
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for size := 10; size <= 100000; size *= 100 {
				b.Run(fmt.Sprintf("%7d", size), func(b *testing.B) {
					list := make([]T, size)
					for i := range size {
						if i%2 == 0 {
							list[i] = trueE
							continue
						}
						list[i] = falseE
					}

					b.Run("v1", func(b *testing.B) {
						b.ResetTimer()
						var lll []int
						for i := 0; i < b.N; i++ {
							lll = sliceutil.IndexFuncs(list, funcs...)
						}
						b.StopTimer()
						if !slices.IsSorted(lll) {
							panic("s1")
						}
						if len(lll) != size/2 {
							panic("v1")
						}
					})

				})
			}
		})
	}
}

// Assume Interleave and Interleave2 are defined here or in an imported package.
func BenchmarkInterleaveInts(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprint("Size", size), func(b *testing.B) {
			slice := make([]int, size)
			element := 0
			step := 2

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = sliceutil.Interleave(slice, element, step)
			}
		})
	}
}

func BenchmarkInterleaveStrings(b *testing.B) {
	for _, size := range []int{10, 100, 1000} {
		b.Run(fmt.Sprint("Size", size), func(b *testing.B) {
			slice := make([]string, size)
			element := "element"
			step := 3

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = sliceutil.Interleave(slice, element, step)
			}
		})
	}
}

func TestDeleteFuncs(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		funcs []func(int) bool
		want  []int
	}{
		{
			name:  "Delete even numbers",
			slice: []int{1, 2, 3, 4, 5},
			funcs: []func(int) bool{
				func(n int) bool { return n%2 == 0 },
			},
			want: []int{1, 3, 5},
		},
		{
			name:  "Delete specific number",
			slice: []int{1, 2, 3, 4, 5},
			funcs: []func(int) bool{
				func(n int) bool { return n == 3 },
			},
			want: []int{1, 2, 4, 5},
		},
		{
			name:  "Multiple funcs",
			slice: []int{1, 2, 3, 4, 5, 6},
			funcs: []func(int) bool{
				func(n int) bool { return n%2 == 0 },
				func(n int) bool { return n == 5 },
			},
			want: []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceutil.DeleteFuncs(tt.slice, tt.funcs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFuncs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDeleteFuncs(b *testing.B) {
	slice := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
	}
	funcs := []func(int) bool{
		func(n int) bool { return n%2 == 0 },
		func(n int) bool { return n%3 == 0 },
	}

	b.ResetTimer()
	b.Run("v1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = sliceutil.DeleteFuncs(append([]int(nil), slice...), funcs...)
		}
	})
}
