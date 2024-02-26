package flags_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"utilgo/pkg/flags"
)

func TestFileValue(t *testing.T) {
	tests := []struct {
		name            string
		setup           func(t *testing.T) (string, func())
		content         string
		expectNameCheck bool
		expectedErr     bool
	}{
		{
			name: "valid file set and get",
			setup: func(t *testing.T) (string, func()) {
				content := "Hello, test!"
				tmpfile, err := ioutil.TempFile("", "example")
				if err != nil {
					t.Fatalf("Failed to create temp file: %v", err)
				}
				if _, err := tmpfile.WriteString(content); err != nil {
					t.Fatalf("Failed to write to temp file: %v", err)
				}
				if err := tmpfile.Close(); err != nil {
					t.Fatalf("Failed to close temp file: %v", err)
				}
				return tmpfile.Name(), func() { os.Remove(tmpfile.Name()) }
			},
			content:         "Hello, test!",
			expectNameCheck: true,
			expectedErr:     false,
		},
		{
			name: "file not found",
			setup: func(t *testing.T) (string, func()) {
				return "/path/to/nonexistent/file", func() {}
			},
			content:         "",
			expectNameCheck: false,
			expectedErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename, cleanup := tt.setup(t)
			defer cleanup()

			fv := &flags.FileValue{}
			err := fv.Set(filename)

			if (err != nil) != tt.expectedErr {
				t.Fatalf("FileValue.Set() error = %v, wantErr %v", err, tt.expectedErr)
			}

			if tt.expectNameCheck && fv.Filename != filename {
				t.Errorf("Expected filename %s, got %s", filename, fv.Filename)
			}

			if !tt.expectedErr {
				content := fv.Get().([]byte)
				if !reflect.DeepEqual(content, []byte(tt.content)) {
					t.Errorf("Expected content \"%s\", got \"%s\"", tt.content, string(content))
				}

				readContent, readErr := ioutil.ReadAll(fv.Reader())
				if readErr != nil {
					t.Fatalf("Failed to read file: %v", readErr)
				}
				if string(readContent) != tt.content {
					t.Errorf("Expected content \"%s\", got \"%s\"", tt.content, string(readContent))
				}
			}
		})
	}
}
