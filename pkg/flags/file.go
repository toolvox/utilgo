// Package flags introduces supplementary [flag.Value] implementations for utilization with the [flag.Var](...) function.
package flags

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// FileValue is a command-line flag that handles an input file.
//
// The file is opened when Set is called and fails if the file cannot be opened.
type FileValue struct {
	Filename string
	Content  []byte
}

// FileDefault sets the default value of a [FileValue] flag, returning the pointer for use in:
//
//	flag.Var(flags.FileDefault(&o.InputFile, "default.txt"), "input", "usage usage")
func FileDefault(fileValue *FileValue, def string) *FileValue {
	*fileValue = FileValue{
		Filename: def,
		Content:  []byte{},
	}
	return fileValue
}

// String returns the input path of the file.
//
// This method implements the [flag.Value] interface.
func (f FileValue) String() string {
	return fmt.Sprintf(`"%s"`, f.Filename)
}

// Set reads the file named by value and stores its content.
//
// The value is saved even if the file fails to open.
//
// This method implements the [flag.Value] interface.
func (f *FileValue) Set(value string) error {
	if value == "" {
		return nil
	}
	content, err := os.ReadFile(value)
	f.Filename = value
	if err != nil {
		return fmt.Errorf("unable to read file '%s': %w", value, err)
	}

	f.Content = content
	return nil
}

// Get returns the content of the file.
// [FileValue.Get] will try to call [FileValue.Set] if it wasn't called yet.
//
// The return type is always []byte.
//
// This method implements the [flag.Getter] interface.
func (f FileValue) Get() any {
	if len(f.Content) == 0 {
		if err := f.Set(f.Filename); err != nil {
			return err
		}
	}
	return f.Content
}

// Reader returns an [io.Reader] for the content of the file.
// Uses return value from [FileValue.Get].
func (f FileValue) Reader() io.Reader {
	switch t := f.Get().(type) {
	case []byte:
		return bytes.NewReader(t)

	default:
		return nil
	}
}
