package flags

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// FileValue holds information about a file specified by a flag.
// It stores the name of the file and its content in memory.
type FileValue struct {
	Filename string // Path to the input file
	Content  []byte // The content read from the file
}

// String returns the name of the file.
func (fv FileValue) String() string {
	return fv.Filename
}

// Set reads the file named by value and stores its content.
// The value is saved even if the file fails to open.
// This method implements the [flag.Value] interface.
func (fv *FileValue) Set(value string) error {
	content, err := os.ReadFile(value)
	fv.Filename = value
	if err != nil {
		return fmt.Errorf("unable to read file '%s': %w", value, err)
	}

	fv.Content = content
	return nil
}

// Get returns the content of the file.
// This method implements the [flag.Getter] interface.
// The return type is always []byte.
func (fv FileValue) Get() any {
	return fv.Content
}

// Reader returns an [io.Reader] for the content of the file.
func (fv FileValue) Reader() io.Reader {
	return bytes.NewReader(fv.Content)
}
