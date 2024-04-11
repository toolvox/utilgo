package mathutil

import (
	"fmt"
	"math"
	"sort"
)

// Max returns the maximum value among the provided values.
func Max[N Number](values ...N) N {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

// Min returns the minimum value among the provided values.
func Min[N Number](values ...N) N {
	if len(values) == 0 {
		return 0
	}
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}

// Sum calculates the sum of the provided values.
func Sum[N Number](values ...N) N {
	var sum N
	for _, v := range values {
		sum += v
	}
	return sum
}

// Average calculates the average of the provided values.
func Average[N Number](values ...N) N {
	if len(values) == 0 {
		return 0
	}
	return Sum(values...) / N(len(values))
}

// Product calculates the product of the provided values.
func Product[N Number](values ...N) N {
	if len(values) == 0 {
		return 0
	}
	var prod N = values[0]
	for _, v := range values[1:] {
		prod *= v
	}
	return prod
}

// GeometricMean calculates the geometric mean of the provided values.
func GeometricMean[N Number](values ...N) N {
	var prod float64 = 1.0
	for _, v := range values {
		prod *= float64(v)
	}
	resF := math.Pow(prod, 1.0/float64(len(values)))
	resRound := N(math.Round(resF))
	if math.Abs(float64(resRound-N(resF))) >= 1 {
		return N(resRound)
	}
	return N(resF)
}

// HarmonicMean calculates the harmonic mean of the provided values.
func HarmonicMean[N Number](values ...N) N {
	var sum float64 = 0
	for _, v := range values {
		if v == 0 {
			return 0
		}
		sum += 1 / float64(v)
	}
	if sum == 0 {
		return 0
	}
	resF := float64(len(values)) / sum
	resRound := N(math.Round(resF))
	if math.Abs(float64(resRound-N(resF))) >= 1 {
		return N(resRound)
	}
	return N(resF)
}

// XenoSum calculates a weighted sum where each subsequent (sorted largest to smallest) value is half as significant as the previous one.
func XenoSum[N Number](values ...N) N {
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	var sum, factor float64
	sum, factor = 0, 1
	for _, v := range values {
		sum += float64(v) / factor
		factor *= 2
	}
	resRound := N(math.Round(sum))
	if math.Abs(float64(resRound-N(sum))) >= 1 {
		return N(resRound)
	}
	return N(sum)
}

// Aggregator defines an interface for aggregating values.
type Aggregator[N Number] interface {
	Aggregate(values ...N) N
}

// AggregatorFunc is a function type that implements Aggregator.
type AggregatorFunc[N Number] func(values ...N) N

// Aggregate calls the AggregatorFunc with the provided values.
func (af AggregatorFunc[N]) Aggregate(values ...N) N {
	return af(values...)
}

// AggregatorKey is a string key representing an aggregator function.
type AggregatorKey string

// Constants for AggregatorKey representing different aggregator functions.
const (
	AggMax           AggregatorKey = "Max"
	AggMin           AggregatorKey = "Min"
	AggSum           AggregatorKey = "Sum"
	AggAverage       AggregatorKey = "Avg"
	AggProduct       AggregatorKey = "Prod"
	AggGeometricMean AggregatorKey = "Geom"
	AggHarmonicMean  AggregatorKey = "Harm"
	AggXenoSum       AggregatorKey = "Xeno"
)

// GetAggregator returns the AggregatorFunc corresponding to the given AggregatorKey.
func GetAggregator[N Number](key AggregatorKey) AggregatorFunc[N] {
	switch key {
	case AggMax:
		return Max[N]
	case AggMin:
		return Min[N]
	case AggSum:
		return Sum[N]
	case AggAverage:
		return Average[N]
	case AggProduct:
		return Product[N]
	case AggGeometricMean:
		return GeometricMean[N]
	case AggHarmonicMean:
		return HarmonicMean[N]
	case AggXenoSum:
		return XenoSum[N]
	default:
		panic(fmt.Errorf("unknown aggregator: %s", key))
	}
}
