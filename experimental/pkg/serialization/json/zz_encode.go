package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Encode uses [json.Encoder] to encode the object to a json byte slice.
// Returns the encoded json byte slice or the error encountered.
func Encode[T any](obj T) (data []byte, err error) {
	writer := bytes.NewBuffer(data)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		return
	}
	return
}

// MustEncode uses [json.Encoder] to encode the object to a json byte slice.
// Returns the encoded json byte slice or panics if an error is encountered.
func MustEncode[T any](obj T) (data []byte) {
	var err error
	writer := bytes.NewBuffer(data)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		panic(err)
	}
	return
}

// EncodeIndent uses [json.Encoder] to encode the object to an indented json byte slice.
// Returns the encoded json byte slice or the error encountered.
func EncodeIndent[T any](obj T) (data []byte, err error) {
	writer := bytes.NewBuffer(data)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		return
	}
	return
}

// MustEncodeIndent uses [json.Encoder] to encode the object to an indented json byte slice.
// Returns the encoded json byte slice or panics if an error is encountered.
func MustEncodeIndent[T any](obj T) (data []byte) {
	var err error
	writer := bytes.NewBuffer(data)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		panic(err)
	}
	return
}

// EncodeFile uses [json.Encoder] to encode the object to a json file.
// Saves the encoded json file or the error encountered.
func EncodeFile[T any](obj T, path string) (err error) {
	writer, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer writer.Close()
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		return
	}
	return
}

// MustEncodeFile uses [json.Encoder] to encode the object to a json file.
// Saves the encoded json file or panics if an error is encountered.
func MustEncodeFile[T any](obj T, path string) {
	var err error
	writer, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		panic(err)
	}
	defer writer.Close()
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		panic(err)
	}
	return
}

// EncodeFileIndent uses [json.Encoder] to encode the object to an indented json file.
// Saves the encoded json file or the error encountered.
func EncodeFileIndent[T any](obj T, path string) (err error) {
	writer, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer writer.Close()
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		return
	}
	return
}

// MustEncodeFileIndent uses [json.Encoder] to encode the object to an indented json file.
// Saves the encoded json file or panics if an error is encountered.
func MustEncodeFileIndent[T any](obj T, path string) {
	var err error
	writer, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		panic(err)
	}
	defer writer.Close()
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		panic(err)
	}
	return
}

// EncodeWriter uses [json.Encoder] with the provided [io.writer] to encode the object to a json writer's data.
// Writes the encoded json writer's data or the error encountered.
func EncodeWriter[T any](obj T, writer io.Writer) (err error) {
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		return
	}
	return
}

// MustEncodeWriter uses [json.Encoder] with the provided [io.writer] to encode the object to a json writer's data.
// Writes the encoded json writer's data or panics if an error is encountered.
func MustEncodeWriter[T any](obj T, writer io.Writer) {
	var err error
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		panic(err)
	}
	return
}

// EncodeWriterIndent uses [json.Encoder] with the provided [io.writer] to encode the object to an indented json writer's data.
// Writes the encoded json writer's data or the error encountered.
func EncodeWriterIndent[T any](obj T, writer io.Writer) (err error) {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		return
	}
	return
}

// MustEncodeWriterIndent uses [json.Encoder] with the provided [io.writer] to encode the object to an indented json writer's data.
// Writes the encoded json writer's data or panics if an error is encountered.
func MustEncodeWriterIndent[T any](obj T, writer io.Writer) {
	var err error
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(obj)
	if err != nil {
		err = fmt.Errorf("encode json: %w", err)
		panic(err)
	}
	return
}
