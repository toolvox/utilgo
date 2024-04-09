package sets_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/toolvox/utilgo/pkg/sets"
)

func Benchmark_Add(b *testing.B) {
	for n := 0; n < 100; n++ {

		b.Run(fmt.Sprint(n), func(b *testing.B) {
			var setTook time.Duration
			var tinySetTook time.Duration
			b.Run("Set", func(b *testing.B) {
				testSet := sets.NewSet[int]()
				for i := 0; i < b.N; i++ {
					for v := range n {
						testSet.Add(v)
					}
				}
				b.StopTimer()
				setTook = b.Elapsed() / time.Duration(b.N)
			})
			b.Run("Set2", func(b *testing.B) {
				testSet := sets.NewSet[int]()
				for i := 0; i < b.N; i++ {
					for v := range n {
						testSet.Add(v)
					}
				}
			})
			b.Run("TinySet", func(b *testing.B) {
				testSet := sets.NewTinySet[int]()
				for i := 0; i < b.N; i++ {
					for v := range n {
						testSet.Add(v)
					}
					for v := range n / 2 {
						testSet.Add(v * 2)
					}
				}
				b.StopTimer()
				tinySetTook = b.Elapsed() / time.Duration(b.N)
			})
			b.Run("TinySet2", func(b *testing.B) {
				testSet := sets.NewTinySet[int]()
				for i := 0; i < b.N; i++ {
					for v := range n {
						testSet.Add(v)
					}
					for v := range n / 2 {
						testSet.Add(v * 2)
					}
				}
			})
			log.Println(setTook - tinySetTook)
		})
	}
}
