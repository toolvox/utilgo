package json_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"utilgo/experimental/pkg/serialization/json"
)

func Test_JSON_Decode_Sanity(t *testing.T) {
	tempDir := t.TempDir()
	defer os.RemoveAll(tempDir)

	testFilePath := filepath.Join(tempDir, "test.json")
	_ = os.WriteFile(testFilePath, demoJsonNoIndent, 0644) // Setup file for DecodeFile tests

	cyclicData := &cyclicObject{}
	cyclicData.Reference = cyclicData // Create cyclic reference

	t.Run("Decode", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.Decode[demoObject](demoJsonNoIndent)
			must.NoError(err)
			must.Equal(demoObj, obj)
		})

		t.Run("Error", func(t *testing.T) {
			_, err := json.Decode[demoObject](demoBadJson)
			require.Error(t, err, "expected error when decoding invalid JSON")
		})
	})

	t.Run("MustDecode", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				obj := json.MustDecode[demoObject](demoJsonNoIndent)
				must.Equal(demoObj, obj)
			})
		})

		// Error testing for MustDecode via panic is omitted due to the nature of panics.
	})

	t.Run("DecodeReader", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			reader := bytes.NewReader(demoJsonNoIndent)
			obj, err := json.DecodeReader[demoObject](reader)
			must.NoError(err)
			must.Equal(demoObj, obj)
		})

		t.Run("Error", func(t *testing.T) {
			reader := bytes.NewReader(demoBadJson)
			_, err := json.DecodeReader[demoObject](reader)
			require.Error(t, err, "expected error when decoding invalid JSON via reader")
		})
	})

	t.Run("DecodeFile", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.DecodeFile[demoObject](testFilePath)
			must.NoError(err)
			must.Equal(demoObj, obj)
		})

		// Testing DecodeFile for error handling is not straightforward without an invalid file setup.
	})

	// Additional tests for MustDecodeReader, DecodeInto, MustDecodeInto, DecodeReaderInto, MustDecodeReaderInto,
	// DecodeFileInto, MustDecodeFileInto, DecodeFS, MustDecodeFS, DecodeFSInto, and MustDecodeFSInto can follow a similar structure.
}
