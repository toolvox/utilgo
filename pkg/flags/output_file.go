package flags

import (
	"fmt"
	"io"
	"os"
)

// OutputFileValue represents a command-line flag that specifies an output file.
type OutputFileValue struct {
	Filename string // Path to the output file
	Flag     int    // File opening flag (e.g., [os.O_CREATE])
}

// NewOutputFileValue creates a new [OutputFileValue] with a default file and flag.
func NewOutputFileValue(defaultFile string, flag int) OutputFileValue {
	return OutputFileValue{Filename: defaultFile, Flag: flag}
}

// String returns the path of the output file.
func (o OutputFileValue) String() string {
	return o.Filename
}

// Set assigns a new value to the Filename field.
func (o *OutputFileValue) Set(value string) error {
	o.Filename = value
	return nil
}

// Get opens the file and returns the *[os.File] if it succeeds.
// If Filename is empty, returns [os.Stdout].
// If an error occurs, return thr error instead.
func (o OutputFileValue) Get() any {
	if o.Filename == "" {
		return os.Stdout
	}
	file, err := os.OpenFile(o.Filename, o.Flag, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file '%s': %w", o.Filename, err)
	}
	return file
}

// Writer returns an [io.WriteCloser] for the target of [OutputFileValue].
func (o OutputFileValue) Writer() (io.WriteCloser, error) {
	switch r := o.Get().(type) {
	case error:
		return nil, r
	case *os.File:
		return r, nil
	default:
		panic(fmt.Errorf("unknown result from Get(): %v - %v", o, r))
	}
}
