package stringutil_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"utilgo/pkg/stringutil"
)

type String = stringutil.String

type sliceTest struct {
	from, to int
	expect   string
}

func Test_String(t *testing.T) {
	tests := []struct {
		source          string
		expectedByteLen int
		expectedRuneLen int
		byteSlices      []sliceTest
		runeSlices      []sliceTest
	}{
		{source: "", expectedByteLen: 0, expectedRuneLen: 0,
			byteSlices: []sliceTest{
				{from: 0, to: 0, expect: ""},
				{from: 3, to: 0, expect: ""},
				{from: 0, to: 3, expect: ""},
				{from: -1, to: 0, expect: ""},
				{from: 0, to: -1, expect: ""},
				{from: -2, to: -2, expect: ""},
				{from: 2, to: 0, expect: ""},
			},
			runeSlices: []sliceTest{
				{from: 0, to: 0, expect: ""},
				{from: 3, to: 0, expect: ""},
				{from: 0, to: 3, expect: ""},
				{from: -1, to: 0, expect: ""},
				{from: 0, to: -1, expect: ""},
				{from: -2, to: -2, expect: ""},
				{from: 2, to: 0, expect: ""},
			},
		},
		{source: "👍", expectedByteLen: 4, expectedRuneLen: 1,
			byteSlices: []sliceTest{
				{from: -1, to: 0, expect: ""},
				{from: 0, to: -1, expect: "👍"},
			},
			runeSlices: []sliceTest{
				{from: -1, to: 0, expect: ""},
				{from: 0, to: -1, expect: "👍"},
				{from: 0, to: 10, expect: "👍"},
			},
		},
		{source: "👍👎", expectedByteLen: 2 * 4, expectedRuneLen: 2,
			byteSlices: []sliceTest{
				{from: 0, to: 4, expect: "👍"},
				{from: 4, to: 8, expect: "👎"},
				{from: -8, to: -1, expect: "👍👎"},
			},
			runeSlices: []sliceTest{
				{from: 0, to: -2, expect: "👍"},
				{from: 1, to: 3, expect: "👎"},
				{from: -2, to: -1, expect: "👍👎"},
			},
		},
		{source: "👍What \the \nope👎", expectedByteLen: 4 + 13 + 4, expectedRuneLen: 15,
			byteSlices: []sliceTest{
				{from: 0, to: 5, expect: "👍W"},
				{from: 5, to: 11, expect: "hat \th"},
				{from: 11, to: 16, expect: "e \nop"},
				{from: 16, to: 22, expect: "e👎"},
			},
			runeSlices: []sliceTest{
				{from: 0, to: 2, expect: "👍W"},
				{from: 2, to: 8, expect: "hat \th"},
				{from: 8, to: 13, expect: "e \nop"},
				{from: 13, to: 15, expect: "e👎"},
			},
		},
		{source: "Hello, 世界", expectedByteLen: 7 + 2*3, expectedRuneLen: 9,
			byteSlices: []sliceTest{
				{from: 0, to: 5, expect: "Hello"},
				{from: 7, to: 13, expect: "世界"},
			},
			runeSlices: []sliceTest{
				{from: 0, to: 5, expect: "Hello"},
				{from: 7, to: 9, expect: "世界"},
			},
		},
		{source: "世界", expectedByteLen: 2 * 3, expectedRuneLen: 2,
			byteSlices: []sliceTest{
				{from: 0, to: 3, expect: "世"},
				{from: 3, to: 6, expect: "界"},
			},
			runeSlices: []sliceTest{
				{from: 0, to: 1, expect: "世"},
				{from: 1, to: 2, expect: "界"},
			},
		},
		{source: "世👍界👎", expectedByteLen: 2*3 + 2*4, expectedRuneLen: 4,
			byteSlices: []sliceTest{
				{from: 0, to: 3, expect: "世"},
				{from: 3, to: 7, expect: "👍"},
				{from: 7, to: 10, expect: "界"},
				{from: 10, to: 14, expect: "👎"},
			},
			runeSlices: []sliceTest{
				{from: 0, to: 1, expect: "世"},
				{from: 1, to: 2, expect: "👍"},
				{from: 2, to: 3, expect: "界"},
				{from: 3, to: 4, expect: "👎"},
			},
		},
	}
	for ti, tt := range tests {
		t.Run(fmt.Sprint(ti, ".", tt.source), func(t *testing.T) {
			testString := String(tt.source)
			if testString.String() != tt.source {
				t.Errorf("expected: '%s', actual: '%s'", tt.source, testString.String())
			}
			t.Run("Length", func(t *testing.T) {
				byteLen := testString.ByteLen()
				if byteLen != tt.expectedByteLen {
					t.Errorf("byte len: expected: '%d', actual: '%d'", tt.expectedByteLen, byteLen)
				}

				runeLen := testString.Len()
				if runeLen != tt.expectedRuneLen {
					t.Errorf("rune len: expected: '%d', actual: '%d'", tt.expectedRuneLen, runeLen)
				}
			})

			t.Run("Slice", func(t *testing.T) {
				for _, bs := range tt.byteSlices {
					t.Run(fmt.Sprint("byte: ", bs), func(t *testing.T) {
						require.NotPanics(t, func() {
							byteSlice := testString.ByteSlice(bs.from, bs.to)
							if slices.Compare([]byte(bs.expect), byteSlice.Bytes()) != 0 {
								t.Errorf("byte slice: expected: '%v', actual: '%v'", bs.expect, string(byteSlice))
							}
						})
					})
				}
				for _, bs := range tt.runeSlices {
					t.Run(fmt.Sprint("rune: ", bs), func(t *testing.T) {
						runeSlice := testString.Slice(bs.from, bs.to)
						if slices.Compare([]rune(bs.expect), runeSlice) != 0 {
							t.Errorf("rune slice: expected: '%v', actual: '%v'", bs.expect, string(runeSlice))
						}
					})
				}
			})
		})
	}
}
