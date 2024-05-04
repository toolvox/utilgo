package mathutil_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/pkg/mathutil"
)

var tests = []struct {
	intervals [][2]int
	expected  []int
}{
	{
		intervals: [][2]int{},
		expected:  []int{},
	}, {
		intervals: [][2]int{{2, 3}},
		expected:  []int{2, 3},
	}, {
		intervals: [][2]int{{2, 3}, {5, 5}},
		expected:  []int{2, 3, 5, 5},
	}, {
		intervals: [][2]int{{13, 15}, {2, 4}},
		expected:  []int{2, 4, 13, 15},
	}, {
		intervals: [][2]int{{2, 3}, {4, 4}},
		expected:  []int{2, 4},
	}, {
		intervals: [][2]int{{7, 8}, {3, 4}, {5, 6}, {10, 12}, {1, 2}},
		expected:  []int{1, 8, 10, 12},
	}, {
		intervals: [][2]int{{3, 5}, {2, 4}},
		expected:  []int{2, 5},
	}, {
		intervals: [][2]int{{10, 12}, {4, 6}, {0, 8}, {7, 9}, {1, 3}},
		expected:  []int{0, 12},
	}, {
		intervals: [][2]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
		expected:  []int{1, 10},
	}, {
		intervals: [][2]int{{128, 191}, {192, 223}},
		expected:  []int{128, 223},
	}, {
		intervals: [][2]int{{32, 127}, {128, 159}},
		expected:  []int{32, 159},
	}, {
		intervals: [][2]int{{64, 159}, {160, 191}},
		expected:  []int{64, 191},
	}}

func Test_FindCover(t *testing.T) {
	for ti, tt := range tests {
		t.Run(fmt.Sprint(ti), func(t *testing.T) {
			res := mathutil.FindCover(tt.intervals)
			require.Equal(t, tt.expected, res)
		})
	}
}
