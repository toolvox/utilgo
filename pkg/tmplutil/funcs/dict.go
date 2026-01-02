package funcs

import (
	"fmt"
	"reflect"
)

func DictFunc(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("dict expects an even number of arguments")
	}

	result := make(map[string]interface{})
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}

		value := values[i+1]
		valueRef := reflect.ValueOf(value)
		if !valueRef.IsValid() {
			if !canBeNil(valueRef.Type()) {
				return nil, fmt.Errorf("value is nil; should be of type %s", valueRef.Type())
			}
			valueRef = reflect.Zero(valueRef.Type())
		}

		result[key] = valueRef.Interface()
	}

	return result, nil
}

func ContainsKeyFunc(data interface{}, key interface{}) (bool, error) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Map {
		return false, fmt.Errorf("input is not a map")
	}

	for _, k := range v.MapKeys() {
		if reflect.DeepEqual(k.Interface(), key) {
			return true, nil
		}
	}
	return false, nil
}
