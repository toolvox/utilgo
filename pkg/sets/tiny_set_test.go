package sets_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"utilgo/pkg/sets"
)

func TestTinySet(t *testing.T) {
	tests := []struct {
		name          string
		operation     string
		set1          sets.TinySet[int]
		set2          sets.TinySet[int]
		addElements   []int
		expected      interface{}
		expectedThree [3]sets.TinySet[int]
	}{
		{
			name:      "UnionWith with Set",
			operation: "Union",
			set1:      sets.NewTinySet(1, 2),
			set2:      sets.NewTinySet(2, 3),
			expected:  sets.NewTinySet(1, 2, 3),
		},
		{
			name:        "UnionWith with Elements",
			operation:   "UnionWith",
			set1:        sets.NewTinySet(1, 2),
			addElements: []int{3, 2},
			expected:    sets.NewTinySet(1, 2, 3),
		},
		{
			name:      "IntersectionWith with Set",
			operation: "Intersection",
			set1:      sets.NewTinySet(1, 2),
			set2:      sets.NewTinySet(2, 3),
			expected:  sets.NewTinySet(2),
		},
		{
			name:        "IntersectionWith with Elements",
			operation:   "IntersectionWith",
			set1:        sets.NewTinySet(1, 2),
			addElements: []int{2, 3},
			expected:    sets.NewTinySet(2),
		},
		{
			name:      "DifferenceWith with Set",
			operation: "Difference",
			set1:      sets.NewTinySet(1, 2),
			set2:      sets.NewTinySet(2, 3),
			expected:  sets.NewTinySet(1),
		},
		{
			name:        "DifferenceWith with Elements",
			operation:   "DifferenceWith",
			set1:        sets.NewTinySet(1, 2),
			addElements: []int{2, 3},
			expected:    sets.NewTinySet(1),
		},
		{
			name:      "ThreeWay comparison",
			operation: "ThreeWay",
			set1:      sets.NewTinySet(1, 2),
			set2:      sets.NewTinySet(2, 3),
			expectedThree: [3]sets.TinySet[int]{
				sets.NewTinySet(2), // Common
				sets.NewTinySet(1), // Only in set1
				sets.NewTinySet(3), // Only in set2
			},
		},
		{
			name:        "Contains",
			operation:   "Contains",
			set1:        sets.NewTinySet(1, 2, 3),
			addElements: []int{2},
			expected:    true,
		},
		{
			name:      "String representation",
			operation: "String",
			set1:      sets.NewTinySet(1, 2, 3),
			expected:  "{1, 2, 3}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var result interface{}

			switch tc.operation {
			case "Union":
				result = tc.set1.Union(tc.set2)
			case "UnionWith":
				result = tc.set1.UnionWith(tc.addElements...)
			case "Intersection":
				result = tc.set1.Intersection(tc.set2)
			case "IntersectionWith":
				result = tc.set1.IntersectionWith(tc.addElements...)
			case "Difference":
				result = tc.set1.Difference(tc.set2)
			case "DifferenceWith":
				result = tc.set1.DifferenceWith(tc.addElements...)
			case "ThreeWay":
				result = tc.set1.ThreeWay(tc.set2)
				if !reflect.DeepEqual(result, tc.expectedThree) {
					t.Errorf("Expected %v, got %v", tc.expectedThree, result)
				}
				return
			case "Contains":
				result = tc.set1.Contains(tc.addElements[0])
			case "String":
				result = tc.set1.String()
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func strElement(i int) string {
	return fmt.Sprint(i)
}

func strElements(from, to int) []string {
	var res []string
	for i := from; i < to; i++ {
		res = append(res, strElement(i))
	}
	return res
}


func Benchmark_TinySet_New(b *testing.B) {
	for N := 5; N < 100; N += 25 {
		b.Run(fmt.Sprintf("Specific_%03d", N), func(b *testing.B) {
			elements1 := strElements(0, N)
			log.Println()
			b.Run("New0", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sets.NewTinySet(elements1...)
				}
			})

			b.Run("New1", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sets.NewTinySet(elements1...)
				}
			})

			b.Run("New2", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sets.NewTinySet(elements1...)
				}
			})

			log.Println()
		})
	}
}

func Benchmark_TinySet_Contains(b *testing.B) {
	for N := 5; N < 100; N += 25 {
		b.Run(fmt.Sprintf("Specific_%03d", N), func(b *testing.B) {
			elements1 := strElements(0, N)
			elements2 := strElements(N/2, 2*N)

			log.Println()
			b.Run("Contains0", func(b *testing.B) {
				tSet := sets.NewTinySet(elements1...)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = tSet.Contains(elements1[1])
					for _, v := range elements2 {
						_ = tSet.Contains(v)
					}
				}
			})

			b.Run("Contains1", func(b *testing.B) {
				tSet := sets.NewTinySet(elements1...)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = tSet.Contains(elements1[1])
					for _, v := range elements2 {
						_ = tSet.Contains(v)
					}
				}
			})

			log.Println()
		})
	}
}

func Benchmark_TinySet_Add(b *testing.B) {
	for N := 5; N < 100; N += 25 {
		b.Run(fmt.Sprintf("Specific_%03d", N), func(b *testing.B) {
			elements1 := strElements(0, N)
			elements2 := strElements(N/2, 2*N)

			b.Run("Add0/Batch", func(b *testing.B) {
				tSet := sets.NewTinySet(elements1...)
				for i := 0; i < b.N; i++ {
					tSet.Add(elements2...)
				}
			})

			b.Run("Add1/Batch", func(b *testing.B) {
				tSet := sets.NewTinySet(elements1...)
				for i := 0; i < b.N; i++ {
					tSet.Add(elements2...)
				}
			})

			log.Println()
		})
	}
}

func Benchmark_TinySet_Add_Single(b *testing.B) {
	for N := 5; N < 100; N += 25 {
		b.Run(fmt.Sprintf("Specific_%03d", N), func(b *testing.B) {
			elements1 := strElements(0, N)
			elements3 := strElements(N, 4*N)

			log.Println()

			b.Run("Add0/Single", func(b *testing.B) {
				tSet := sets.NewTinySet(elements1...)
				for i := 0; i < b.N; i++ {
					for j := 0; j < len(elements3); j++ {
						tSet.Add(elements3[j])
					}
				}
			})

			b.Run("Add1/Single", func(b *testing.B) {
				tSet := sets.NewTinySet(elements1...)
				for i := 0; i < b.N; i++ {
					for j := 0; j < len(elements3); j++ {
						tSet.Add(elements3[j])
					}
				}
			})

			log.Println()
		})
	}
}
