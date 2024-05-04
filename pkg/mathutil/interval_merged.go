package mathutil

import (
	"github.com/toolvox/utilgo/api"
	"github.com/toolvox/utilgo/pkg/sliceutil"
)

type mergedIntervals[N Number] []api.Interval[N]

func RawMerged[N Number](intervals ...api.Interval[N]) api.Interval[N] {
	intervals = sliceutil.CropElements(intervals, nil, NewEmpty[N]())
	if len(intervals) == 0 {
		return NewEmpty[N]()
	}
	if len(intervals) == 1 {
		return NewClosed(intervals[0].Min(), intervals[0].Max())
	}
	return mergedIntervals[N](intervals)
}

func NewMerged[N Number](intervals ...api.Interval[N]) api.Interval[N] {
	intervals = sliceutil.CropElements(intervals, nil, NewEmpty[N]())
	if len(intervals) == 0 {
		return NewEmpty[N]()
	}
	if len(intervals) == 1 {
		return intervals[0]
	}

	rawIntervals := make([][2]N, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		subInts := intervals[i].Intervals()
		if subInts == nil {
			rawIntervals = append(rawIntervals, [2]N{intervals[i].Min(), intervals[i].Max()})
			continue
		}
		for _, subInt := range subInts {
			rawIntervals = append(rawIntervals, [2]N{subInt.Min(), subInt.Max()})
		}
	}
	rawResult := FindCover(rawIntervals)
	var res []api.Interval[N] = make([]api.Interval[N], len(rawResult)/2)
	for i := 0; i < len(res); i++ {
		res[i] = NewClosed(rawResult[i*2], rawResult[i*2+1])
	}
	return mergedIntervals[N](res)
}

func (mis mergedIntervals[N]) Min() N { return mis[0].Min() }
func (mis mergedIntervals[N]) Max() N { return mis[len(mis)-1].Max() }
func (mis mergedIntervals[N]) Len() N {
	var result N
	for _, v := range mis.Intervals() {
		result += v.Len()
	}
	return result
}

func (mis mergedIntervals[N]) IsEmpty() bool     { return false }
func (mis mergedIntervals[N]) IsSingleton() bool { return false }
func (mis mergedIntervals[N]) IsCompound() bool  { return true }

func (mis mergedIntervals[N]) Enumerate(step N) []N {
	var res []N
	for _, subInt := range mis {
		res = append(res, subInt.Enumerate(step)...)
	}
	return res
}

func (mis mergedIntervals[N]) Intervals() []api.Interval[N] {
	for _, v := range mis {
		if v.IsCompound() {
			panic("DO NOT WANT!")
		}
	}
	return mis
}

func (mis mergedIntervals[N]) Contains(value N) bool {
	for _, subInt := range mis {
		if subInt.Contains(value) {
			return true
		}
	}
	return false
}

func (mis mergedIntervals[N]) Overlaps(other api.Interval[N]) bool {
	for _, subInt := range mis {
		if subInt.Overlaps(other) {
			return true
		}
	}
	return false
}

func (mis mergedIntervals[N]) Equals(other api.Interval[N]) bool {
	if other == nil || other.IsEmpty() || other.IsSingleton() || !other.IsCompound() {
		return false
	}
	subInts := other.Intervals()
	if len(subInts) != len(mis) {
		return false
	}
	for si, subInt := range subInts {
		if !subInt.Equals(mis[si]) {
			return false
		}
	}
	return true

}

func (mis mergedIntervals[N]) Union(other api.Interval[N]) api.Interval[N] {
	if other == nil || other.IsEmpty() {
		return mis
	}
	if !other.IsCompound() {
		return other.Union(mis)
	}

	return NewMerged(other, mis)
}

func (mis mergedIntervals[N]) Intersection(other api.Interval[N]) api.Interval[N] {
	if other == nil || other.IsEmpty() {
		return NewEmpty[N]()
	}
	var res []api.Interval[N]
	for _, subInt := range mis {
		res = append(res, other.Intersection(subInt))
	}
	return RawMerged(res...)
}

func (mis mergedIntervals[N]) Difference(other api.Interval[N]) api.Interval[N] {
	if other == nil || other.IsEmpty() || !other.Overlaps(mis) {
		return mis
	}

	var res []api.Interval[N]
	if !other.IsCompound() {
		for _, subInt := range mis {
			res = append(res, subInt.Difference(other).Intervals()...)
		}
		return RawMerged(res...)
	}
	otherIntervals := other.Intervals()
	for _, subInt := range mis {
		for _, otherSubInt := range otherIntervals {
			subInt = subInt.Difference(otherSubInt)
		}
		res = append(res, subInt)
	}
	return RawMerged(res...)
}

func (mis mergedIntervals[N]) Resize(newSize N, growMode api.GrowFlags) api.Interval[N] {
	panic("#TODO: not implemented")
}
func (mis mergedIntervals[N]) Scale(scale float64, growMode api.GrowFlags) api.Interval[N] {
	panic("#TODO: not implemented")
}
func (mis mergedIntervals[N]) Translate(offset N, back bool) api.Interval[N] {
	panic("#TODO: not implemented")
}
