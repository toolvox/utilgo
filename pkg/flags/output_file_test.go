package flags_test

import (
	"os"
	"testing"

	"utilgo/pkg/flags"
)

func TestOutputFileValue_Get(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	tests := []struct {
		name      string
		ofv       flags.OutputFileValue
		errGet    bool
		errWriter bool
	}{
		{
			name:      "stdout",
			ofv:       flags.NewOutputFileValue("", os.O_RDWR|os.O_CREATE),
			errGet:    false,
			errWriter: false,
		},
		{
			name:      "valid file",
			ofv:       flags.NewOutputFileValue(tmpfile.Name(), os.O_RDWR|os.O_CREATE),
			errGet:    false,
			errWriter: false,
		},
		{
			name:      "invalid file",
			ofv:       flags.NewOutputFileValue("/invalid/path", os.O_RDWR|os.O_CREATE),
			errGet:    true,
			errWriter: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ofv.Get()
			if _, ok := got.(error); ok != tt.errGet {
				t.Errorf("OutputFileValue.Get() error = %v, wantErr %v", got, tt.errGet)
			}
			_, err := tt.ofv.Writer()
			if (err != nil) != tt.errWriter {
				t.Errorf("OutputFileValue.Writer() error = %v, wantErr %v", err, tt.errWriter)
			}
		})
	}
}
