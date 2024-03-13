package yaml

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Marshal encodes the object to a yaml byte slice.
// Returns the encoded yaml byte slice or the error encountered.
func Marshal[T any](obj T) (data []byte, err error) {
	data, err = yaml.Marshal(obj)
	if err != nil {
		err = fmt.Errorf("marshal yaml: %w", err)
		return
	}
	return
}

// MustMarshal encodes the object to a yaml byte slice.
// Returns the encoded yaml byte slice or panics if an error is encountered.
func MustMarshal[T any](obj T) (data []byte) {
	var err error
	data, err = yaml.Marshal(obj)
	if err != nil {
		err = fmt.Errorf("marshal yaml: %w", err)
		panic(err)
	}
	return
}

// MarshalFile encodes the object to a yaml file.
// Saves the encoded yaml file or the error encountered.
func MarshalFile[T any](obj T, path string) (err error) {
	var data []byte
	data, err = yaml.Marshal(obj)
	if err != nil {
		err = fmt.Errorf("marshal yaml: %w", err)
		return
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		err = fmt.Errorf("write file '%s': %w", path, err)
		return
	}
	return
}

// MustMarshalFile encodes the object to a yaml file.
// Saves the encoded yaml file or panics if an error is encountered.
func MustMarshalFile[T any](obj T, path string) {
	var err error
	var data []byte
	data, err = yaml.Marshal(obj)
	if err != nil {
		err = fmt.Errorf("marshal yaml: %w", err)
		panic(err)
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		err = fmt.Errorf("write file '%s': %w", path, err)
		panic(err)
	}
	return
}
