package mathutil

import (
	"fmt"
)

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
