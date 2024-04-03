package stringutil_test

import (
	"testing"

	"github.com/toolvox/utilgo/pkg/stringutil"

	"github.com/stretchr/testify/require"
)

func Test_Indent(t *testing.T) {
	io := stringutil.IndentOption{
		EmptyLines:   stringutil.EmptyLine_TrimHead | stringutil.EmptyLine_TrimBody | stringutil.EmptyLine_TrimTail,
		IndentPrefix: "\t",
		IndentString: "\t",
	}
	t.Run("Sanity_1", func(t *testing.T) {
		expect := `
	// Comment
	func foo () {
		// Body
		for {
			// Inner
		}
		// Back Out after 2 spaces and another following
		return
	} // With an extra empty line below
`[1:]

		r := io.Indent(`
		
		// Comment
		func foo () {
			// Body
			for {
				// Inner
			}

			 
			// Back Out after 2 spaces and another following
			 
			return
		} // With an extra empty line below
		
		`)
		if r != expect {
			require.EqualValues(t, expect, r)
		}
	})

	t.Run("Bug #1", func(t *testing.T) {
		var io = stringutil.IndentOption{stringutil.EmptyLine_TrimAll, "\t", "\t"}
		require.NotPanics(t, func() {

			res := io.Indent(`
			
			`)
			require.Equal(t, "", res)
		})
	})
}
