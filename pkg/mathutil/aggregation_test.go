package mathutil_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/api"
	"github.com/toolvox/utilgo/pkg/mathutil"
)

type testCase[N api.Number] struct {
	name   string
	values []N
	want   N
}

func runTestCases[N api.Number](t *testing.T, tests []testCase[N], aggKey mathutil.AggregatorKey) {
	t.Run(string(aggKey), func(t *testing.T) {
		agg := mathutil.GetAggregator[N](aggKey)
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := agg.Aggregate(tt.values...)
				require.InDelta(t, tt.want, got, 0.000000001)
			})
		}
	})
}

func TestMaxInt(t *testing.T) {
	tests := []testCase[int]{
		{"single element", []int{1}, 1},
		{"positive numbers", []int{1, 2, 3}, 3},
		{"negative numbers", []int{-3, -2, -1}, -1},
		{"mixed numbers", []int{-1, 0, 1}, 1},
		{"with zero", []int{0, 2, 3}, 3},
		{"empty slice", []int{}, 0},
	}
	runTestCases(t, tests, mathutil.AggMax)
}

func TestMaxFloat64(t *testing.T) {
	tests := []testCase[float64]{
		{"single element", []float64{1.5}, 1.5},
		{"floating points", []float64{1.1, 2.2, 3.3}, 3.3},
		{"negative floats", []float64{-1.1, -2.2, -0.5}, -0.5},
		{"mixed floats", []float64{-1.5, 0.0, 1.5}, 1.5},
		{"with zero", []float64{0.0, 2.5, 3.5}, 3.5},
		{"empty slice", []float64{}, 0},
	}
	runTestCases(t, tests, mathutil.AggMax)
}

// Completing TestMinInt with additional test cases
func TestMinInt(t *testing.T) {
	tests := []testCase[int]{
		{"single element", []int{2}, 2},
		{"positive numbers", []int{4, 1, 2, 3}, 1},
		{"negative numbers", []int{-2, -3, -2, -1}, -3},
		{"mixed numbers", []int{-1, 0, 1}, -1},
		{"with zero", []int{0, 2, 3}, 0},
		{"empty slice", []int{}, 0},
	}
	runTestCases(t, tests, mathutil.AggMin)
}

// Implementing TestMinFloat64 function
func TestMinFloat64(t *testing.T) {
	tests := []testCase[float64]{
		{"single element", []float64{2.5}, 2.5},
		{"positive floats", []float64{1.1, 2.2, 3.3}, 1.1},
		{"negative floats", []float64{-3.3, -2.2, -1.1}, -3.3},
		{"mixed floats", []float64{-1.1, 0.0, 1.1}, -1.1},
		{"with zero", []float64{0.0, 1.1, 2.2}, 0.0},
		{"empty slice", []float64{}, 0},
	}
	runTestCases(t, tests, mathutil.AggMin)
}

// Tests for Sum function
func TestSumInt(t *testing.T) {
	tests := []testCase[int]{
		{"all positive", []int{1, 2, 3, 4}, 10},
		{"all negative", []int{-1, -2, -3, -4}, -10},
		{"mixed", []int{-2, 4, -6, 8}, 4},
		{"with zero", []int{0, 1, 2, 3}, 6},
		{"single element", []int{5}, 5},
		{"empty slice", []int{}, 0},
	}
	runTestCases(t, tests, mathutil.AggSum)
}

func TestSumFloat64(t *testing.T) {
	tests := []testCase[float64]{
		{"all positive", []float64{1.1, 2.2, 3.3}, 6.6},
		{"all negative", []float64{-1.1, -2.2, -3.3}, -6.6},
		{"mixed", []float64{-2.2, 4.4, -6.6, 8.8}, 4.4},
		{"with zero", []float64{0, 1.1, 2.2, 3.3}, 6.6},
		{"single element", []float64{5.5}, 5.5},
		{"empty slice", []float64{}, 0},
	}
	runTestCases(t, tests, mathutil.AggSum)
}

// Tests for Average function
func TestAverageInt(t *testing.T) {
	tests := []testCase[int]{
		{"all positive", []int{2, 4, 6, 8}, 5},
		{"all negative", []int{-2, -4, -6, -8}, -5},
		{"mixed", []int{-2, 4, -6, 8}, 1},
		{"with zero", []int{0, 2, 4, 6}, 3},
		{"single element", []int{7}, 7},
		{"empty slice", []int{}, 0},
	}
	runTestCases(t, tests, mathutil.AggAverage)
}

func TestAverageFloat64(t *testing.T) {
	tests := []testCase[float64]{
		{"all positive", []float64{1.5, 3.5, 5.5}, 3.5},
		{"all negative", []float64{-1.5, -3.5, -5.5}, -3.5},
		{"mixed", []float64{-2.5, 5.0, -7.5, 10.0}, 1.25},
		{"with zero", []float64{0.0, 2.0, 4.0, 6.0}, 3.0},
		{"single element", []float64{7.7}, 7.7},
		{"empty slice", []float64{}, 0},
	}
	runTestCases(t, tests, mathutil.AggAverage)
}

// Tests for Product function
func TestProductInt(t *testing.T) {
	tests := []testCase[int]{
		{"all positive", []int{2, 3, 4}, 24},
		{"all negative", []int{-2, -3, -4}, -24},
		{"mixed", []int{-2, 3, -4, 5}, 120},
		{"with zero", []int{0, 2, 3, 4}, 0},
		{"single element", []int{6}, 6},
		{"empty slice", []int{}, 0},
	}
	runTestCases(t, tests, mathutil.AggProduct)
}

func TestProductFloat64(t *testing.T) {
	tests := []testCase[float64]{
		{"all positive", []float64{1.5, 2.5, 3.5}, 13.125},
		{"all negative", []float64{-1.5, -2.5, -3.5}, -13.125},
		{"mixed", []float64{-1.5, 2.5, -3.5, 4.5}, 59.0625},
		{"with zero", []float64{0.0, 1.5, 2.5, 3.5}, 0.0},
		{"single element", []float64{5.5}, 5.5},
		{"empty slice", []float64{}, 0},
	}
	runTestCases(t, tests, mathutil.AggProduct)
}

func TestGeometricMeanInt(t *testing.T) {
	tests := []testCase[int]{
		{"empty slice", []int{}, 1},
		{"all positive", []int{1, 2, 3, 4}, 2},
		{"two numbers", []int{1, 256}, 16},
		{"single element", []int{100}, 100},
		{"identical numbers", []int{3, 3, 3}, 3},
		{"large numbers", []int{1000, 10000, 100000}, 10000},
	}
	runTestCases(t, tests, mathutil.AggGeometricMean)
}

func TestGeometricMeanFloat64(t *testing.T) {
	tests := []testCase[float64]{
		{"all positive", []float64{1.0, 2.0, 4.0, 8.0}, 2.8284271247461903},
		{"two numbers", []float64{0.01, 10_000.0}, 10.0},
		{"single element", []float64{64.0}, 64.0},
		{"identical numbers", []float64{2.5, 2.5, 2.5}, 2.5},
		{"large numbers", []float64{1000.0, 10000.0, 100000.0}, 10000.0},
		{"empty slice", []float64{}, 1.0},
	}
	runTestCases(t, tests, mathutil.AggGeometricMean)
}

func TestHarmonicMeanInt(t *testing.T) {
	tests := []testCase[int]{
		{"all positive", []int{1, 2, 4}, 2},
		{"two numbers", []int{1, 3}, 2},
		{"single element", []int{10}, 10},
		{"identical numbers", []int{5, 50, 5}, 7},
		{"zero in values", []int{0, 1, 2}, 0},
		{"empty slice", []int{}, 0},
	}
	runTestCases(t, tests, mathutil.AggHarmonicMean)
}

func TestHarmonicMeanFloat64(t *testing.T) {
	tests := []testCase[float64]{
		{"all positive", []float64{1.0, 2.0, 4.0}, 1.7142857142857142},
		{"two numbers", []float64{1.5, 4.5}, 2.25},
		{"single element", []float64{20.0}, 20.0},
		{"identical numbers", []float64{3.3, 3.3, 3.3}, 3.3},
		{"zero in values", []float64{0.0, 2.2, 3.3}, 0.0},
		{"empty slice", []float64{}, 0},
	}
	runTestCases(t, tests, mathutil.AggHarmonicMean)
}

func TestXenoSumInt(t *testing.T) {
	tests := []testCase[int]{
		{"all positive", []int{1, 2, 3, 4}, 6},
		{"two numbers", []int{3, 6}, 8},
		{"single element", []int{10}, 10},
		{"identical numbers", []int{2, 2, 2}, 4},
		{"empty slice", []int{}, 0},
	}
	runTestCases(t, tests, mathutil.AggXenoSum)
}

func TestXenoSumFloat64(t *testing.T) {
	tests := []testCase[float64]{
		{"all positive", []float64{1.5, 3.0, 6.0}, 7.875},
		{"two numbers", []float64{2.5, 5.0}, 6.25},
		{"single element", []float64{20.0}, 20.0},
		{"identical numbers", []float64{4.4, 4.4, 4.4}, 7.7},
		{"empty slice", []float64{}, 0},
	}
	runTestCases(t, tests, mathutil.AggXenoSum)
}
