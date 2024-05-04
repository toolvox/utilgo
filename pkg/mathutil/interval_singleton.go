package mathutil

import (
	"github.com/toolvox/utilgo/api"
)

type singletonInterval[N Number] [1]N

func NewSingleton[N Number](v N) api.Interval[N] {
	return singletonInterval[N]{v}
}

func (si singletonInterval[N]) Min() N { return si[0] }
func (si singletonInterval[N]) Max() N { return si[0] }
func (si singletonInterval[N]) Len() N { return 1 }

func (si singletonInterval[N]) IsEmpty() bool     { return false }
func (si singletonInterval[N]) IsSingleton() bool { return true }
func (si singletonInterval[N]) IsCompound() bool  { return false }

func (si singletonInterval[N]) Enumerate(_ N) []N            { return []N{si[0]} }
func (si singletonInterval[N]) Intervals() []api.Interval[N] { return []api.Interval[N]{si} }

func (si singletonInterval[N]) Contains(value N) bool { return si[0] == value }
func (si singletonInterval[N]) Overlaps(other api.Interval[N]) bool {
	return other != nil && other.Contains(si[0])
}
func (si singletonInterval[N]) Equals(other api.Interval[N]) bool {
	return other != nil && other.IsSingleton() && other.Contains(si[0])
}

func (si singletonInterval[N]) Union(other api.Interval[N]) api.Interval[N] {
	if other == nil || other.IsEmpty() {
		return si
	}
	if other.Contains(si[0]) {
		return other
	}
	otherInts := other.Intervals()
	var res []api.Interval[N]
	var firstAfter int
	for oi, otherInt := range otherInts {
		if otherInt.Max() < si.Min() {
			res = append(res, otherInt)
			continue
		}
		firstAfter = oi
		break
	}
	res = append(res, si)
	res = append(res, otherInts[firstAfter:]...)
	return RawMerged(res...)
}

func (si singletonInterval[N]) Intersection(other api.Interval[N]) api.Interval[N] {
	if other == nil || other.IsEmpty() || !other.Contains(si[0]) {
		return NewEmpty[N]()
	}
	return si
}

func (si singletonInterval[N]) Difference(other api.Interval[N]) api.Interval[N] {
	if other == nil || other.IsEmpty() || !other.Contains(si[0]) {
		return si
	}
	return NewEmpty[N]()
}

func (si singletonInterval[N]) Resize(newSize N, growMode api.GrowFlags) api.Interval[N] {
	if newSize <= 1 {
		return si
	}

	left, right := Growths[N](growMode,
		MinValue[N](), si.Min(), newSize-1, si.Max(), MaxValue[N]())

	return NewClosed(left, right)
}

func (si singletonInterval[N]) Scale(scale float64, growMode api.GrowFlags) api.Interval[N] {
	newSize := N(scale * float64(si.Len()))
	return si.Resize(newSize, growMode)
}

func (si singletonInterval[N]) Translate(offset N, back bool) api.Interval[N] {
	minT := MinValue[N]()
	if back {
		leftSlack := si[0] - minT
		if offset >= leftSlack {
			return NewSingleton(minT)
		}
		return NewSingleton(si[0] - offset)
	}
	maxT := MaxValue[N]()
	rightSlack := maxT - si[0]
	if offset >= rightSlack {
		return NewSingleton(maxT)
	}
	return NewSingleton(si[0] + offset)
}
