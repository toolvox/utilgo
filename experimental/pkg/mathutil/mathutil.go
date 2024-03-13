package mathutil

import (
	"math"
	"sort"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Sign[N Number](n N) N {
	switch {
	case n > 0:
		return n / n
	case n < 0:
		return 0 - (n / n)
	default:
		return 0
	}
}

func Abs[N Number](n N) N {
	if n >= 0 {
		return n
	}
	return -n
}

func Max[N Number](vs ...N) N {
	if len(vs) == 0 {
		return 0
	}
	max := vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] > max {
			max = vs[i]
		}
	}
	return max
}

func Min[N Number](vs ...N) N {
	if len(vs) == 0 {
		return 0
	}
	min := vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] < min {
			min = vs[i]
		}
	}
	return min
}

func Bounds[N Number](vs ...N) (N, N) {
	if len(vs) == 0 {
		return 0, 0
	}
	min := vs[0]
	max := vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] > max {
			max = vs[i]
		}
		if vs[i] < min {
			min = vs[i]
		}
	}
	return min, max
}

func Clamp[N Number](min, val, max N) N {
	switch {
	case val < min:
		val = min
	case val > max:
		val = max
	}
	return val
}

func Sum[N Number](vs ...N) N {
	var sum N
	for _, v := range vs {
		sum += v
	}
	return sum
}

func Average[N Number](vs ...N) N {
	var sum N
	for _, v := range vs {
		sum += v
	}
	return sum / N(len(vs))
}

func Prod[N Number](vs ...N) N {
	var prod N = vs[0]
	for _, v := range vs[1:] {
		prod *= v
	}
	return prod
}

func GeometricMean[N Number](vs ...N) N {
	var prod float64 = 1.0
	for _, v := range vs {
		prod *= float64(v)
	}
	return N(math.Round(math.Pow(prod, 1.0/float64(len(vs)))))
}

func HarmonicMean[N Number](vs ...N) N {
	var sum float64 = 0
	for _, v := range vs {
		if v == 0 {
			return 0
		}
		sum += 1 / float64(v)
	}
	if sum == 0 {
		return 0
	}
	return N(math.Round(float64(len(vs)) / sum))
}

func XenoSum[N Number](vs ...N) N {
	sort.Slice(vs, func(i, j int) bool {
		return vs[i] < vs[j]
	})
	var sum, factor N
	sum, factor = 0, 1
	for i := len(vs) - 1; i >= 0; i-- {
		sum += vs[i] / factor
		factor *= 2
	}
	return sum
}
