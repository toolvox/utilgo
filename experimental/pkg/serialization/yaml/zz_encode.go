package yaml

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// Encode uses [yaml.Encoder] to encode the object to a yaml byte slice.
// Returns the encoded yaml byte slice or the error encountered.
func Encode[T any](obj T) (data []byte, err error) {
	writer := bytes.NewBuffer(data)
	encoder := yaml.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode yaml: %w", err)
		return
	}
	return
}

// MustEncode uses [yaml.Encoder] to encode the object to a yaml byte slice.
// Returns the encoded yaml byte slice or panics if an error is encountered.
func MustEncode[T any](obj T) (data []byte) {
	var err error
	writer := bytes.NewBuffer(data)
	encoder := yaml.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode yaml: %w", err)
		panic(err)
	}
	return
}

// EncodeFile uses [yaml.Encoder] to encode the object to a yaml file.
// Saves the encoded yaml file or the error encountered.
func EncodeFile[T any](obj T, path string) (err error) {
	writer, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer writer.Close()
	encoder := yaml.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode yaml: %w", err)
		return
	}
	return
}

// MustEncodeFile uses [yaml.Encoder] to encode the object to a yaml file.
// Saves the encoded yaml file or panics if an error is encountered.
func MustEncodeFile[T any](obj T, path string) {
	var err error
	writer, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		panic(err)
	}
	defer writer.Close()
	encoder := yaml.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode yaml: %w", err)
		panic(err)
	}
	return
}

// EncodeWriter uses [yaml.Encoder] with the provided [io.writer] to encode the object to a yaml writer's data.
// Writes the encoded yaml writer's data or the error encountered.
func EncodeWriter[T any](obj T, writer io.Writer) (err error) {
	encoder := yaml.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode yaml: %w", err)
		return
	}
	return
}

// MustEncodeWriter uses [yaml.Encoder] with the provided [io.writer] to encode the object to a yaml writer's data.
// Writes the encoded yaml writer's data or panics if an error is encountered.
func MustEncodeWriter[T any](obj T, writer io.Writer) {
	var err error
	encoder := yaml.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode yaml: %w", err)
		panic(err)
	}
	return
}
