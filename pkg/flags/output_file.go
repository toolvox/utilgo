package flags

import (
	"io"
	"os"

	"utilgo/pkg/errs"
)

// OutputFileValue is a command-line flag that handles an output file.
//
// The file is opened when Set is called and fails if the file cannot be opened.
type OutputFileValue struct {
	Filename string
	Flag     int
	File     *os.File
}

// OutputFileDefault sets the default values of an [OutputFileDefault] flag, returning the pointer to use in:
//
//	flag.Var(flags.OutputFileDefault(&o.OutputFile, "default.txt", 0644), "output", "usage usage")
//
// `O_RDONLY (0)` is not a valid flag and will default to `os.O_CREATE | os.O_TRUNC`
func OutputFileDefault(outputFileValue *OutputFileValue, defaultPath string, defaultFlag int) *OutputFileValue {
	if defaultFlag == 0 {
		defaultFlag = os.O_CREATE | os.O_TRUNC
	}
	*outputFileValue = OutputFileValue{
		Filename: defaultPath,
		Flag:     defaultFlag,
	}
	return outputFileValue
}

// String returns the path of the output file.
//
// This method implements the [flag.Value] interface.
func (o OutputFileValue) String() string {
	if o.File != nil {
		return o.File.Name()
	}
	return o.Filename
}

// Set assign the File field depending on the value.
//
// Returns an error if a file is requested but the file cannot be opened using the provided flag.
//
// This method implements the [flag.Value] interface.
func (o *OutputFileValue) Set(value string) error {
	switch value {
	case "", "stdout":
		o.File = os.Stdout

	case "stderr":
		o.File = os.Stderr

	default:
		var err error
		if o.Flag == 0 {
			o.Flag = os.O_CREATE | os.O_TRUNC
		}
		o.File, err = os.OpenFile(value, o.Flag, 0644)
		if err != nil {
			return errs.Wrapf("unable to open file '%s'", value, err)
		}
	}
	o.Filename = value
	return nil
}

// Get returns the open *[os.File] if the call to [OutputFileValue.Set] succeeded and an error otherwise.
// [OutputFileValue.Get] will try to call [OutputFileValue.Set] if it wasn't called yet.
//
// This method implements the [flag.Getter] interface.
func (o OutputFileValue) Get() any {
	if o.File == nil {
		if err := o.Set(o.Filename); err != nil {
			return err
		}
	}
	return o.File
}

// Writer returns an [io.WriteCloser] for the target of [OutputFileValue].
// Or nil if [OutputFileValue.Set] failed.
// Uses return value from [FileValue.Get].
func (o OutputFileValue) Writer() io.WriteCloser {
	switch t := o.Get().(type) {
	case os.File:
		return &t
	
	case *os.File:
		return t
	
	case io.WriteCloser:
		return t

	default:
		return nil
	}
}
