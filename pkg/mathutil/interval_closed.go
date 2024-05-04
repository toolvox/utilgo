package mathutil

import (
	"github.com/toolvox/utilgo/api"
)

type closedInterval[N Number] [2]N

func NewClosed[N Number](min, max N) api.Interval[N] {
	if max < min {
		return NewEmpty[N]()
	}
	if max == min {
		return NewSingleton(min)
	}

	return closedInterval[N]{min, max}
}

func NewInterval[N Number](min N, includeMin bool, max N, includeMax bool) api.Interval[N] {
	if max == min && includeMin && includeMax {
		return NewSingleton(min)
	}
	if max <= min {
		return NewEmpty[N]()
	}
	switch {
	case includeMin && includeMax:
		return closedInterval[N]{min, max}
	case includeMax:
		return closedInterval[N]{min + 1, max}
	case includeMin:
		return closedInterval[N]{min, max - 1}
	default:
		return closedInterval[N]{min + 1, max - 1}
	}
}

func (ci closedInterval[N]) Min() N { return ci[0] }
func (ci closedInterval[N]) Max() N { return ci[1] }
func (ci closedInterval[N]) Len() N { return ci[1] - ci[0] + 1 }

func (ci closedInterval[N]) IsEmpty() bool     { return false }
func (ci closedInterval[N]) IsSingleton() bool { return false }
func (ci closedInterval[N]) IsCompound() bool  { return false }

func (ci closedInterval[N]) Enumerate(step N) []N {
	if step == 0 {
		step = 1
	}

	var res []N
	for i := ci[0]; i <= ci[1]; i += step {
		res = append(res, i)
	}
	return res
}

func (ci closedInterval[N]) Intervals() []api.Interval[N] { return []api.Interval[N]{ci} }

func (ci closedInterval[N]) Contains(value N) bool { return value >= ci[0] && value <= ci[1] }

// overlapsOne assumes:
//
//	other != nil && !other.IsEmpty() && !other.IsSingleton() && !other.IsCompound()
func (ci closedInterval[N]) overlapsOne(other api.Interval[N]) bool {
	return ci[0] <= other.Max() && ci[1] >= other.Min()
}

func (ci closedInterval[N]) Overlaps(other api.Interval[N]) bool {
	if other == nil || other.IsEmpty() {
		return false
	}
	if other.IsSingleton() {
		return ci.Contains(other.Min())
	}

	for _, subInt := range other.Intervals() {
		if !ci.overlapsOne(subInt) {
			continue
		}
		return true
	}
	return false
}

func (ci closedInterval[N]) Equals(other api.Interval[N]) bool {
	if other == nil || other.IsEmpty() || other.IsSingleton() || other.IsCompound() {
		return false
	}

	return ci[0] == other.Min() && ci[1] == other.Max()
}

func (ci closedInterval[N]) Union(other api.Interval[N]) api.Interval[N] {
	if other == nil || other.IsEmpty() ||
		(other.IsSingleton() && ci.Contains(other.Min())) ||
		(!other.IsCompound() && ci.Contains(other.Min()) && ci.Contains(other.Max())) {
		return ci
	}

	var myMin, myMax N = ci[0], ci[1]
	subInts := other.Intervals()
	var res []api.Interval[N] = make([]api.Interval[N], 0, len(subInts))
	for si, subInt := range subInts {
		if subInt.Max() < myMin {
			res = append(res, subInt)
			continue
		}
		if subInt.Min() > myMax {
			res = append(res, NewClosed(myMin, myMax))
			res = append(res, subInts[si:]...)
			return RawMerged(res...)
		}
		myMin, myMax = min(myMin, subInt.Min()), max(myMax, subInt.Max())
	}
	res = append(res, NewClosed(myMin, myMax))
	return RawMerged(res...)
}

func (ci closedInterval[N]) Intersection(other api.Interval[N]) api.Interval[N] {
	if other == nil || other.IsEmpty() {
		return other
	}
	if other.IsSingleton() {
		if ci.Contains(other.Max()) {
			return other
		}
		return NewEmpty[N]()
	}
	if !other.IsCompound() {
		if !ci.overlapsOne(other) {
			return NewEmpty[N]()
		}
		return NewClosed(max(ci[0], other.Min()), min(ci[1], other.Max()))
	}

	return other.Intersection(ci)
}

func (ci closedInterval[N]) Difference(other api.Interval[N]) api.Interval[N] {
	if other == nil || other.IsEmpty() || !ci.Overlaps(other) {
		return ci
	}

	if other.IsSingleton() {
		if ci[0] == other.Max() || ci[1] == other.Min() {
			return NewInterval(ci[0], ci[0] != other.Min(), ci[1], ci[1] != other.Max())
		}
		return RawMerged(NewClosed(ci[0], other.Min()-1), NewClosed(other.Max()+1, ci[1]))
	}

	if !other.IsCompound() {
		if !ci.overlapsOne(other) {
			return ci
		}

		if other.Contains(ci[0]) {
			if other.Max() >= ci[1] {
				return NewEmpty[N]()
			}
			return NewClosed(other.Max()+1, ci[1])
		}

		if other.Contains(ci[1]) {
			if other.Min() <= ci[0] {
				return NewEmpty[N]()
			}
			return NewClosed(ci[0], other.Min()-1)
		}

		return RawMerged(NewClosed(ci[0], other.Min()-1), NewClosed(other.Max()+1, ci[1]))
	}

	var res api.Interval[N] = ci
	for _, subInt := range other.Intervals() {
		res = res.Difference(subInt)
	}
	return res
}

func (ci closedInterval[N]) Resize(newSize N, growMode api.GrowFlags) api.Interval[N] {
	currentSize := ci.Len()
	minT, maxT := MinValue[N](), MaxValue[N]()
	growth := newSize - currentSize

	left, right := Growths[N](growMode, minT, ci.Min(), growth, ci.Max(), maxT)
	return NewClosed(left, right)
}

func (ci closedInterval[N]) Scale(scale float64, growMode api.GrowFlags) api.Interval[N] {
	if scale <= 0 {
		return NewEmpty[N]()
	}

	newSize := N(float64(ci.Len()) * scale)
	return ci.Resize(newSize, growMode)
}

func (ci closedInterval[N]) Translate(offset N, back bool) api.Interval[N] {
	minT, maxT := MinValue[N](), MaxValue[N]()
	var newMin, newMax N

	if back {
		newMin = ci[0] - offset
		newMax = ci[1] - offset
		if newMin < minT {
			difference := minT - newMin
			newMin = minT
			if newMax-difference < minT {
				newMax = minT
			} else {
				newMax -= difference
			}
		}
	} else {
		newMin = ci[0] + offset
		newMax = ci[1] + offset
		if newMax > maxT {
			difference := newMax - maxT
			newMax = maxT
			if newMin+difference > maxT {
				newMin = maxT
			} else {
				newMin += difference
			}
		}
	}

	return NewClosed(newMin, newMax)
}
