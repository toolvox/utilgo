package json

import (
	"encoding/json"
	"fmt"
	"os"
)

// Marshal encodes the object to a json byte slice.
// Returns the encoded json byte slice or the error encountered.
func Marshal[T any](obj T) (data []byte, err error) {
	data, err = json.Marshal(obj)
	if err != nil {
		err = fmt.Errorf("marshal json: %w", err)
		return
	}
	return
}

// MustMarshal encodes the object to a json byte slice.
// Returns the encoded json byte slice or panics if an error is encountered.
func MustMarshal[T any](obj T) (data []byte) {
	var err error
	data, err = json.Marshal(obj)
	if err != nil {
		err = fmt.Errorf("marshal json: %w", err)
		panic(err)
	}
	return
}

// MarshalIndent encodes the object to an indented json byte slice.
// Returns the encoded json byte slice or the error encountered.
func MarshalIndent[T any](obj T) (data []byte, err error) {
	data, err = json.MarshalIndent(obj, "", "\t")
	if err != nil {
		err = fmt.Errorf("marshal json: %w", err)
		return
	}
	return
}

// MustMarshalIndent encodes the object to an indented json byte slice.
// Returns the encoded json byte slice or panics if an error is encountered.
func MustMarshalIndent[T any](obj T) (data []byte) {
	var err error
	data, err = json.MarshalIndent(obj, "", "\t")
	if err != nil {
		err = fmt.Errorf("marshal json: %w", err)
		panic(err)
	}
	return
}

// MarshalFile encodes the object to a json file.
// Saves the encoded json file or the error encountered.
func MarshalFile[T any](obj T, path string) (err error) {
	var data []byte
	data, err = json.Marshal(obj)
	if err != nil {
		err = fmt.Errorf("marshal json: %w", err)
		return
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		err = fmt.Errorf("write file '%s': %w", path, err)
		return
	}
	return
}

// MustMarshalFile encodes the object to a json file.
// Saves the encoded json file or panics if an error is encountered.
func MustMarshalFile[T any](obj T, path string) {
	var err error
	var data []byte
	data, err = json.Marshal(obj)
	if err != nil {
		err = fmt.Errorf("marshal json: %w", err)
		panic(err)
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		err = fmt.Errorf("write file '%s': %w", path, err)
		panic(err)
	}
	return
}

// MarshalFileIndent encodes the object to an indented json file.
// Saves the encoded json file or the error encountered.
func MarshalFileIndent[T any](obj T, path string) (err error) {
	var data []byte
	data, err = json.MarshalIndent(obj, "", "\t")
	if err != nil {
		err = fmt.Errorf("marshal json: %w", err)
		return
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		err = fmt.Errorf("write file '%s': %w", path, err)
		return
	}
	return
}

// MustMarshalFileIndent encodes the object to an indented json file.
// Saves the encoded json file or panics if an error is encountered.
func MustMarshalFileIndent[T any](obj T, path string) {
	var err error
	var data []byte
	data, err = json.MarshalIndent(obj, "", "\t")
	if err != nil {
		err = fmt.Errorf("marshal json: %w", err)
		panic(err)
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		err = fmt.Errorf("write file '%s': %w", path, err)
		panic(err)
	}
	return
}
