package randutil_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/toolvox/utilgo/pkg/randutil"
)

func Test_Derange(t *testing.T) {
	for n := 2; n <= 1024*1024*16; n *= 2 {
		rng := rand.New(rand.NewSource(int64(n)))
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			res := randutil.Derange(rng, n)
			for i, v := range res {
				if v == i {
					t.Errorf("res[i]==i (%d)", i)
				}
			}
		})
	}
}

// Benchmark_Derange ASD
func Benchmark_Derange(b *testing.B) {
	for n := 2; n <= 1024*1024*16; n *= 2 {
		rng := rand.New(rand.NewSource(int64(n)))
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = randutil.Derange(rng, n)
			}
		})
	}
}
