package json_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"utilgo/experimental/pkg/serialization/json"
)

type cyclicObject struct {
	Reference *cyclicObject `json:"Reference"`
}

var demoJsonNoIndent = []byte(`{"Name":"Demo","Value":-667,"SubObject":{"Flag":true}}`)

const ILLEGAL_CYCLIC_JSON_ERROR = "marshal json: json: unsupported value: encountered a cycle via *json_test.cyclicObject"

func Test_JSON_Marshal_Sanity(t *testing.T) {
	tempDir := t.TempDir()
	testFilePath := filepath.Join(tempDir, "test.json")
	testIndentedFilePath := filepath.Join(tempDir, "test_indented.json")

	cyclicData := &cyclicObject{}
	cyclicData.Reference = cyclicData // create a cyclic reference to cause an error

	t.Run("Marshal", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			data, err := json.Marshal(demoObj)
			must.NoError(err)
			must.JSONEq(string(demoJsonNoIndent), string(data))
		})

		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			_, err := json.Marshal(cyclicData)
			must.Error(err)
			must.Equal(ILLEGAL_CYCLIC_JSON_ERROR, err.Error())
		})
	})

	t.Run("MustMarshal", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				data := json.MustMarshal(demoObj)
				must.JSONEq(string(demoJsonNoIndent), string(data))
			})
		})

		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(ILLEGAL_CYCLIC_JSON_ERROR, func() {
				data := json.MustMarshal(cyclicData)
				must.JSONEq(string(demoJsonNoIndent), string(data))
			})
		})
	})

	t.Run("MarshalIndent", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			data, err := json.MarshalIndent(demoObj)
			must.NoError(err)
			must.JSONEq(string(demoJson), string(data))
		})

		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			_, err := json.MarshalIndent(cyclicData)
			must.Error(err)
			must.Equal(ILLEGAL_CYCLIC_JSON_ERROR, err.Error())
		})
	})

	t.Run("MustMarshalIndent", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				data := json.MustMarshalIndent(demoObj)
				must.JSONEq(string(demoJson), string(data))
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(ILLEGAL_CYCLIC_JSON_ERROR, func() {
				data := json.MustMarshalIndent(cyclicData)
				must.JSONEq(string(demoJsonNoIndent), string(data))
			})
		})
	})

	t.Run("MarshalFile", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NoError(json.MarshalFile(demoObj, testFilePath))
			data, err := os.ReadFile(testFilePath)
			must.NoError(err)
			must.JSONEq(string(demoJsonNoIndent), string(data))
		})

		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			err := json.MarshalFile(cyclicData, testFilePath)
			must.Error(err)
			must.Equal(ILLEGAL_CYCLIC_JSON_ERROR, err.Error())
		})
	})

	t.Run("MustMarshalFile", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				json.MustMarshalFile(demoObj, testFilePath)
				data, err := os.ReadFile(testFilePath)
				must.NoError(err)
				must.JSONEq(string(demoJsonNoIndent), string(data))
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(ILLEGAL_CYCLIC_JSON_ERROR, func() {
				json.MustMarshalFile(cyclicData, testFilePath)
			})
		})
	})

	t.Run("MarshalFileIndent", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NoError(json.MarshalFileIndent(demoObj, testIndentedFilePath))
			data, err := os.ReadFile(testIndentedFilePath)
			must.NoError(err)
			must.JSONEq(string(demoJson), string(data))
		})

		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			err := json.MarshalFileIndent(cyclicData, testFilePath)
			must.Error(err)
			must.Equal(ILLEGAL_CYCLIC_JSON_ERROR, err.Error())
		})
	})

	t.Run("MustMarshalFileIndent", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				json.MustMarshalFileIndent(demoObj, testIndentedFilePath)
				data, err := os.ReadFile(testIndentedFilePath)
				must.NoError(err)
				must.JSONEq(string(demoJson), string(data))
			})
		})

		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(ILLEGAL_CYCLIC_JSON_ERROR, func() {
				json.MustMarshalFileIndent(cyclicData, testFilePath)
			})
		})
	})
}
