package mathutil_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/toolvox/utilgo/pkg/mathutil"
)

func Run_Test_NumberType_Funcs[N mathutil.Number](max, min, epsilon, one N) func(*testing.T) {
	return func(t *testing.T) {
		require.Equal(t, max, mathutil.MaxValue[N]())
		require.Equal(t, min, mathutil.MinValue[N]())
		require.Equal(t, epsilon, mathutil.EpsilonValue[N]())
		require.Equal(t, one, mathutil.One[N]())
	}
}

// TestNumberFunctions runs a unified table-driven test for all number functions.
func Test_NumberType_Funcs(t *testing.T) {

	t.Run("int", Run_Test_NumberType_Funcs(int(math.MaxInt), int(math.MinInt), int(1), int(1)))
	t.Run("int8", Run_Test_NumberType_Funcs(int8(math.MaxInt8), int8(math.MinInt8), int8(1), int8(1)))
	t.Run("int16", Run_Test_NumberType_Funcs(int16(math.MaxInt16), int16(math.MinInt16), int16(1), int16(1)))
	t.Run("int32", Run_Test_NumberType_Funcs(int32(math.MaxInt32), int32(math.MinInt32), int32(1), int32(1)))
	t.Run("int64", Run_Test_NumberType_Funcs(int64(math.MaxInt64), int64(math.MinInt64), int64(1), int64(1)))
	t.Run("uint", Run_Test_NumberType_Funcs(uint(math.MaxUint), uint(0), uint(1), uint(1)))
	t.Run("uint8", Run_Test_NumberType_Funcs(uint8(math.MaxUint8), uint8(0), uint8(1), uint8(1)))
	t.Run("uint16", Run_Test_NumberType_Funcs(uint16(math.MaxUint16), uint16(0), uint16(1), uint16(1)))
	t.Run("uint32", Run_Test_NumberType_Funcs(uint32(math.MaxUint32), uint32(0), uint32(1), uint32(1)))
	t.Run("uint64", Run_Test_NumberType_Funcs(uint64(math.MaxUint64), uint64(0), uint64(1), uint64(1)))
	t.Run("float32", Run_Test_NumberType_Funcs(float32(math.MaxFloat32), float32(-math.MaxFloat32), float32(math.SmallestNonzeroFloat32), float32(1.0)))
	t.Run("float64", Run_Test_NumberType_Funcs(float64(math.MaxFloat64), float64(-math.MaxFloat64), float64(math.SmallestNonzeroFloat64), float64(1.0)))

}
