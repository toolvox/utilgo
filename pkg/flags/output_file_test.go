package flags_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"utilgo/pkg/errs"
	"utilgo/pkg/flags"
)

func prepTemp(t *testing.T) string {
	t.Helper()
	tempDir := t.TempDir()
	assert.NoError(t, os.WriteFile(filepath.Join(tempDir, "junk"), []byte("junk"), 0644))
	assert.NoError(t, os.WriteFile(filepath.Join(tempDir, "prefix"), []byte("<prefix>\n"), 0644))
	assert.NoError(t, os.Mkdir(filepath.Join(tempDir, "deep"), 0444))
	assert.NoError(t, os.WriteFile(filepath.Join(tempDir, "deep", "zero"), []byte("secret-data"), 0000))
	assert.NoError(t, os.WriteFile(filepath.Join(tempDir, "deep", "read"), []byte("data\n\tto\n\tread\n."), 0111))
	return tempDir
}

func TestOutputFileValue(t *testing.T) {
	tempDir := prepTemp(t)

	tests := []struct {
		name          string
		filename      string
		overrideFlag  int
		expectSpecial *os.File
		errSet        error
		errWriter     error
	}{
		{name: "empty to stdout", filename: "", expectSpecial: os.Stdout},
		{name: "stdout", filename: "stdout", expectSpecial: os.Stdout},
		{name: "stderr", filename: "stderr", expectSpecial: os.Stderr},
		{name: "default truncate junk", filename: "junk"},
		{name: "flag append to prefix", filename: "prefix", overrideFlag: os.O_APPEND},
		{name: "new file", filename: "new"},
		{name: "invalid file", filename: "?.-/#/?\\0/dev/null", errSet: errs.New("invalid value \"?.-/#/?\\\\0/dev/null\" for flag -test: unable to open file '?.-/#/?\\0/dev/null': open ?.-/#/?\\0/dev/null: The filename, directory name, or volume label syntax is incorrect.")},
		{name: "no mode file", filename: "deep/zero", errSet: errs.New(`invalid value "deep/zero" for flag -test: unable to open file 'deep/zero': open deep/zero: Access is denied.`)},
		{name: "readonly file", filename: "deep/read", errSet: errs.New(`invalid value "deep/read" for flag -test: unable to open file 'deep/read': open deep/read: Access is denied.`)},
	}

	for ti, tt := range tests {
		t.Run(fmt.Sprintf("%d_%s", ti, tt.name), func(t *testing.T) {
			os.Chdir(tempDir)
			defer os.Chdir("../..")
			testSet := flag.NewFlagSet(tt.name, flag.ContinueOnError)
			var testFlag flags.OutputFileValue
			if tt.overrideFlag != 0 {
				testFlag = flags.OutputFileValue{Flag: tt.overrideFlag}
			}
			testSet.Var(flags.OutputFileDefault(&testFlag, "", tt.overrideFlag), "test", "output file")
			must := require.New(t)
			must.NotPanics(func() {
				err := testSet.Parse([]string{"-test", tt.filename})
				if tt.errSet != nil {
					must.Error(err)
					must.Equal(err.Error(), tt.errSet.Error())
					return
				}
				must.NoError(err)
				if tt.expectSpecial != nil {
					must.Equal(tt.expectSpecial, testFlag.Get())
					return
				}

				wc := testFlag.Writer()
				must.NotNil(wc)

				testContent := "Test1\n\tTest2"
				n, err := wc.Write([]byte(testContent))
				if tt.errWriter != nil {
					must.Error(err)
					must.Equal(err.Error(), tt.errWriter.Error())
					return
				}
				must.NoError(err)
				must.Len([]byte(testContent), n)
				must.NoError(wc.Close())

				actualContent, err := os.ReadFile(tt.filename)
				must.NoError(err)
				if strings.Contains(tt.name, "append") {
					testContent = fmt.Sprint("<prefix>\n", testContent)
				}
				must.Equal(testContent, string(actualContent))

			})
		})
	}
}
