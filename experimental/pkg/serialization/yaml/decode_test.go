package yaml_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"utilgo/experimental/pkg/serialization/yaml"
)

func Test_YAML_Decode_Sanity(t *testing.T) {
	var obj demoObject
	var objWithError demoObject

	tempDir := t.TempDir()
	defer os.RemoveAll(tempDir)

	demoFile, err := os.CreateTemp(tempDir, "demo")
	if err != nil {
		t.Fatal(err)
	}
	defer demoFile.Close()
	demoFile.WriteString(string(demoYaml))

	demonFile, err := os.CreateTemp(tempDir, "demon")
	if err != nil {
		t.Fatal(err)
	}
	defer demonFile.Close()
	demonFile.WriteString(string(demoInvalidYaml))

	badFile, err := os.CreateTemp(tempDir, "bad")
	if err != nil {
		t.Fatal(err)
	}
	defer badFile.Close()
	badFile.WriteString(string(demoBadYaml))
	t.Run("Decode", func(t *testing.T) {
		obj, err := yaml.Decode[demoObject](demoYaml)
		require.NoError(t, err)
		assert.Equal(t, demoObj, obj)
	})

	t.Run("MustDecode", func(t *testing.T) {
		assert.NotPanics(t, func() {
			obj := yaml.MustDecode[demoObject](demoYaml)
			assert.Equal(t, demoObj, obj)
		})
	})

	t.Run("DecodeWithError", func(t *testing.T) {
		_, err := yaml.Decode[demoObject](demoBadYaml)
		assert.Error(t, err)
	})

	t.Run("MustDecodeWithError", func(t *testing.T) {
		assert.Panics(t, func() {
			yaml.MustDecode[demoObject](demoBadYaml)
		})
	})

	t.Run("DecodeFile", func(t *testing.T) {
		obj, err := yaml.DecodeFile[demoObject](demoFile.Name())
		require.NoError(t, err)
		assert.Equal(t, demoObj, obj)
	})

	t.Run("MustDecodeFile", func(t *testing.T) {
		assert.NotPanics(t, func() {
			obj := yaml.MustDecodeFile[demoObject](demoFile.Name())
			assert.Equal(t, demoObj, obj)
		})
	})

	t.Run("DecodeReader", func(t *testing.T) {
		reader := bytes.NewReader(demoYaml)
		obj, err := yaml.DecodeReader[demoObject](reader) // Assuming your function signature.
		require.NoError(t, err)
		assert.Equal(t, demoObj, obj)
	})

	t.Run("MustDecodeReader", func(t *testing.T) {
		reader := bytes.NewReader(demoYaml)
		assert.NotPanics(t, func() {
			obj := yaml.MustDecodeReader[demoObject](reader) // Assuming your function signature.
			assert.Equal(t, demoObj, obj)
		})
	})

	t.Run("DecodeReaderWithError", func(t *testing.T) {
		reader := bytes.NewReader(demoBadYaml)
		_, err := yaml.DecodeReader[demoObject](reader)
		assert.Error(t, err)
	})

	t.Run("MustDecodeReaderWithError", func(t *testing.T) {
		reader := bytes.NewReader(demoBadYaml)
		assert.Panics(t, func() {
			yaml.MustDecodeReader[demoObject](reader)
		})
	})

	t.Run("DecodeInto", func(t *testing.T) {
		err := yaml.DecodeInto(demoYaml, &obj) // Assuming your function signature.
		require.NoError(t, err)
		assert.Equal(t, demoObj, obj)
	})

	var objForMust demoObject
	t.Run("MustDecodeInto", func(t *testing.T) {
		assert.NotPanics(t, func() {
			yaml.MustDecodeInto(demoYaml, &objForMust) // Assuming your function signature.
			assert.Equal(t, demoObj, objForMust)
		})
	})

	t.Run("DecodeIntoWithError", func(t *testing.T) {
		err := yaml.DecodeInto(demoBadYaml, &objWithError)
		assert.Error(t, err)
	})

	t.Run("MustDecodeIntoWithError", func(t *testing.T) {
		assert.Panics(t, func() {
			yaml.MustDecodeInto(demoBadYaml, &objWithError)
		})
	})

	t.Run("DecodeFS", func(t *testing.T) {
		obj, err := yaml.DecodeFS[demoObject](filepath.Base(demoFile.Name()), os.DirFS(tempDir)) // Adjust path and FS as needed.
		require.NoError(t, err)
		assert.Equal(t, demoObj, obj)
	})

	t.Run("MustDecodeFS", func(t *testing.T) {
		assert.NotPanics(t, func() {
			obj := yaml.MustDecodeFS[demoObject](filepath.Base(demoFile.Name()), os.DirFS(tempDir)) // Adjust path and FS as needed.
			assert.Equal(t, demoObj, obj)
		})
	})

	t.Run("DecodeFSWithError", func(t *testing.T) {
		_, err := yaml.DecodeFS[demoObject](filepath.Base(badFile.Name()), os.DirFS(tempDir)) // Adjust path and FS as needed.
		assert.Error(t, err)
	})

	t.Run("MustDecodeFSWithError", func(t *testing.T) {
		assert.Panics(t, func() {
			yaml.MustDecodeFS[demoObject](filepath.Base(badFile.Name()), os.DirFS(tempDir)) // Adjust path and FS as needed.
		})
	})
}
