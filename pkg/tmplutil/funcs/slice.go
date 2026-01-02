package funcs

import "reflect"

func SliceContains(s reflect.Value, k reflect.Value) bool {
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == k.Interface() {
			return true
		}
	}
	return false
}
