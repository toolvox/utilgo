package sets_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/toolvox/utilgo/pkg/sets"
	"github.com/toolvox/utilgo/test"
	test_sets "github.com/toolvox/utilgo/test/sets_test"
)

// func Benchmark_Contains(b *testing.B) {
// 	for n := 1; n <= 10_000_000; n *= 10 {
// 		b.Run(fmt.Sprintf("%8d", n), func(b *testing.B) {
// 			testSet := sets.NewSet(0)
// 			items := []int{}
// 			for i := range n {
// 				testSet.Add(i + 1)
// 				items = append(items, i+1)
// 			}
// 			b.Run("one_item", func(b *testing.B) {
// 				b.Run("1", func(b *testing.B) {
// 					for i := 0; i < b.N; i++ {
// 						testSet.Contains2(i)
// 					}
// 				})
// 				b.Run("2", func(b *testing.B) {
// 					for i := 0; i < b.N; i++ {
// 						testSet.Contains(i)
// 					}
// 				})
// 			})
// 			b.Run("5_item", func(b *testing.B) {
// 				b.Run("1", func(b *testing.B) {
// 					for i := 0; i < b.N; i++ {
// 						testSet.Contains2(i-2, i-1, i, i+1, i+2)
// 					}
// 				})
// 				b.Run("2", func(b *testing.B) {
// 					for i := 0; i < b.N; i++ {
// 						testSet.Contains(i-2, i-1, i, i+1, i+2)
// 					}
// 				})
// 			})
// 			b.Run("n/2_item", func(b *testing.B) {
// 				b.Run("1", func(b *testing.B) {
// 					for i := 0; i < b.N; i++ {
// 						testSet.Contains2(items[n/2:]...)
// 					}
// 				})
// 				b.Run("2", func(b *testing.B) {
// 					for i := 0; i < b.N; i++ {
// 						testSet.Contains(items[n/2:]...)
// 					}
// 				})
// 			})
// 		})
// 	}
// }

func Test_Set(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		t.Run("new", func(t *testing.T) {
			testSet := sets.NewSet(0, 1, 2, 3, 1, 3, 5)
			require.True(t, testSet.Contains(1, 2, 3, 5))
		})
		test_sets.Run_Test_Set(t,
			test.TestConstructorFor[sets.Set[int]]{
				NewFunc: func() sets.Set[int] {
					return sets.NewSet[int]()
				},
			},
			test.TestDataFor[int]{
				ElementFunc: func(i int) int { return i },
			},
			10_000,
		)
	})
	t.Run("string", func(t *testing.T) {
		t.Run("new", func(t *testing.T) {
			testSet := sets.NewSet("0", "1", "2", "3", "1", "3", "5")
			require.True(t, testSet.Contains("1", "2", "3", "5"))
		})
		test_sets.Run_Test_Set(t,
			test.TestConstructorFor[sets.Set[string]]{
				NewFunc: func() sets.Set[string] {
					return sets.NewSet[string]()
				},
			},
			test.TestDataFor[string]{
				ElementFunc: func(i int) string { return fmt.Sprint(i) },
			},
			10_000,
		)
	})
}

func TestUnion(t *testing.T) {
	s1 := sets.NewSet(1, 2)
	s2 := sets.NewSet(2, 3)
	UnionWith := s1.Union(s2)
	if len(UnionWith) != 3 || !UnionWith.Contains(1) || !UnionWith.Contains(2) || !UnionWith.Contains(3) {
		t.Errorf("UnionWith does not contain the correct elements")
	}
}

func TestUnionWith(t *testing.T) {
	s := sets.NewSet(1, 2)
	UnionWith := s.UnionWith(3, 4)
	if len(UnionWith) != 4 || !UnionWith.Contains(1) || !UnionWith.Contains(4) {
		t.Errorf("UnionWith does not contain the correct elements")
	}
}

func TestIntersection(t *testing.T) {
	s1 := sets.NewSet(1, 2, 3)
	s2 := sets.NewSet(2, 3, 4)
	IntersectionWith := s1.Intersection(s2)
	if len(IntersectionWith) != 2 || !IntersectionWith.Contains(2) || !IntersectionWith.Contains(3) {
		t.Errorf("IntersectionWith does not contain the correct elements")
	}
}

func TestIntersectionWith(t *testing.T) {
	s := sets.NewSet(1, 2, 3)
	IntersectionWith := s.IntersectionWith(2, 3, 4)
	if len(IntersectionWith) != 2 || !IntersectionWith.Contains(2) || !IntersectionWith.Contains(3) {
		t.Errorf("IntersectionWith does not contain the correct elements")
	}
}

func TestDifference(t *testing.T) {
	s1 := sets.NewSet(1, 2, 3)
	s2 := sets.NewSet(2, 3, 4)
	DifferenceWith := s1.Difference(s2)
	if len(DifferenceWith) != 1 || !DifferenceWith.Contains(1) {
		t.Errorf("DifferenceWith does not contain the correct elements")
	}
}

func TestDifferenceWith(t *testing.T) {
	s := sets.NewSet(1, 2, 3)
	DifferenceWith := s.DifferenceWith(2, 3)
	if len(DifferenceWith) != 1 || !DifferenceWith.Contains(1) {
		t.Errorf("DifferenceWith does not contain the correct elements")
	}
}

func TestThreeWay(t *testing.T) {
	s1 := sets.NewSet(1, 2)
	s2 := sets.NewSet(2, 3)
	threeWay := s1.ThreeWay(s2)
	if len(threeWay[0]) != 1 || len(threeWay[1]) != 1 || len(threeWay[2]) != 1 {
		t.Errorf("ThreeWay does not contain the correct elements")
	}
}

func Test_Set_Encoding(t *testing.T) {
	testSet := sets.NewSet("hello", "goodbye", "salute")
	t.Run("YAML", func(t *testing.T) {
		var bs []byte
		var err error
		t.Run("Marshal", func(t *testing.T) {
			bs, err = yaml.Marshal(testSet)
			require.NoError(t, err)
			for v := range testSet {
				require.Contains(t, string(bs), v)
			}
		})
		t.Run("Unmarshal", func(t *testing.T) {
			var cycleSet sets.Set[string]
			err = yaml.Unmarshal(bs, &cycleSet)
			require.NoError(t, err)
			require.Equal(t, testSet, cycleSet)
		})
	})

	t.Run("JSON", func(t *testing.T) {
		var bs []byte
		var err error
		t.Run("Marshal", func(t *testing.T) {
			bs, err = json.Marshal(testSet)
			require.NoError(t, err)
			for v := range testSet {
				require.Contains(t, string(bs), v)
			}
		})
		t.Run("Unmarshal", func(t *testing.T) {
			var cycleSet sets.Set[string]
			err = json.Unmarshal(bs, &cycleSet)
			require.NoError(t, err)
			require.Equal(t, testSet, cycleSet)
		})
	})

}
