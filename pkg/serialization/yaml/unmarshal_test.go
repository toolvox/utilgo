package yaml_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/pkg/serialization/yaml"

	"github.com/toolvox/utilgo/pkg/errs"
)

type demoSubObject struct {
	Flag bool `yaml:"Flag"`
}

type demoObject struct {
	Name      string        `yaml:"Name"`
	Value     int           `yaml:"Value"`
	SubObject demoSubObject `yaml:"SubObject"`
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

var demoYaml = []byte(`
Name: Demo
Value: -667
SubObject:
    Flag: true
`)

var demoInvalidYaml = []byte(`
Name: Demon
Value: 666
SubObject:
    Flag: true
`)

var demoBadYaml = []byte(`
Name:Bad
Value: Values
SubObject:
    Flag: :O
`)

func Test_YAML_Unmarshal_Sanity(t *testing.T) {
	_, demoSyntaxError := yaml.Unmarshal[demoObject](demoBadYaml)
	_, demonValidationError := yaml.UnmarshalValid[demoValidator](demoInvalidYaml)
	_, readFileError := os.ReadFile("x")
	readFileError = fmt.Errorf("read file 'x': %w", readFileError)

	tempDir := t.TempDir()
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

	tempFS := os.DirFS(tempDir)

	t.Run("Unmarshal", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.Unmarshal[demoObject](demoInvalidYaml)
			must.NoError(err)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.Unmarshal[demoObject](demoBadYaml)
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.Unmarshal[demoValidator](demoYaml)
			must.NoError(err)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalValid[demoValidator](demoInvalidYaml)
			must.Error(err)
			must.Equal(demonValidationError, err)
			must.Equal(validDemonObj, obj)
		})
	})

	t.Run("MustUnmarshal", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				obj := yaml.MustUnmarshal[demoObject](demoInvalidYaml)
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demoSyntaxError.Error(), func() {
				obj := yaml.MustUnmarshal[demoObject](demoBadYaml)
				must.Equal(demoObject{}, obj)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				obj := yaml.MustUnmarshal[demoValidator](demoYaml)
				must.Equal(validDemoObj, obj)
			})
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demonValidationError.Error(), func() {
				obj := yaml.MustUnmarshalValid[demoValidator](demoInvalidYaml)
				must.Equal(validDemonObj, obj)
			})
		})
	})

	t.Run("UnmarshalInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			must.NoError(yaml.UnmarshalInto(demoInvalidYaml, &obj))
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			err := yaml.UnmarshalInto(demoBadYaml, &obj)
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.NoError(yaml.UnmarshalInto(demoYaml, &obj))
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := yaml.UnmarshalValidInto(demoInvalidYaml, &obj)
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
				yaml.MustUnmarshalInto(demoInvalidYaml, &obj)
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demoSyntaxError.Error(), func() {
				var obj demoObject
				yaml.MustUnmarshalInto(demoBadYaml, &obj)
				must.Equal(demoObject{}, obj)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			yaml.MustUnmarshalInto(demoYaml, &obj)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demonValidationError.Error(), func() {
				var obj demoValidator
				var vav errs.Validator = obj
				_ = vav
				yaml.MustUnmarshalValidInto(demoInvalidYaml, &obj)
				must.Equal(demonObj, obj)
			})
		})
	})

	t.Run("UnmarshalFile", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalFile[demoObject](demonFile.Name())
			must.NoError(err)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalFile[demoObject](badFile.Name())
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalFile[demoValidator](demoFile.Name())
			must.NoError(err)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalValidFile[demoValidator](demonFile.Name())
			must.Error(err)
			must.Equal(demonValidationError, err)
			must.Equal(validDemonObj, obj)
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalFile[demoValidator]("x")
			must.Error(err)
			must.Equal(readFileError, err)
			must.Equal(demoValidator{}, obj)
		})
	})

	t.Run("MustUnmarshalFile", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				obj := yaml.MustUnmarshalFile[demoObject](demonFile.Name())
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demoSyntaxError.Error(), func() {
				obj := yaml.MustUnmarshalFile[demoObject](badFile.Name())
				must.Equal(demoObject{}, obj)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			must.NotPanics(func() {
				obj := yaml.MustUnmarshalFile[demoValidator](demoFile.Name())
				must.Equal(validDemoObj, obj)
			})
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(demonValidationError.Error(), func() {
				obj := yaml.MustUnmarshalValidFile[demoValidator](demonFile.Name())
				must.Equal(demonObj, obj)
			})
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			must.PanicsWithError(readFileError.Error(), func() {
				obj := yaml.MustUnmarshalFile[demoValidator]("x")
				must.Equal(demoObject{}, obj)
			})
		})
	})
	t.Run("UnmarshalFileInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			yaml.UnmarshalFileInto(demonFile.Name(), &obj)
			must.NoError(yaml.UnmarshalFileInto(demonFile.Name(), &obj))
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			err := yaml.UnmarshalFileInto(badFile.Name(), &obj)
			must.Error(err)
			must.Equal(demoSyntaxError, err)
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			yaml.UnmarshalFileInto(demoFile.Name(), &obj)
			must.NoError(yaml.UnmarshalFileInto(demoFile.Name(), &obj))
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := yaml.UnmarshalValidFileInto(demonFile.Name(), &obj)
			must.Error(err)
			must.Equal(demonValidationError.Error(), err.Error())
			must.Equal(validDemonObj, obj)
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := yaml.UnmarshalFileInto("x", &obj)
			must.Error(err)
			must.Equal(readFileError.Error(), err.Error())
			must.Equal(demoValidator{}, obj)
		})
	})

	t.Run("MustUnmarshalFileInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			must.NotPanics(func() {
				yaml.MustUnmarshalFileInto(demonFile.Name(), &obj)
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			must.Panics(func() {
				yaml.MustUnmarshalFileInto(badFile.Name(), &obj)
				must.Equal(demoValidator{}, obj)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.NotPanics(func() {
				yaml.MustUnmarshalFileInto(demoFile.Name(), &obj)
				must.Equal(validDemoObj, obj)
			})
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.Panics(func() {
				yaml.MustUnmarshalValidFileInto(demonFile.Name(), &obj)
				must.Equal(demoValidator{}, obj)
			})
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.Panics(func() {
				yaml.MustUnmarshalFileInto("x", &obj)
				must.Equal(demoValidator{}, obj)
			})
		})
	})

	t.Run("UnmarshalFS", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalFS[demoObject](filepath.Base(demonFile.Name()), tempFS)
			must.NoError(err)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalFS[demoObject](filepath.Base(badFile.Name()), tempFS)
			must.Error(err)
			must.Equal(demoSyntaxError.Error(), err.Error())
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalFS[demoValidator](filepath.Base(demoFile.Name()), tempFS)
			must.NoError(err)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalValidFS[demoValidator](filepath.Base(demonFile.Name()), tempFS)
			must.Error(err)
			must.Equal(demonValidationError.Error(), err.Error())
			must.Equal(validDemonObj, obj)
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			obj, err := yaml.UnmarshalFS[demoValidator]("x", tempFS)
			must.Error(err)
			must.Equal(readFileError.Error(), err.Error())
			must.Equal(demoValidator{}, obj)
		})
	})

	t.Run("MustUnmarshalFS", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			obj := yaml.MustUnmarshalFS[demoObject](filepath.Base(demonFile.Name()), tempFS)
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			must.Panics(func() {
				obj := yaml.MustUnmarshalFS[demoObject](filepath.Base(badFile.Name()), tempFS)
				must.Equal(demoValidator{}, obj)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			obj := yaml.MustUnmarshalFS[demoValidator](filepath.Base(demoFile.Name()), tempFS)
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			must.Panics(func() {
				obj := yaml.MustUnmarshalValidFS[demoValidator](filepath.Base(demonFile.Name()), tempFS)
				must.Equal(demoValidator{}, obj)
			})
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			must.Panics(func() {
				obj := yaml.MustUnmarshalFS[demoValidator]("x", tempFS)
				must.Equal(demoValidator{}, obj)
			})
		})
	})

	t.Run("UnmarshalFSInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			must.NoError(yaml.UnmarshalFSInto(filepath.Base(demonFile.Name()), tempFS, &obj))
			must.Equal(demonObj, obj)
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			err := yaml.UnmarshalFSInto(filepath.Base(badFile.Name()), tempFS, &obj)
			must.Error(err)
			must.Equal(demoSyntaxError.Error(), err.Error())
			must.Equal(demoObject{}, obj)
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.NoError(yaml.UnmarshalFSInto(filepath.Base(demoFile.Name()), tempFS, &obj))
			must.Equal(validDemoObj, obj)
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := yaml.UnmarshalValidFSInto(filepath.Base(demonFile.Name()), tempFS, &obj)
			must.Error(err)
			must.Equal(demonValidationError.Error(), err.Error())
			must.Equal(validDemonObj, obj)
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			err := yaml.UnmarshalFSInto("x", tempFS, &obj)
			must.Error(err)
			must.Equal(readFileError.Error(), err.Error())
			must.Equal(demoValidator{}, obj)
		})
	})

	t.Run("MustUnmarshalFSInto", func(t *testing.T) {
		t.Run("NoError", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			must.NotPanics(func() {
				yaml.MustUnmarshalFSInto(filepath.Base(demonFile.Name()), tempFS, &obj)
				must.Equal(demonObj, obj)
			})
		})
		t.Run("Error", func(t *testing.T) {
			must := require.New(t)
			var obj demoObject
			must.Panics(func() {
				yaml.MustUnmarshalFSInto(filepath.Base(badFile.Name()), tempFS, &obj)
				must.Equal(demoValidator{}, obj)
			})
		})
		t.Run("NoValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.NotPanics(func() {
				yaml.MustUnmarshalFSInto(filepath.Base(demoFile.Name()), tempFS, &obj)
				must.Equal(validDemoObj, obj)
			})
		})
		t.Run("ValidationError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.Panics(func() {
				yaml.MustUnmarshalValidFSInto(filepath.Base(demonFile.Name()), tempFS, &obj)
				must.Equal(demoValidator{}, obj)
			})
		})
		t.Run("FileError", func(t *testing.T) {
			must := require.New(t)
			var obj demoValidator
			must.Panics(func() {
				yaml.MustUnmarshalFSInto("x", tempFS, &obj)
			})
		})
	})
}
