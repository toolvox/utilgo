package mathutil

import "github.com/toolvox/utilgo/api"

func FindCover[T Number](intervals [][2]T) []T {
	if len(intervals) == 0 {
		return []T{}
	}
	if len(intervals) == 1 {
		return intervals[0][:]
	}
	// Custom insertion sort for potentially better performance in nearly-sorted data
	for i := 1; i < len(intervals); i++ {
		j := i
		for j > 0 && intervals[j][0] < intervals[j-1][0] {
			intervals[j], intervals[j-1] = intervals[j-1], intervals[j]
			j--
		}
	}
	// In-place merge intervals
	idx := 0
	for _, interval := range intervals {
		if MinValue[T]() == interval[0] || interval[0]-1 <= intervals[idx][1] {
			if interval[1] > intervals[idx][1] {
				intervals[idx][1] = interval[1]
			}
		} else {
			idx++
			intervals[idx] = interval
		}
	}
	// Flatten the intervals into a single slice
	result := make([]T, 0, 2*(idx+1))
	for i := 0; i <= idx; i++ {
		result = append(result, intervals[i][0], intervals[i][1])
	}
	return result
}

func initialGrowths[T Number](gf api.GrowFlags, growth T) (T, T) {
	switch gf & 3 {
	case 0:
		return growth / 2, growth - (growth / 2)
	case 1:
		return growth, 0
	case 2:
		return 0, growth
	case 3:
		return growth - (growth / 2), growth / 2
	default:
		panic("what how!")
	}
}

func Growths[T Number](gf api.GrowFlags, leftBound, leftEdge, growth, rightEdge, rightBound T) (T, T) {
	leftGrow, rightGrow := initialGrowths(gf, growth)
	leftSlack, rightSlack := leftEdge-leftBound, rightBound-rightEdge
	var doOverflow bool = gf&4 == 0

	if !doOverflow {
		leftGrow = min(leftSlack, leftGrow)
		rightGrow = min(rightSlack, rightGrow)
	}

	if leftGrow > 0 {
		if leftGrow > leftSlack {
			leftEdge = leftBound
			if doOverflow {
				rightGrow += leftGrow - leftSlack
			}
			leftSlack = 0
		} else {
			leftEdge -= leftGrow
			leftSlack -= leftGrow
		}
		leftGrow = 0
	}

	if rightGrow > 0 {
		if rightGrow > rightSlack {
			rightEdge = rightBound
			if doOverflow {
				leftGrow += rightGrow - rightSlack
			}
			rightSlack = 0
		} else {
			rightEdge += rightGrow
			rightSlack -= rightGrow
		}
		rightGrow = 0
	}

	if leftGrow > 0 {
		if leftGrow > leftSlack {
			return leftBound, rightBound
		}
		leftEdge -= leftGrow
		leftSlack -= leftGrow
		leftGrow = 0
	}

	return leftEdge, rightEdge
}
