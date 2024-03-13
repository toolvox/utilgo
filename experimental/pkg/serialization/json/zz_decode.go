package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"

	"utilgo/pkg/errs"
)

// Decode uses [json.Decoder] to decode the object encoded in the json byte slice.
// Returns the decoded object or the error encountered.
func Decode[T any](data []byte) (obj T, err error) {
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	return
}

// MustDecode uses [json.Decoder] to decode the object encoded in the json byte slice.
// Returns the decoded object or panics if an error is encountered.
func MustDecode[T any](data []byte) (obj T) {
	var err error
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	return
}

// DecodeValid uses [json.Decoder] to decode the object encoded in the json byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func DecodeValid[T errs.Validator](data []byte) (obj T, err error) {
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValid uses [json.Decoder] to decode the object encoded in the json byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeValid[T errs.Validator](data []byte) (obj T) {
	var err error
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeInto uses [json.Decoder] to decode the object encoded in the json byte slice.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeInto[T any](data []byte, obj T) (err error) {
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	return
}

// MustDecodeInto uses [json.Decoder] to decode the object encoded in the json byte slice.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeInto[T any](data []byte, obj T) {
	var err error
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	return
}

// DecodeValidInto uses [json.Decoder] to decode the object encoded in the json byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeValidInto[T errs.Validator](data []byte, obj T) (err error) {
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidInto uses [json.Decoder] to decode the object encoded in the json byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeValidInto[T errs.Validator](data []byte, obj T) {
	var err error
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeFile uses [json.Decoder] to decode the object encoded in the json file.
// Returns the decoded object or the error encountered.
func DecodeFile[T any](path string) (obj T, err error) {
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	return
}

// MustDecodeFile uses [json.Decoder] to decode the object encoded in the json file.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeFile[T any](path string) (obj T) {
	var err error
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	return
}

// DecodeValidFile uses [json.Decoder] to decode the object encoded in the json file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func DecodeValidFile[T errs.Validator](path string) (obj T, err error) {
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidFile uses [json.Decoder] to decode the object encoded in the json file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeValidFile[T errs.Validator](path string) (obj T) {
	var err error
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeFileInto uses [json.Decoder] to decode the object encoded in the json file.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeFileInto[T any](path string, obj T) (err error) {
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	return
}

// MustDecodeFileInto uses [json.Decoder] to decode the object encoded in the json file.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeFileInto[T any](path string, obj T) {
	var err error
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	return
}

// DecodeValidFileInto uses [json.Decoder] to decode the object encoded in the json file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeValidFileInto[T errs.Validator](path string, obj T) (err error) {
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidFileInto uses [json.Decoder] to decode the object encoded in the json file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeValidFileInto[T errs.Validator](path string, obj T) {
	var err error
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeFS uses [json.Decoder] to decode the object encoded in the json file from the provided FS.
// Returns the decoded object or the error encountered.
func DecodeFS[T any](path string, f fs.FS) (obj T, err error) {
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		return
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	return
}

// MustDecodeFS uses [json.Decoder] to decode the object encoded in the json file from the provided FS.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeFS[T any](path string, f fs.FS) (obj T) {
	var err error
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	return
}

// DecodeValidFS uses [json.Decoder] to decode the object encoded in the json file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func DecodeValidFS[T errs.Validator](path string, f fs.FS) (obj T, err error) {
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		return
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidFS uses [json.Decoder] to decode the object encoded in the json file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeValidFS[T errs.Validator](path string, f fs.FS) (obj T) {
	var err error
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeFSInto uses [json.Decoder] to decode the object encoded in the json file from the provided FS.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeFSInto[T any](path string, f fs.FS, obj T) (err error) {
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		return
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	return
}

// MustDecodeFSInto uses [json.Decoder] to decode the object encoded in the json file from the provided FS.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeFSInto[T any](path string, f fs.FS, obj T) {
	var err error
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	return
}

// DecodeValidFSInto uses [json.Decoder] to decode the object encoded in the json file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeValidFSInto[T errs.Validator](path string, f fs.FS, obj T) (err error) {
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		return
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidFSInto uses [json.Decoder] to decode the object encoded in the json file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeValidFSInto[T errs.Validator](path string, f fs.FS, obj T) {
	var err error
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeReader uses [json.Decoder] with the provided [io.reader] to decode the object encoded in the json data read by the reader.
// Returns the decoded object or the error encountered.
func DecodeReader[T any](reader io.Reader) (obj T, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	return
}

// MustDecodeReader uses [json.Decoder] with the provided [io.reader] to decode the object encoded in the json data read by the reader.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeReader[T any](reader io.Reader) (obj T) {
	var err error
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	return
}

// DecodeValidReader uses [json.Decoder] with the provided [io.reader] to decode the object encoded in the json data read by the reader.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func DecodeValidReader[T errs.Validator](reader io.Reader) (obj T, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidReader uses [json.Decoder] with the provided [io.reader] to decode the object encoded in the json data read by the reader.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeValidReader[T errs.Validator](reader io.Reader) (obj T) {
	var err error
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeReaderInto uses [json.Decoder] with the provided [io.reader] to decode the object encoded in the json data read by the reader.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeReaderInto[T any](reader io.Reader, obj T) (err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	return
}

// MustDecodeReaderInto uses [json.Decoder] with the provided [io.reader] to decode the object encoded in the json data read by the reader.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeReaderInto[T any](reader io.Reader, obj T) {
	var err error
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	return
}

// DecodeValidReaderInto uses [json.Decoder] with the provided [io.reader] to decode the object encoded in the json data read by the reader.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeValidReaderInto[T errs.Validator](reader io.Reader, obj T) (err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidReaderInto uses [json.Decoder] with the provided [io.reader] to decode the object encoded in the json data read by the reader.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeValidReaderInto[T errs.Validator](reader io.Reader, obj T) {
	var err error
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}
