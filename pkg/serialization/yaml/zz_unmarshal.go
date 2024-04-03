package yaml

import (
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/toolvox/utilgo/pkg/errs"
)

// Unmarshal parses the object encoded in the provided yaml byte slice.
// Returns the decoded object or the error encountered.
func Unmarshal[T any](data []byte) (obj T, err error) {
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	return
}

// MustUnmarshal parses the object encoded in the provided yaml byte slice.
// Returns the decoded object or panics if an error is encountered.
func MustUnmarshal[T any](data []byte) (obj T) {
	var err error
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	return
}

// UnmarshalValid parses the object encoded in the provided yaml byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func UnmarshalValid[T errs.Validator](data []byte) (obj T, err error) {
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustUnmarshalValid parses the object encoded in the provided yaml byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or panics if an error is encountered.
func MustUnmarshalValid[T errs.Validator](data []byte) (obj T) {
	var err error
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// UnmarshalInto parses the object encoded in the provided yaml byte slice.
// Assigns the decoded object to the object pointer or the error encountered.
func UnmarshalInto[T any](data []byte, obj T) (err error) {
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	return
}

// MustUnmarshalInto parses the object encoded in the provided yaml byte slice.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustUnmarshalInto[T any](data []byte, obj T) {
	var err error
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	return
}

// UnmarshalValidInto parses the object encoded in the provided yaml byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func UnmarshalValidInto[T errs.Validator](data []byte, obj T) (err error) {
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustUnmarshalValidInto parses the object encoded in the provided yaml byte slice.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustUnmarshalValidInto[T errs.Validator](data []byte, obj T) {
	var err error
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// UnmarshalFile parses the object encoded in the provided yaml file.
// Returns the decoded object or the error encountered.
func UnmarshalFile[T any](path string) (obj T, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		return
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	return
}

// MustUnmarshalFile parses the object encoded in the provided yaml file.
// Returns the decoded object or panics if an error is encountered.
func MustUnmarshalFile[T any](path string) (obj T) {
	var err error
	data, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		panic(err)
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	return
}

// UnmarshalValidFile parses the object encoded in the provided yaml file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func UnmarshalValidFile[T errs.Validator](path string) (obj T, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		return
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustUnmarshalValidFile parses the object encoded in the provided yaml file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or panics if an error is encountered.
func MustUnmarshalValidFile[T errs.Validator](path string) (obj T) {
	var err error
	data, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		panic(err)
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// UnmarshalFileInto parses the object encoded in the provided yaml file.
// Assigns the decoded object to the object pointer or the error encountered.
func UnmarshalFileInto[T any](path string, obj T) (err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		return
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	return
}

// MustUnmarshalFileInto parses the object encoded in the provided yaml file.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustUnmarshalFileInto[T any](path string, obj T) {
	var err error
	data, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		panic(err)
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	return
}

// UnmarshalValidFileInto parses the object encoded in the provided yaml file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func UnmarshalValidFileInto[T errs.Validator](path string, obj T) (err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		return
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustUnmarshalValidFileInto parses the object encoded in the provided yaml file.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustUnmarshalValidFileInto[T errs.Validator](path string, obj T) {
	var err error
	data, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		panic(err)
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// UnmarshalFS parses the object encoded in the provided yaml file from the provided FS.
// Returns the decoded object or the error encountered.
func UnmarshalFS[T any](path string, f fs.FS) (obj T, err error) {
	data, err := fs.ReadFile(f, path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		return
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	return
}

// MustUnmarshalFS parses the object encoded in the provided yaml file from the provided FS.
// Returns the decoded object or panics if an error is encountered.
func MustUnmarshalFS[T any](path string, f fs.FS) (obj T) {
	var err error
	data, err := fs.ReadFile(f, path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		panic(err)
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	return
}

// UnmarshalValidFS parses the object encoded in the provided yaml file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or the error encountered.
func UnmarshalValidFS[T errs.Validator](path string, f fs.FS) (obj T, err error) {
	data, err := fs.ReadFile(f, path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		return
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustUnmarshalValidFS parses the object encoded in the provided yaml file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Returns the decoded object or panics if an error is encountered.
func MustUnmarshalValidFS[T errs.Validator](path string, f fs.FS) (obj T) {
	var err error
	data, err := fs.ReadFile(f, path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		panic(err)
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}

// UnmarshalFSInto parses the object encoded in the provided yaml file from the provided FS.
// Assigns the decoded object to the object pointer or the error encountered.
func UnmarshalFSInto[T any](path string, f fs.FS, obj T) (err error) {
	data, err := fs.ReadFile(f, path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		return
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	return
}

// MustUnmarshalFSInto parses the object encoded in the provided yaml file from the provided FS.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustUnmarshalFSInto[T any](path string, f fs.FS, obj T) {
	var err error
	data, err := fs.ReadFile(f, path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		panic(err)
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	return
}

// UnmarshalValidFSInto parses the object encoded in the provided yaml file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or the error encountered.
func UnmarshalValidFSInto[T errs.Validator](path string, f fs.FS, obj T) (err error) {
	data, err := fs.ReadFile(f, path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		return
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		return
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		return
	}
	return
}

// MustUnmarshalValidFSInto parses the object encoded in the provided yaml file from the provided FS.
// Expects the decoded type to implement [pkg/utilgo/pkg/errs.Validator] and treats validation errors as decoding errors.
// Assigns the decoded object to the object pointer or panics if an error is encountered.
func MustUnmarshalValidFSInto[T errs.Validator](path string, f fs.FS, obj T) {
	var err error
	data, err := fs.ReadFile(f, path)
	if err != nil {
		err = fmt.Errorf("read file '%s': %w", path, err)
		panic(err)
	}
	if err = yaml.Unmarshal(data, &obj); err != nil {
		err = fmt.Errorf("unmarshal yaml: %w", err)
		panic(err)
	}
	if err = obj.Validate(); err != nil {
		err = fmt.Errorf("validation: %w", err)
		panic(err)
	}
	return
}
