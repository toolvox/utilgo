package lexer_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"utilgo/pkg/cli/lexer"
)

var tests = []struct {
	args            string
	boolFlags       []string
	expectedActions []string
	expectedFlags   map[string]string
}{
	/*0*/ {args: "x", expectedFlags: map[string]string{}},
	/*1*/ {args: "x -", expectedFlags: map[string]string{}},
	/*2*/ {args: "x --", expectedFlags: map[string]string{}},
	/*3*/ {args: "x ---", expectedFlags: map[string]string{}},
	/*4*/ {args: "x --=", expectedFlags: map[string]string{}},
	/*5*/ {args: "x --==", expectedFlags: map[string]string{}},

	/*6*/ {args: "x", boolFlags: []string{"a"},
		expectedFlags: map[string]string{"a": "false"}},
	/*7*/ {args: "x a",
		expectedActions: []string{"a"},
		expectedFlags:   map[string]string{}},
	/*8*/ {args: "x -a",
		expectedFlags: map[string]string{"a": ""}},
	/*9*/ {args: "x -a", boolFlags: []string{"a"},
		expectedFlags: map[string]string{"a": "true"}},

	/*10*/ {args: "x --a", boolFlags: []string{"a"},
		expectedFlags: map[string]string{"a": "true"}},
	/*11*/ {args: "x ---a", boolFlags: []string{"a"},
		expectedFlags: map[string]string{"a": "true"}},
	/*12*/ {args: "x --a=", boolFlags: []string{"a"},
		expectedFlags: map[string]string{"a": "false"}},
	/*13*/ {args: "x --a==", boolFlags: []string{"a"},
		expectedFlags: map[string]string{"a": "="}},
	/*14*/ {args: "x --a===", boolFlags: []string{"a"},
		expectedFlags: map[string]string{"a": "=="}},

	/*15*/ {
		args:            "x a b",
		boolFlags:       []string{},
		expectedActions: []string{"a", "b"},
		expectedFlags:   map[string]string{},
	},
	/*16*/ {
		args:            "x -a b",
		boolFlags:       []string{},
		expectedActions: []string{},
		expectedFlags:   map[string]string{"a": "b"},
	},
	/*17*/ {
		args:            "x -a b",
		boolFlags:       []string{"a"},
		expectedActions: []string{"b"},
		expectedFlags:   map[string]string{"a": "true"},
	},
	/*18*/ {
		args:            "x -a b",
		boolFlags:       []string{"b"},
		expectedActions: []string{},
		expectedFlags:   map[string]string{"a": "b", "b": "false"},
	},
	/*19*/ {
		args:            "x -a b",
		boolFlags:       []string{"a", "b"},
		expectedActions: []string{"b"},
		expectedFlags:   map[string]string{"a": "true", "b": "false"},
	},
	/*20*/ {
		args:            "x -a=b",
		boolFlags:       []string{},
		expectedActions: []string{},
		expectedFlags:   map[string]string{"a": "b"},
	},
	/*21*/ {
		args:            "x -a=b",
		boolFlags:       []string{"a"},
		expectedActions: []string{},
		expectedFlags:   map[string]string{"a": "b"},
	},
	/*22*/ {
		args:            "x -a=b",
		boolFlags:       []string{"b"},
		expectedActions: []string{},
		expectedFlags:   map[string]string{"a": "b", "b": "false"},
	},
	/*23*/ {
		args:            "x -a=b",
		boolFlags:       []string{"b", "a"},
		expectedActions: []string{},
		expectedFlags:   map[string]string{"a": "b", "b": "false"},
	},
	/*24*/ {
		args:            "x -a= b",
		boolFlags:       []string{"a"},
		expectedActions: []string{"b"},
		expectedFlags:   map[string]string{"a": "false"},
	},
	/*25*/ {
		args:            "x -a== b",
		boolFlags:       []string{"a"},
		expectedActions: []string{"b"},
		expectedFlags:   map[string]string{"a": "="},
	},
	/*26*/ {
		args:            "x -a== b",
		boolFlags:       []string{},
		expectedActions: []string{"b"},
		expectedFlags:   map[string]string{"a": "="},
	},
	/*27*/ {
		args:            "x =a- --b=== ---c=d e --f g -h --i==---j k",
		boolFlags:       []string{"b", "h"},
		expectedActions: []string{"=a-", "e", "k"},
		expectedFlags: map[string]string{
			"b": "==",
			"c": "d",
			"f": "g",
			"h": "true",
			"i": "=---j",
		},
	},
}

func Test_TokenizeFlags(t *testing.T) {
	t.Run("Sanity", func(t *testing.T) {
		for ti, tt := range tests {
			t.Run(fmt.Sprintf("%d_%s", ti, tt.args), func(t *testing.T) {
				must := require.New(t)
				must.NotPanics(func() {
					testTokenizer := lexer.NewFlagTokenizer(tt.boolFlags...)
					must.NotNil(testTokenizer)

					testTokenizer.TokenizeFlags(strings.Split(tt.args, " "))

					must.NotNil(testTokenizer.Actions)
					// must.Len(testTokenizer.Actions, len(tt.expectedActions))
					must.ElementsMatch(tt.expectedActions, testTokenizer.Actions[1:])

					must.NotNil(testTokenizer.Tokenized)
					// must.Len(testTokenizer.Tokenized, len(tt.expectedFlags))
					must.Equal(tt.expectedFlags, testTokenizer.Tokenized)
				})
			})
		}
	})
}
