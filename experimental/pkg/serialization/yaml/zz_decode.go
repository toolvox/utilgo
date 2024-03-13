package yaml

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"

	"utilgo/pkg/errs"
)

// Decode uses [yaml.Decoder] to decode the object encoded in the yaml byte slice.
// Returns the decoded object or the error encountered.
func Decode[T any](data []byte) (obj T, err error) {
	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	return
}

// MustDecode uses [yaml.Decoder] to decode the object encoded in the yaml byte slice.
// Returns the decoded object or panics if an error is encountered.
func MustDecode[T any](data []byte) (obj T) {
	var err error
	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	return
}

// DecodeValid uses [yaml.Decoder] to decode the object encoded in the yaml byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func DecodeValid[T errs.Validator](data []byte) (obj T, err error) {
	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValid uses [yaml.Decoder] to decode the object encoded in the yaml byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeValid[T errs.Validator](data []byte) (obj T) {
	var err error
	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeInto uses [yaml.Decoder] to decode the object encoded in the yaml byte slice.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeInto[T any](data []byte, obj T) (err error) {
	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	return
}

// MustDecodeInto uses [yaml.Decoder] to decode the object encoded in the yaml byte slice.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeInto[T any](data []byte, obj T) {
	var err error
	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	return
}

// DecodeValidInto uses [yaml.Decoder] to decode the object encoded in the yaml byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeValidInto[T errs.Validator](data []byte, obj T) (err error) {
	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidInto uses [yaml.Decoder] to decode the object encoded in the yaml byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeValidInto[T errs.Validator](data []byte, obj T) {
	var err error
	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeFile uses [yaml.Decoder] to decode the object encoded in the yaml file.
// Returns the decoded object or the error encountered.
func DecodeFile[T any](path string) (obj T, err error) {
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	return
}

// MustDecodeFile uses [yaml.Decoder] to decode the object encoded in the yaml file.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeFile[T any](path string) (obj T) {
	var err error
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	return
}

// DecodeValidFile uses [yaml.Decoder] to decode the object encoded in the yaml file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func DecodeValidFile[T errs.Validator](path string) (obj T, err error) {
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidFile uses [yaml.Decoder] to decode the object encoded in the yaml file.
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
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeFileInto uses [yaml.Decoder] to decode the object encoded in the yaml file.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeFileInto[T any](path string, obj T) (err error) {
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	return
}

// MustDecodeFileInto uses [yaml.Decoder] to decode the object encoded in the yaml file.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeFileInto[T any](path string, obj T) {
	var err error
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	return
}

// DecodeValidFileInto uses [yaml.Decoder] to decode the object encoded in the yaml file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeValidFileInto[T errs.Validator](path string, obj T) (err error) {
	reader, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file: %w", err)
		return
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidFileInto uses [yaml.Decoder] to decode the object encoded in the yaml file.
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
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeFS uses [yaml.Decoder] to decode the object encoded in the yaml file from the provided FS.
// Returns the decoded object or the error encountered.
func DecodeFS[T any](path string, f fs.FS) (obj T, err error) {
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		return
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	return
}

// MustDecodeFS uses [yaml.Decoder] to decode the object encoded in the yaml file from the provided FS.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeFS[T any](path string, f fs.FS) (obj T) {
	var err error
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	return
}

// DecodeValidFS uses [yaml.Decoder] to decode the object encoded in the yaml file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func DecodeValidFS[T errs.Validator](path string, f fs.FS) (obj T, err error) {
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		return
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidFS uses [yaml.Decoder] to decode the object encoded in the yaml file from the provided FS.
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
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeFSInto uses [yaml.Decoder] to decode the object encoded in the yaml file from the provided FS.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeFSInto[T any](path string, f fs.FS, obj T) (err error) {
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		return
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	return
}

// MustDecodeFSInto uses [yaml.Decoder] to decode the object encoded in the yaml file from the provided FS.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeFSInto[T any](path string, f fs.FS, obj T) {
	var err error
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		panic(err)
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	return
}

// DecodeValidFSInto uses [yaml.Decoder] to decode the object encoded in the yaml file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeValidFSInto[T errs.Validator](path string, f fs.FS, obj T) (err error) {
	reader, err := f.Open(path)
	if err != nil {
		err = fmt.Errorf("open fs file: %w", err)
		return
	}
	defer reader.Close()
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidFSInto uses [yaml.Decoder] to decode the object encoded in the yaml file from the provided FS.
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
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeReader uses [yaml.Decoder] with the provided [io.reader] to decode the object encoded in the yaml data read by the reader.
// Returns the decoded object or the error encountered.
func DecodeReader[T any](reader io.Reader) (obj T, err error) {
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	return
}

// MustDecodeReader uses [yaml.Decoder] with the provided [io.reader] to decode the object encoded in the yaml data read by the reader.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeReader[T any](reader io.Reader) (obj T) {
	var err error
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	return
}

// DecodeValidReader uses [yaml.Decoder] with the provided [io.reader] to decode the object encoded in the yaml data read by the reader.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func DecodeValidReader[T errs.Validator](reader io.Reader) (obj T, err error) {
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidReader uses [yaml.Decoder] with the provided [io.reader] to decode the object encoded in the yaml data read by the reader.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or panics if an error is encountered.
func MustDecodeValidReader[T errs.Validator](reader io.Reader) (obj T) {
	var err error
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// DecodeReaderInto uses [yaml.Decoder] with the provided [io.reader] to decode the object encoded in the yaml data read by the reader.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeReaderInto[T any](reader io.Reader, obj T) (err error) {
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	return
}

// MustDecodeReaderInto uses [yaml.Decoder] with the provided [io.reader] to decode the object encoded in the yaml data read by the reader.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeReaderInto[T any](reader io.Reader, obj T) {
	var err error
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	return
}

// DecodeValidReaderInto uses [yaml.Decoder] with the provided [io.reader] to decode the object encoded in the yaml data read by the reader.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func DecodeValidReaderInto[T errs.Validator](reader io.Reader, obj T) (err error) {
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustDecodeValidReaderInto uses [yaml.Decoder] with the provided [io.reader] to decode the object encoded in the yaml data read by the reader.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustDecodeValidReaderInto[T errs.Validator](reader io.Reader, obj T) {
	var err error
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&obj)
	if err != nil {
		err = fmt.Errorf("decode yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}
