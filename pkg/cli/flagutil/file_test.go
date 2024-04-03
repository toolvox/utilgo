package flagutil_test

import (
	"flag"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/pkg/cli/flagutil"
	"github.com/toolvox/utilgo/pkg/errs"
)

func TestFileValue(t *testing.T) {
	tempDir := prepTemp(t)

	tests := []struct {
		name      string
		filename  string
		content   string
		expectErr error
	}{
		{
			name:     "read file",
			filename: "prefix",
			content:  "<prefix>\n",
		},
		{
			name:     "no file",
			filename: "",
			content:  "",
		},
		{
			name:      "read not existing file",
			filename:  "x",
			expectErr: errs.New(`invalid value "x" for flag -test: unable to read file 'x': open x: The system cannot find the file specified.`),
		},
	}
	for ti, tt := range tests {
		t.Run(fmt.Sprintf("%d_%s", ti, tt.name), func(t *testing.T) {
			os.Chdir(tempDir)
			defer os.Chdir("../..")

			testSet := flag.NewFlagSet(tt.name, flag.ContinueOnError)
			testSet.SetOutput(io.Discard)

			var testFlag flagutil.FileValue
			testSet.Var(&testFlag, "test", "input file")
			must := require.New(t)
			must.NotPanics(func() {
				err := testSet.Parse([]string{"-test", tt.filename})
				if tt.expectErr != nil {
					must.Error(err)
					must.Equal(tt.expectErr.Error(), err.Error())
					return
				}
				must.NoError(err)
				must.Equal(fmt.Sprintf(`"%s"`, tt.filename), testFlag.String())
				must.Equal(tt.content, string(testFlag.Get().([]byte)))
				reader := testFlag.Reader()
				must.NotNil(reader)
				bs, err := io.ReadAll(reader)
				must.NoError(err)
				must.Equal(tt.content, string(bs))
			})
		})
	}
}
