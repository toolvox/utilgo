package mathutil

import (
	"math"
	"unsafe"

	"github.com/toolvox/utilgo/api"
)

type Number = api.Number

func MaxValue[N Number]() N {
	var val N
	switch any(val).(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr:
		var maxUintVal uint64 = ^uint64(0)
		return *(*N)(unsafe.Pointer(&maxUintVal))

	case int8, int16, int32, int64, int:
		var maxIntVal int64 = math.MaxInt8
		size := sizeof[N]()
		for i := 1; i < size; i++ {
			maxIntVal <<= 8
			maxIntVal += 0xFF
		}
		return *(*N)(unsafe.Pointer(&maxIntVal))

	case float32:
		var maxFloat32Val float32 = math.MaxFloat32
		return *(*N)(unsafe.Pointer(&maxFloat32Val))

	case float64:
		var maxFloat64Val float64 = math.MaxFloat64
		return *(*N)(unsafe.Pointer(&maxFloat64Val))
	default:
		panic("unsupported type")
	}
}

func MinValue[N Number]() N {
	var val N
	switch any(val).(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr:
		var minUintVal uint64 = uint64(0)
		return *(*N)(unsafe.Pointer(&minUintVal))

	case int8, int16, int32, int64, int:
		var minIntVal int64 = math.MinInt8
		size := sizeof[N]()
		for i := 1; i < size; i++ {
			minIntVal <<= 8
		}
		return *(*N)(unsafe.Pointer(&minIntVal))

	case float32:
		var minFloat32Val float32 = -math.MaxFloat32
		return *(*N)(unsafe.Pointer(&minFloat32Val))

	case float64:
		var minFloat64Val float64 = -math.MaxFloat64
		return *(*N)(unsafe.Pointer(&minFloat64Val))

	default:
		panic("unsupported type")
	}
}

func EpsilonValue[N Number]() N {
	var val N
	switch any(val).(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr,
		int8, int16, int32, int64, int:
		var epsilonUintVal uint64 = 1
		return *(*N)(unsafe.Pointer(&epsilonUintVal))
	case float32:
		var epsilonFloat32Val float32 = math.SmallestNonzeroFloat32
		return *(*N)(unsafe.Pointer(&epsilonFloat32Val))

	case float64:
		var epsilonFloat64Val float64 = math.SmallestNonzeroFloat64
		return *(*N)(unsafe.Pointer(&epsilonFloat64Val))

	default:
		panic("unsupported type")
	}
}

func One[N Number]() N {
	return N(1)
}

func sizeof[T any]() int {
	var val T
	return int(unsafe.Sizeof(val))
}
