// Package reflectutil provides reflection related helpers and helpers that heavily rely on reflections.
package reflectutil

import (
	"reflect"
)

// IsZero reflects a variables to determine if it is equivalent to a non-value.
//
// Empty lists and strings are considered zero. Default values of primitives are also considered zero.
//
// Warning! This function has been partially tested, use at your own risk :O
func IsZero[T any](v T) bool {
	var rv reflect.Value
	var ok bool
	if rv, ok = any(v).(reflect.Value); !ok {
		rv = reflect.ValueOf(v)
	}
	switch rk := rv.Kind(); rk {
	case reflect.Invalid:
		return true

	case reflect.Bool:
		return !rv.Bool()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.String,
		reflect.Complex64, reflect.Complex128:
		return rv.IsZero()

	case reflect.Array:
		if rv.Len() == 0 || rv.IsZero() {
			return true
		}
		for i := 0; i < rv.Len(); i++ {
			element := rv.Index(i)
			if !IsZero(element) {
				return false
			}
		}
		return true

	case reflect.Chan:
		if rv.Len() == 0 || rv.IsZero() || rv.IsNil() {
			return true
		}

	case reflect.Func:
		if rv.IsZero() || rv.IsNil() {
			return true
		}

	case reflect.Map:
		if rv.Len() == 0 || rv.IsZero() || rv.IsNil() {
			return true
		}

		keys := rv.MapKeys()
		for i := 0; i < rv.Len(); i++ {
			key := keys[i]
			val := rv.MapIndex(key)
			if !IsZero(key) && !IsZero(val) {
				return false
			}
		}
		return true

	case reflect.Pointer:
		return !rv.Elem().IsValid() || IsZero(rv.Elem())

	case reflect.Slice:
		if rv.Len() == 0 || rv.IsZero() || rv.IsNil() {
			return true
		}
		for i := 0; i < rv.Len(); i++ {
			element := rv.Index(i)
			if !IsZero(element) {
				return false
			}
		}
		return true

	case reflect.Struct:
		return rv.IsZero()

	case reflect.Interface:
		return rv.IsNil() || rv.IsZero()

	case reflect.UnsafePointer:
		panic("case reflect.UnsafePointer")

	default:
		panic("unknown")

	}
	return false
}
