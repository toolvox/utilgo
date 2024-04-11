package maputil

import "encoding/json"

// MapToStruct converts a map to a struct by converting through json.
func MapToStruct[T any](m map[string]any) (T, error) {
	var result T
	data, err := json.Marshal(m)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// StructToMap converts a struct to a map by converting through json.
func StructToMap[T any](t T) (map[string]any, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	var out map[string]any
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
