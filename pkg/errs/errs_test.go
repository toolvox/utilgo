package errs_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/pkg/errs"
)

// TestWrapperError FOO
func Test_Error(t *testing.T) {
	t.Run("Sanity", func(t *testing.T) {
		err1string := "example error 1"
		err2string := "example error 2"

		err1 := errs.New(err1string)
		err2 := errs.New(err2string)

		must := require.New(t)
		must.Equal(err1string, err1.Error())
		must.Error(err1)
		must.Equal(err2string, err2.Error())
		must.Error(err2)

		err2again := errs.Newf("%s error %d", "example", 2)
		must.Equal(err2, err2again)
		must.Error(err2again)
		must.ErrorIs(err2again, err2)
	})
}
