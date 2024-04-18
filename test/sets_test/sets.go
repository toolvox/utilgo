package sets_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/api/sets"
	"github.com/toolvox/utilgo/test"
)

func Run_Test_Set[S sets.BasicSet[C], C comparable](t *testing.T, ctor test.TestConstructorFor[S], data test.TypedTestData[C], maxN int) {
	t.Run("BasicSet", func(t *testing.T) {
		t.Run("Len|Add", func(t *testing.T) {
			must := require.New(t)
			testSet := ctor.New()
			must.NotNil(testSet)
			must.Equal(0, testSet.Len())
			for n := 1; n <= maxN; n *= 10 {
				t.Run(fmt.Sprint("n=", n), func(t *testing.T) {
					for i := 0; i < n; i++ {
						testSet.Add(data.Element(i))
					}
					must.Equal(n, testSet.Len())
				})
			}
		})
		t.Run("Contains", func(t *testing.T) {
			must := require.New(t)
			testSet := ctor.New()
			must.NotNil(testSet)
			must.Equal(0, testSet.Len())
			for n := 1; n <= maxN; n *= 10 {
				t.Run(fmt.Sprint("n=", n), func(t *testing.T) {
					for i := 0; i < n; i++ {
						testSet.Add(data.Element(i))
					}
					for i := -n; i < 2*n; i += max(1, n/10) {
						must.Equal(i >= 0 && i < n, testSet.Contains(data.Element(i)), "i=%d, n=%d", i, n)
					}
				})
			}
		})
		t.Run("Elements", func(t *testing.T) {
			must := require.New(t)
			testSet := ctor.New()
			must.NotNil(testSet)
			must.Empty(testSet.Elements())
			var testElements []C = make([]C, 0, maxN)
			for v := range maxN {
				testElements = append(testElements, data.Element(v))
			}
			for n := 1; n <= maxN; n *= 10 {
				t.Run(fmt.Sprint("n=", n), func(t *testing.T) {
					for i := 0; i < n; i++ {
						testSet.Add(data.Element(i))
					}
					must.ElementsMatch(testElements[:n], testSet.Elements())
				})
			}
		})
		t.Run("String", func(t *testing.T) {
			must := require.New(t)
			testSet := ctor.New()
			must.NotNil(testSet)
			must.Equal("{ }", testSet.String())
			for n := 1; n <= maxN; n *= 10 {
				t.Run(fmt.Sprint("n=", n), func(t *testing.T) {
					expectLines := []string{}
					for i := 0; i < n; i++ {
						testSet.Add(data.Element(i))
						expectLines = append(expectLines, fmt.Sprint(data.Element(i)))
					}
					actualString := testSet.String()
					for _, line := range expectLines {
						must.Contains(actualString, line)
					}
					must.Equal(len(expectLines)-1, strings.Count(actualString, ", "))
				})
			}
		})
		t.Run("Remove", func(t *testing.T) {
			must := require.New(t)
			testSet := ctor.New()
			must.NotNil(testSet)
			must.Equal(0, testSet.Len())
			for n := 1; n <= maxN; n *= 10 {
				t.Run(fmt.Sprint("n=", n), func(t *testing.T) {
					for i := 0; i < n; i++ {
						testSet.Add(data.Element(i))
					}
					for i := -n; i < 2*n; i += max(1, n/10) {
						must.Equal(i >= 0 && i < n, testSet.Contains(data.Element(i)), "i=%d, n=%d", i, n)
						testSet.Remove(data.Element(i))
						must.Equal(false, testSet.Contains(data.Element(i)), "i=%d, n=%d", i, n)
					}
				})
			}
		})
	})
}
