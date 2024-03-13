package tmplz

import (
	"fmt"
	"maps"
	"reflect"
)

// Assignments holds key-value pairs for template variable assignments.
type Assignments map[string]any

// Get retrieves the value associated with the given key as a [String], and a boolean indicating if the key was found.
func (a Assignments) Get(key string) (String, bool) {
	if res, ok := a[key]; ok {
		return String(fmt.Sprint(res)), ok
	}
	return String{}, false
}

// Assign sets a value for a key in [Assignments]. If the key already exists, its value is updated.
func (a Assignments) Assign(key string, value any) {
	a[key] = value
}

// AssignMap assigns multiple key-value pairs from a given map to the [Assignments].
func (a Assignments) AssignMap(kvps map[string]any) {
	for k, v := range kvps {
		a[k] = v
	}
}

// AssignObj uses the exported fields of the provided object to assign key-value pairs to the Assignments.
// Each field name is used as a key.
func (a Assignments) AssignObj(obj any) {
	dict := utilStructToMap(obj)
	a.AssignMap(dict)
}

// Clone creates a deep copy of the Assignments and returns it.
func (a Assignments) Clone() Assignments {
	return maps.Clone(a)
}

// utilStructToMap is a helper function that converts the exported fields of a struct to a map[string]any.
// Each field name becomes a key in the map.
func utilStructToMap(obj any) map[string]any {
	// TODO: this function wants to be in maputil
	result := make(map[string]any)
	objValue := reflect.ValueOf(obj)

	if objValue.Kind() == reflect.Struct {
		objType := objValue.Type()
		for i := 0; i < objValue.NumField(); i++ {
			field := objValue.Field(i)
			result[objType.Field(i).Name] = field.Interface()
		}
	}
	return result
}
