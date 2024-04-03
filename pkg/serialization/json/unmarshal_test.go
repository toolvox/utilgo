package json_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/pkg/serialization/json"

	"github.com/toolvox/utilgo/pkg/errs"
)

type demoSubObject struct {
	Flag bool `json:"Flag"`
}

type demoObject struct {
	Name      string        `json:"Name"`
	Value     int           `json:"Value"`
	SubObject demoSubObject `json:"SubObject"`
}

var demoObj = demoObject{
	Name:      "Demo",
	Value:     -667,
	SubObject: demoSubObject{Flag: true},
}

var demonObj = demoObject{
	Name:      "Demon",
	Value:     666,
	SubObject: demoSubObject{Flag: true},
}

var validDemoObj = demoValidator(demoObj)
var validDemonObj = demoValidator(demonObj)

type demoValidator demoObject

func (v demoValidator) Validate() error {
	var errs errs.Errors
	if v.Name == "Demon" {
		errs = append(errs, fmt.Errorf("'Name' must not be demonic, was: %s", v.Name))
	}
	if v.Value == 666 {
		errs = append(errs, fmt.Errorf("'Value' must not be demonic, was: %d", v.Value))
	}
	return errs.OrNil()
}

var demoJson = []byte(`{
	"Name": "Demo",
	"Value": -667,
	"SubObject": {
		"Flag": true
	}
}`)

var demoInvalidJson = []byte(`{
	"Name": "Demon",
	"Value": 666,
	"SubObject": {
		"Flag": true
	}
}`)

var demoBadJson = []byte(`{
	"Name": Bad,
	"Value": Values,
	"SubObject": {
		"Flag": :O
	}
}`)

func Test_JSON_Unmarshal_Sanity(t *testing.T) {
	_, demoSyntaxError := json.Unmarshal[demoObject](demoBadJson)
	_, demonValidationError := json.UnmarshalValid[demoValidator](demoInvalidJson)
	_, readFileError := os.ReadFile("x")
	readFileError = fmt.Errorf("read file 'x': %w", readFileError)

	tempDir := t.TempDir()

	demoFile, err := os.CreateTemp(tempDir, "demo")
	if err != nil {
		t.Fatal(err)
	}
	defer demoFile.Close()
	demoFile.WriteString(string(demoJson))

	demonFile, err := os.CreateTemp(tempDir, "demon")
	if err != nil {
		t.Fatal(err)
	}
	defer demonFile.Close()
	demonFile.WriteString(string(demoInvalidJson))

	badFile, err := os.CreateTemp(tempDir, "bad")
	if err != nil {
		t.Fatal(err)
	}
	defer badFile.Close()
	badFile.WriteString(string(demoBadJson))

	tempFS := os.DirFS(tempDir)

	t.Run("Unmarshal", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.Unmarshal[demoObject](demoInvalidJson)
			must.NoError(err)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.Unmarshal[demoObject](demoBadJson)
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.Unmarshal[demoValidator](demoJson)
			must.NoError(err)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalValid[demoValidator](demoInvalidJson)
			must.Error(err)
			must.Equal(demonValidationError, err)
			must.Equal(validDemonObj, obj)
		})
	})

	t.Run("MustUnmarshal", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				obj := json.MustUnmarshal[demoObject](demoInvalidJson)
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demoSyntaxError.Error(), func() {
				json.MustUnmarshal[demoObject](demoBadJson)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				obj := json.MustUnmarshal[demoValidator](demoJson)
				must.Equal(validDemoObj, obj)
			})
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demonValidationError.Error(), func() {
				json.MustUnmarshalValid[demoValidator](demoInvalidJson)
			})
		})
	})

	t.Run("UnmarshalInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			err := json.UnmarshalInto(demoInvalidJson, &obj)
			must.NoError(err)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			err := json.UnmarshalInto(demoBadJson, &obj)
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := json.UnmarshalInto(demoJson, &obj)
			must.NoError(err)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := json.UnmarshalValidInto(demoInvalidJson, &obj)
			must.Error(err)
			must.Equal(demonValidationError, err)
			must.Equal(validDemonObj, obj)
		})
	})

	t.Run("MustUnmarshalInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			must.NotPanics(func() {
				json.MustUnmarshalInto(demoInvalidJson, &obj)
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demoSyntaxError.Error(), func() {
				var obj demoObject
				json.MustUnmarshalInto(demoBadJson, &obj)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.NotPanics(func() {
				json.MustUnmarshalInto(demoJson, &obj)
				must.Equal(validDemoObj, obj)
			})
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demonValidationError.Error(), func() {
				var obj demoValidator
				json.MustUnmarshalValidInto(demoInvalidJson, &obj)
			})
		})
	})

	t.Run("UnmarshalFile", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalFile[demoObject](demonFile.Name())
			must.NoError(err)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalFile[demoObject](badFile.Name())
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalFile[demoValidator](demoFile.Name())
			must.NoError(err)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalValidFile[demoValidator](demonFile.Name())
			must.Error(err)
			must.Equal(demonValidationError, err)
			must.Equal(validDemonObj, obj)
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalFile[demoValidator]("x")
			must.Error(err)
			must.Equal(readFileError, err)
			must.Equal(demoValidator{}, obj)
		})
	})

	t.Run("MustUnmarshalFile", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				obj := json.MustUnmarshalFile[demoObject](demonFile.Name())
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demoSyntaxError.Error(), func() {
				json.MustUnmarshalFile[demoObject](badFile.Name())
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				obj := json.MustUnmarshalFile[demoValidator](demoFile.Name())
				must.Equal(validDemoObj, obj)
			})
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demonValidationError.Error(), func() {
				json.MustUnmarshalValidFile[demoValidator](demonFile.Name())
			})
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(readFileError.Error(), func() {
				json.MustUnmarshalFile[demoValidator]("x")
			})
		})
	})

	t.Run("UnmarshalFileInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			err := json.UnmarshalFileInto(demonFile.Name(), &obj)
			must.NoError(err)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			err := json.UnmarshalFileInto(badFile.Name(), &obj)
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := json.UnmarshalFileInto(demoFile.Name(), &obj)
			must.NoError(err)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := json.UnmarshalValidFileInto(demonFile.Name(), &obj)
			must.Error(err)
			must.Equal(demonValidationError, err)
			must.Equal(validDemonObj, obj)
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := json.UnmarshalFileInto("x", &obj)
			must.Error(err)
			must.Equal(readFileError, err)
			must.Equal(demoValidator{}, obj)
		})
	})

	t.Run("MustUnmarshalFileInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			must.NotPanics(func() {
				json.MustUnmarshalFileInto(demonFile.Name(), &obj)
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demoSyntaxError.Error(), func() {
				var obj demoObject
				json.MustUnmarshalFileInto(badFile.Name(), &obj)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.NotPanics(func() {
				json.MustUnmarshalFileInto(demoFile.Name(), &obj)
				must.Equal(validDemoObj, obj)
			})
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demonValidationError.Error(), func() {
				var obj demoValidator
				json.MustUnmarshalValidFileInto(demonFile.Name(), &obj)
			})
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(readFileError.Error(), func() {
				var obj demoValidator
				json.MustUnmarshalFileInto("x", &obj)
			})
		})
	})

	t.Run("UnmarshalFS", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalFS[demoObject](filepath.Base(demonFile.Name()), tempFS)
			must.NoError(err)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalFS[demoObject](filepath.Base(badFile.Name()), tempFS)
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalFS[demoValidator](filepath.Base(demoFile.Name()), tempFS)
			must.NoError(err)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalValidFS[demoValidator](filepath.Base(demonFile.Name()), tempFS)
			must.Error(err)
			must.Equal(demonValidationError, err)
			must.Equal(validDemonObj, obj)
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			obj, err := json.UnmarshalFS[demoValidator]("x", tempFS)
			must.Error(err)
			must.Equal(readFileError, err)
			must.Equal(demoValidator{}, obj)
		})
	})

	t.Run("MustUnmarshalFS", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj := json.MustUnmarshalFS[demoObject](filepath.Base(demonFile.Name()), tempFS)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demoSyntaxError.Error(), func() {
				json.MustUnmarshalFS[demoObject](filepath.Base(badFile.Name()), tempFS)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			obj := json.MustUnmarshalFS[demoValidator](filepath.Base(demoFile.Name()), tempFS)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demonValidationError.Error(), func() {
				json.MustUnmarshalValidFS[demoValidator](filepath.Base(demonFile.Name()), tempFS)
			})
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(readFileError.Error(), func() {
				json.MustUnmarshalFS[demoValidator]("x", tempFS)
			})
		})
	})

	t.Run("UnmarshalFSInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			err := json.UnmarshalFSInto(filepath.Base(demonFile.Name()), tempFS, &obj)
			must.NoError(err)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			err := json.UnmarshalFSInto(filepath.Base(badFile.Name()), tempFS, &obj)
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := json.UnmarshalFSInto(filepath.Base(demoFile.Name()), tempFS, &obj)
			must.NoError(err)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := json.UnmarshalValidFSInto(filepath.Base(demonFile.Name()), tempFS, &obj)
			must.Error(err)
			must.Equal(demonValidationError, err)
			must.Equal(validDemonObj, obj)
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := json.UnmarshalFSInto("x", tempFS, &obj)
			must.Error(err)
			must.Equal(readFileError, err)
			must.Equal(demoValidator{}, obj)
		})
	})

	t.Run("MustUnmarshalFSInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			must.NotPanics(func() {
				json.MustUnmarshalFSInto(filepath.Base(demonFile.Name()), tempFS, &obj)
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demoSyntaxError.Error(), func() {
				var obj demoObject
				json.MustUnmarshalFSInto(filepath.Base(badFile.Name()), tempFS, &obj)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.NotPanics(func() {
				json.MustUnmarshalFSInto(filepath.Base(demoFile.Name()), tempFS, &obj)
				must.Equal(validDemoObj, obj)
			})
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demonValidationError.Error(), func() {
				var obj demoValidator
				json.MustUnmarshalValidFSInto(filepath.Base(demonFile.Name()), tempFS, &obj)
			})
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(readFileError.Error(), func() {
				var obj demoValidator
				json.MustUnmarshalFSInto("x", tempFS, &obj)
			})
		})
	})
}
