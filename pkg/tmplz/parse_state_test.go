package tmplz_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/pkg/errs"
	. "github.com/toolvox/utilgo/pkg/tmplz"
)

func Test_State_Parse(t *testing.T) {
	t.Run("Incomplete", func(t *testing.T) {
		tests := []string{
			"@ ",      // 0
			"@?",      // 1
			"@",       // 2
			"@@ @",    // 3
			"@@",      // 4
			"@@@ ",    // 5
			"@@@?",    // 6
			"@@@ @",   // 7
			"@@@",     // 8
			"@@@@",    // 9
			"@1",      // 10
			"@3abc",   // 11
			"@@_3abc", // 12
		}

		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti, "_", tt, "_"), func(t *testing.T) {
				expectErr := errs.Newf("could not complete variable from: '%s'", tt)
				state := NewState(tt)
				require.NotPanics(t, func() {
					node, cont, err := state.Parse(0, nil)
					assert.ErrorIs(t, err, expectErr)
					rest := tt[cont:]
					assert.Equal(t, tt, rest)
					require.Nil(t, node)
				})
			})
		}
	})

	t.Run("Incomplete_No@", func(t *testing.T) {
		tests := []string{
			"",   // 0
			"_",  // 1
			"a",  // 2
			"1",  // 3
			" ",  // 4
			"?",  // 5
			"_1", // 6
			"_ ", // 7
			"_?", // 8
			"aa", // 9
			"a1", // 10
			"a ", // 11
			"a?", // 12
			"11", // 13
			"1 ", // 14
			"1?", // 15
			" 1", // 16
			" ?", // 17
			"?1", // 18
			"? ", // 19
			"??", // 20
			"_@", // 21
			"__", // 22
			"_a", // 23
			"a@", // 24
			"a_", // 25
			"1@", // 26
			"1_", // 27
			"1a", // 28
			" @", // 29
			"  ", // 30
			" _", // 31
			" a", // 32
			"?@", // 33
			"?_", // 34
			"?a", // 35
		}

		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti, "_", tt, "_"), func(t *testing.T) {
				expectErr := errs.New("node must start with '@'")
				state := NewState(tt)
				require.NotPanics(t, func() {
					node, cont, err := state.Parse(0, nil)
					assert.ErrorIs(t, err, expectErr)
					rest := tt[cont:]
					assert.Equal(t, tt, rest)
					require.Nil(t, node)
				})
			})
		}
	})

	t.Run("ValidNoRest", func(t *testing.T) {
		tests := []struct {
			input  string
			expect string
		}{
			{input: "@_", expect: "@_"},                     // 0
			{input: "@a", expect: "@a_"},                    // 1
			{input: "@ab", expect: "@ab_"},                  // 2
			{input: "@a7", expect: "@a7_"},                  // 3
			{input: "@aa7x", expect: "@aa7x_"},              // 4
			{input: "@a@b", expect: "@a@b__"},               // 5
			{input: "@aaa@bb", expect: "@aaa@bb__"},         // 6
			{input: "@aaa@bb_@c__", expect: "@aaa@bb_@c__"}, // 7
			{input: "@aaa@bb_@c_", expect: "@aaa@bb_@c__"},  // 8
			{input: "@aaa@bb_@c", expect: "@aaa@bb_@c__"},   // 9
		}

		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti, "_", tt.input), func(t *testing.T) {
				state := NewState(tt.input)
				require.NotPanics(t, func() {
					node, cont, err := state.Parse(0, nil)
					assert.NoError(t, err)
					assert.Len(t, tt.input, cont)
					rest := tt.input[cont:]
					assert.Empty(t, rest, "rest wasn't empty")
					require.NotNil(t, node, "node was nil")
					require.Equal(t, tt.expect, node.Text.String(), "node incorrect")
				})
			})
		}
	})

	t.Run("ValidRest", func(t *testing.T) {
		tests := []struct {
			input      string
			expectNode string
			expectRest string
		}{
			{input: "@aaa?", expectNode: "@aaa_", expectRest: "?"},                // 0
			{input: "@aaa_bb", expectNode: "@aaa_", expectRest: "bb"},             // 1
			{input: "@aaa bb", expectNode: "@aaa_", expectRest: " bb"},            // 2
			{input: "@aaa@?", expectNode: "@aaa_", expectRest: "@?"},              // 3
			{input: "@_@_@", expectNode: "@_", expectRest: "@_@"},                 // 4
			{input: "@@_@__@", expectNode: "@@_@__", expectRest: "@"},             // 5
			{input: "@aaa@bb_@c__@", expectNode: "@aaa@bb_@c__", expectRest: "@"}, // 6
			{input: "@aaa@bb_@c___", expectNode: "@aaa@bb_@c__", expectRest: "_"}, // 7
		}

		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti, "_", tt.input), func(t *testing.T) {
				state := NewState(tt.input)
				require.NotPanics(t, func() {
					node, cont, err := state.Parse(0, nil)
					assert.NoError(t, err)
					rest := tt.input[cont:]
					assert.Equal(t, tt.expectRest, rest, "rest")
					if tt.expectNode != "" {
						assert.NotNil(t, node)
						assert.Equal(t, tt.expectNode, node.Text.String(), "node")
					} else {
						assert.Nil(t, node)
					}
				})
			})
		}
	})
}
