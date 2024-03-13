package yaml_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"utilgo/experimental/pkg/serialization/yaml"
)

const INVALID_FILE_WRITE_ERROR = "write file '?': open ?: The filename, directory name, or volume label syntax is incorrect."

func Test_YAML_Marshal_Sanity(t *testing.T) {
	tempDir := t.TempDir() // Using t.TempDir() to automatically clean up after the test
	testFilePath := filepath.Join(tempDir, "test.yaml")

	t.Run("Marshal", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			data, err := yaml.Marshal(demoObj)
			must.NoError(err)
			must.Equal(strings.TrimSpace(string(demoYaml)), strings.TrimSpace(string(data)))
		})
	})

	t.Run("MustMarshal", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				data := yaml.MustMarshal(demonObj)
				must.Equal(strings.TrimSpace(string(demoInvalidYaml)), strings.TrimSpace(string(data)))
			})
		})
	})

	t.Run("MarshalFile", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NoError(yaml.MarshalFile(demoObj, testFilePath))
			data, err := os.ReadFile(testFilePath)
			must.NoError(err)
			must.Equal(strings.TrimSpace(string(demoYaml)), strings.TrimSpace(string(data)))
		})

		t.Run("FileError", func(t *testing.T) {
			err := yaml.MarshalFile(demoObj, "?")
			require.Error(t, err)
		})
	})

	t.Run("MustMarshalFile", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				yaml.MustMarshalFile(demoObj, testFilePath)
				data, err := os.ReadFile(testFilePath)
				must.NoError(err)
				must.Equal(strings.TrimSpace(string(demoYaml)), strings.TrimSpace(string(data)))
			})
		})

		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(INVALID_FILE_WRITE_ERROR, func() {
				yaml.MustMarshalFile(demoObj, "?")
				data, err := os.ReadFile(testFilePath)
				must.NoError(err)
				must.Equal(strings.TrimSpace(string(demoYaml)), strings.TrimSpace(string(data)))
			})
		})
	})
}
